package config

import (
	"fmt"
	"github.com/rhysd/abspath"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Path struct {
	Sync       bool   `yaml:"sync"`
	Name       string `yaml:"name"`
	Owner      string `yaml:"owner"`
	Repository string `yaml:"repository"`
	Local      string `yaml:"local"`
	Branch     string `yaml:"branch"`
	Dir        string `yaml:"dir"`
	Token      string `yaml:"token"`
}

type Configs struct {
	Server struct {
		Theme    string `yaml:"theme"`
		Data     string `yaml:"data"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Port     int    `yaml:"port"`
		Favicon  string `yaml:"favicon"`
		Domain   string `yaml:"domain"`
	} `yaml:"server"`
	Sync struct {
		Open  bool   `yaml:"open"`
		Repo  string `yaml:"repo"`
		Owner string `yaml:"owner"`
		Token string `yaml:"token"`
	} `yaml:"sync"`
	Cdn     string `yaml:"cdn"`
	Plugins []string
}

var Config Configs

var SitePath = ""
var SiteDbPath = ""
var SiteWebPath = ""
var SiteAttachmentsPath = ""
var SiteBackgroundPath = ""

func init() {
	a, err := abspath.ExpandFrom("./")
	filename, _ := filepath.Abs(a.String() + "/config.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}

	SitePath, _ = filepath.Abs(a.String() + "/site_template")
	SiteBackgroundPath = fmt.Sprintf(SitePath + "/background")
	SiteDbPath = fmt.Sprintf(SitePath + "/data/" + Config.Server.Data)
	SiteWebPath = fmt.Sprintf(SitePath + "/background/" + Config.Server.Theme)
	SiteAttachmentsPath = fmt.Sprintf(SiteWebPath + "/attachments")

	for _, value := range []string{SitePath, SiteAttachmentsPath} {
		if err := EnsureDir(value); err != nil {
			fmt.Println("Directory creation failed with error: " + err.Error())
			os.Exit(1)
		}
	}

	Config.Plugins = []string{"Favicon", "ReadFileFromRequest", "AddDatabaseData", "AddCdn", "Zip"}
}

// 确定最后一个目录
func EnsureDir(dirName string) error {
	err := os.Mkdir(dirName, 0777)
	if os.IsExist(err) {
		return nil
	}

	return os.Chmod(dirName, 0777)
}
