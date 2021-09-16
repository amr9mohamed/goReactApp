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

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	server.DB, err = gorm.Open(postgres.New(postgres.Config{
		DriverName: Dbdriver,
		DSN:        DBURL,
	}))
	if err != nil {
		log.Fatal("Failed to connect to database, error:", err)
	} else {
		fmt.Printf("We are connected to the database")
	}

	server.DB.Debug().AutoMigrate(&models.User{})

	server.Router = gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("ORIGIN")}
	server.Router.Use(cors.New(config))
	server.Router.Use(middlewares.JsonMiddleware())

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
