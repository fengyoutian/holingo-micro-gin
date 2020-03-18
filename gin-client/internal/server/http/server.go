package http

import (
	"fmt"
	"net/http"
	"time"

	holingo "github.com/fengyoutian/holingo-micro-gin/micro-server/api"
	"github.com/golang/protobuf/ptypes/empty"

	"github.com/fengyoutian/holingo-micro-gin/tool"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"

	"github.com/fengyoutian/holingo-util/file"
)

type ServerConfig struct {
	Network      string
	Addr         string
	Timeout      time.Duration
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	IdleTimeout  time.Duration `yaml:"idleTimeout"`
}

var svc holingo.HolingoHandler

func New(s holingo.HolingoHandler) (h *http.Server, err error) {
	var (
		cfg ServerConfig
		y   *file.YAML
	)
	if y, err = file.Load(tool.Config.GetConfigPath("http.yaml")); err != nil {
		return
	}
	if err = y.Unmarshal("server", &cfg); err != nil {
		return
	}
	logrus.Infof("cfg: %s\n ", cfg)
	svc = s

	engine := gin.Default()
	h = &http.Server{
		Addr:         cfg.Addr,
		Handler:      engine,
		ReadTimeout:  cfg.ReadTimeout * time.Second,
		WriteTimeout: cfg.WriteTimeout * time.Second,
		IdleTimeout:  cfg.IdleTimeout * time.Second,
	}
	registerRouter(engine)
	h.ListenAndServe()
	return
}

func registerRouter(e *gin.Engine) {
	e.GET("/ping", ping)
	v1 := e.Group("/v1")
	{
		v1.GET("/hello/:name", sayHello)
	}
}

func ping(c *gin.Context) {
	if err := svc.Ping(c, nil, &empty.Empty{}); err != nil {
		logrus.Errorf("ping error(%v)", err)
		c.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}
	c.String(http.StatusOK, "pong")
}

func sayHello(c *gin.Context) {
	name := c.Param("name")

	err := svc.SayHello(c, &holingo.HelloReq{
		Name:                 name,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}, &holingo.HelloResp{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":      -1,
			"msg":       "error",
			"timestamp": time.Now().UnixNano() / 1e6,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":      0,
		"msg":       "ok",
		"data":      fmt.Sprintf("Hello %s!", name),
		"timestamp": time.Now().UnixNano() / 1e6,
	})
}
