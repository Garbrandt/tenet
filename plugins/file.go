package plugins

import (
	"fmt"
	"github.com/Garbrandt/tenet/pkg/config"
	"github.com/Garbrandt/tenet/pkg/utlis"
	"github.com/gin-gonic/gin"
	"log"
	"net/url"
	"path"
	"path/filepath"
	"strings"
)

var NeedSync []string

type Info struct {
	Path    string
	Remote  string
	UrlPath string
}

var Files map[string]Info
var FilesChannel chan []string

func init() {
	NeedSync = []string{".css", ".png", ".jpg", ".js"}
	Files = map[string]Info{}
	FilesChannel = make(chan []string, 10000)
	go SyncWork()
}

type ReadFileFromRequest struct {
	Dir  string
	Ext  string
	File string
}

func (this *ReadFileFromRequest) Init(c *gin.Context) {
	u, err := url.Parse(c.Request.RequestURI)
	if err != nil {
		return
	}

	dir, file := path.Split(u.Path)
	ext := filepath.Ext(file)
	if file == "" || ext == "" {
		file = "index.html"
		ext = ".html"
	}

	this.Ext = ext
	this.Dir = dir
	this.File = file
}

func (this *ReadFileFromRequest) Do(c *gin.Context, content string) (string, bool) {
	base := config.SiteWebPath
	items := strings.Split(c.Request.URL.String(), "/")
	switch items[1] {
	case "dashboard":
		base = config.SiteBackgroundPath
	case "generate":
		base = config.SiteBackgroundPath
	default:
	}

	local := fmt.Sprintf("%s", base+path.Join(this.Dir, this.File))
	if config.Config.Sync.Open {
		log.Println("开启同步")
		remote := utlis.GetMD5Hash(config.Config.Server.Domain) + path.Join(this.Dir, this.File)
		Sync(local, remote, this.Ext, c.Request.URL.EscapedPath())
	}

	if this.Ext != ".htm" && this.Ext != ".html" {
		c.File(local)
		return content, false
	}

	content = string(utlis.Convert(local))
	return content, true
}

func Sync(local, remote, ext string, urlPath string) {
	if !utlis.Contains(NeedSync, ext) {
		log.Println("不需要同步", local, remote)
		return
	}

	if _, ok := Files[remote]; ok {
		log.Println("已经同步", local, remote)
		return
	}

	log.Println("需要同步", local, remote)
	FilesChannel <- []string{local, remote, urlPath}
}

func SyncWork() {
	local := 0
	remote := 1
	urlPath := 2
	for {
		syncFile := <-FilesChannel
		log.Println("开始同步", syncFile[local], syncFile[remote], syncFile[urlPath])
		success, err := utlis.SyncFormLocal(syncFile[local], syncFile[remote])
		if success {
			syncFile[urlPath] = strings.Replace(syncFile[urlPath], "/dashboard/", "", 1)
			syncFile[urlPath] = strings.Replace(syncFile[urlPath], "/generate/", "", 1)
			Files[syncFile[remote]] = Info{
				Path:    strings.Replace(syncFile[local], config.SitePath, "", -1),
				Remote:  syncFile[remote],
				UrlPath: syncFile[urlPath],
			}
			log.Println("同步成功", syncFile[local], syncFile[remote])
			continue
		}

		if err != nil {
			log.Println("同步失败", syncFile[local], syncFile[remote], err.Error())
		}
	}
}
