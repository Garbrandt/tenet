package plugins

import (
	"bytes"
	"fmt"
	"github.com/Garbrandt/tenet/pkg/config"
	"github.com/gin-gonic/gin"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

var FaviconPath = "/favicon.ico"

type Favicon struct {
	Handler bool
	Context *gin.Context
}

func (f *Favicon) Init(c *gin.Context) {
	f.Handler = false

	if c.Request.URL.String() == FaviconPath {
		f.Handler = true
	}
}

func (f *Favicon) Do(c *gin.Context, content string) (string, bool) {
	if f.Handler {
		c.Status(200)
		c.Writer.Header().Add("Content-type", "image/x-icon")
		buf, err := makeFavicon(config.Config.Server.Favicon)
		if err != nil {
			log.Println(err)
		}
		_, err = io.Copy(c.Writer, bytes.NewReader(buf))
		if err != nil {
			log.Println(err)
		}
		c.Abort()
		return content, false
	}
	return content, true
}

func makeFavicon(text string) ([]byte, error) {
	fontBytes, err := ioutil.ReadFile("./plugins/FZHTJW.TTF")
	if err != nil {
		return nil, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}
	cor,_ := hexToRGBA("#209CEE")
	fg, bg := image.White, image.NewUniform(cor)
	if false {
		fg, bg = image.White, image.Black
	}
	rgba := image.NewRGBA(image.Rect(0, 0, 32, 32))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(300)
	c.SetFont(f)
	c.SetFontSize(8)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	switch "none" {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}

	pt := freetype.Pt(-1, 28)
	_, err2 := c.DrawString(text, pt)
	if err2 != nil {
		return nil, err
	}
	pt.Y += c.PointToFixed(0)

	contents := new(bytes.Buffer)
	err = png.Encode(contents, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	readBuf, err := ioutil.ReadAll(contents)
	return readBuf, nil
}

func hexToRGBA(hex string) (color.RGBA, error) {
	var (
		rgba color.RGBA
		err  error
		errInvalidFormat = fmt.Errorf("invalid")
	)
	rgba.A = 0xff
	if hex[0] != '#' {
		return rgba, errInvalidFormat
	}
	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}
	switch len(hex) {
	case 7:
		rgba.R = hexToByte(hex[1])<<4 + hexToByte(hex[2])
		rgba.G = hexToByte(hex[3])<<4 + hexToByte(hex[4])
		rgba.B = hexToByte(hex[5])<<4 + hexToByte(hex[6])
	case 4:
		rgba.R = hexToByte(hex[1]) * 17
		rgba.G = hexToByte(hex[2]) * 17
		rgba.B = hexToByte(hex[3]) * 17
	default:
		err = errInvalidFormat
	}
	return rgba, err
}
