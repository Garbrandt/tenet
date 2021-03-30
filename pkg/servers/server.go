package servers

import (
	"fmt"
	"github.com/Garbrandt/tenet/pkg/config"
	"github.com/Garbrandt/tenet/pkg/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"time"
)

var server = Server{}

type Server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func (s *Server) Initialize() {
	s.DB = db.DB
	s.Router = gin.Default()
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	s.initializeRoutes()
}

func (s *Server) Run() {
	apiPort := fmt.Sprintf(":%d", config.Config.Server.Port)
	fmt.Printf("Listening to port %s\n", apiPort)
	log.Fatal(http.ListenAndServe(apiPort, s.Router))
}

func Run() {
	server.Initialize()
	server.Run()
}
