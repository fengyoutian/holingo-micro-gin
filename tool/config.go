package tool

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fengyoutian/holingo-util/file"
)

type _Config struct {
	base      string
	configDir string
}

var Config _Config

func init() {
	Config := new(_Config)
	flag.StringVar(&Config.base, "base", "", "工程根目录")
}

// Init: 初始化
func (c *_Config) Init() {
	if c.base != "" || len(c.base) > 0 {
		c.base = filepath.FromSlash(c.base)
	} else {
		c.base = file.GetParentDir(file.GetBuildDir())
		c.base = fmt.Sprintf("%s%c", c.base, os.PathSeparator)
	}
	c.configDir = fmt.Sprintf("%sconfigs%c", c.base, os.PathSeparator)
}

// GetConfigPath: 获取配置文件全路径
func (c *_Config) GetConfigPath(fileName string) string {
	return fmt.Sprintf("%s%s", c.configDir, fileName)
}
