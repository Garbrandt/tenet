package servers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tenet/lib/config"
	"tenet/plugins"
)

func Interpreter(c *gin.Context) {
	content := ""
	nextStep := true
	for _, i := range config.Config.Plugins {
		plugins.PluginSet[i].Init(c)
		content, nextStep = plugins.PluginSet[i].Do(c, content)
		if !nextStep {
			return
		}
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(content))
}
