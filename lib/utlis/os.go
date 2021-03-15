package utlis

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Work(root string) []string {
	exts := []string{".html", ".htm"}
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
		}

		if Contains(exts, filepath.Ext(path)) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	return files
}

func getContent(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
	}
	return content
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
