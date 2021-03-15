package plugins

import (
	"github.com/gin-gonic/gin"
)

type Translation struct {}

func (this *Translation) Init(c *gin.Context) {}

func (this *Translation) Do(c *gin.Context, content string) (string, bool) {
	return content, true
}
