package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/amr9mohamed/mainApp/api/middlewares"
	"github.com/amr9mohamed/mainApp/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (s *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	s.DB, err = gorm.Open(postgres.New(postgres.Config{
		DriverName: Dbdriver,
		DSN:        DBURL,
	}))
	if err != nil {
		log.Fatal("Failed to connect to database, error:", err)
	} else {
		fmt.Printf("We are connected to the database")
	}

	err = s.DB.Debug().AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	s.Router = gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("ORIGIN")}
	s.Router.Use(cors.New(config))
	s.Router.Use(middlewares.JsonMiddleware())

	s.initializeRoutes()
}

func (s *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
