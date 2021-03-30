package plugins

import (
	"fmt"
	"github.com/Garbrandt/tenet/pkg/config"
	"github.com/gin-gonic/gin"
	"strings"
)

var ReplaceStyle = []string{`'%s'`, `"%s"`,`(%s)`}

type AddCdn struct {
}

func (this *AddCdn) Init(c *gin.Context) {

}

func (this *AddCdn) Do(c *gin.Context, content string) (string, bool) {
	for _, value := range Files {
		for _, style := range ReplaceStyle {
			replace := fmt.Sprintf(style, value.UrlPath)
			cdn := fmt.Sprintf(style,config.Config.Cdn + value.Remote)
			fmt.Println("将会替换", replace, "为", cdn)
			content = strings.ReplaceAll(content, replace, cdn)
		}
	}
	return content, true
}
