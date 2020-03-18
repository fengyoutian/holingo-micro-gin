package http

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fengyoutian/holingo-micro-gin/pkg/config"

	holingo "github.com/fengyoutian/holingo-micro-gin/micro-server/api"
	"github.com/golang/protobuf/ptypes/empty"

	"github.com/fengyoutian/holingo-micro-gin/tool"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"

	"github.com/fengyoutian/holingo-util/file"
)

var svc holingo.HolingoHandler

func New(s holingo.HolingoHandler) (h *http.Server, err error) {
	var (
		cfg config.HttpConfig
		y   *file.YAML
	)
	if y, err = file.Load(tool.Config.GetConfigPath("http.yaml")); err != nil {
		return
	}
	if err = y.Unmarshal("server", &cfg); err != nil {
		return
	}
	logrus.Infof("gin: %s\n ", cfg)
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
	article := e.Group("/article")
	{
		article.POST("/add", addArticle)
		article.GET("/search/:id", searchArticle)
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

	var reply holingo.HelloResp
	err := svc.SayHello(c, &holingo.HelloReq{
		Name:                 name,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}, &reply)

	render(c, reply.Content, err)
}

func addArticle(c *gin.Context) {
	req := &holingo.Article{
		Author:  c.PostForm("author"),
		Title:   c.PostForm("title"),
		Content: c.PostForm("content"),
	}

	var reply holingo.Article
	logrus.Infof("server.addArticle req: %v", req)
	err := svc.AddArticle(c, req, &reply)
	render(c, reply, err)
}

func searchArticle(c *gin.Context) {
	var reply holingo.Article
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err == nil {
		logrus.Infof("server.searchArticle id: %d", id)
		err = svc.SearchArticle(c, &holingo.Article{Id: id}, &reply)
	}

	render(c, reply, err)
}

func render(c *gin.Context, data interface{}, err error) {
	var status int
	result := gin.H{
		"timestamp": time.Now().UnixNano() / 1e6,
	}

	if err != nil {
		status = http.StatusInternalServerError
		result["code"] = -1
		result["msg"] = fmt.Sprint(err.Error())
	} else {
		status = http.StatusOK
		result["code"] = 0
		result["msg"] = "ok"
	}
	if data != nil {
		result["data"] = data
	}
	c.JSON(status, result)
}
