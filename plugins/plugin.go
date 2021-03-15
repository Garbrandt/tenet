package plugins

import (
	"github.com/gin-gonic/gin"
)

// 首先这些插件只负责处理对于内容的解析和更改
// 不需要对别的进行处理
var PluginSet map[string]Plugin

type Plugin interface {
	Init(c *gin.Context)
	Do(c *gin.Context, content string) (result string, nextStep bool)
}

func init() {
	PluginSet = make(map[string]Plugin)
	Register("Favicon", &Favicon{})
	Register("AddCdn", &AddCdn{})
	Register("AddDatabaseData", &AddDatabaseData{})
	Register("ReadFileFromRequest", &ReadFileFromRequest{})
	Register("Zip", &Zip{})
}

func Register(k string, v Plugin) {
	if PluginSet == nil {
		PluginSet = make(map[string]Plugin)
	}
	PluginSet[k] = v
}
