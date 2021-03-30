package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"log"

	"github.com/Garbrandt/tenet/pkg/config"
	"github.com/Garbrandt/tenet/pkg/model"
)

var DB *gorm.DB

func init() {
	Initialize("SQLite", "www")
}

func Initialize(DBDriver, DBName string) {
	var err error
	if DBDriver == "SQLite" {
		DB, err = gorm.Open("sqlite3", fmt.Sprintf("%s/%s.db", config.SiteDbPath, DBName))
		if err != nil {
			fmt.Printf("Cannot connect to %s database", DBDriver)
			log.Println("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", DBDriver)
		}
	} else {
		fmt.Println("Unknown Driver")
	}

	DB.AutoMigrate(&model.Content{}, &model.Connection{})
}
