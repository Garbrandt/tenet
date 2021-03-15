package plugins

import (
	"bytes"
	"compress/gzip"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Zip struct {
}

func (this *Zip) Init(c *gin.Context) {

}

func (this *Zip) Do(c *gin.Context, content string) (string, bool) {
	var buffer bytes.Buffer
	zw := gzip.NewWriter(&buffer)
	_, err := zw.Write([]byte(content))
	closeErr := zw.Close()
	if closeErr != nil {
		return content, true
	}
	if err != nil {
		return content, true
	}

	c.Writer.Header().Set("Accept-Encoding", "gzip")
	c.Writer.Header().Set("Content-Encoding", "gzip")
	c.Data(http.StatusOK, "text/html; charset=utf-8", buffer.Bytes())
	c.Abort()
	return content, false
}
