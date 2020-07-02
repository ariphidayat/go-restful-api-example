package controllers

import (
	"fmt"
	"github.com/ariphidayat/go-restful-api-example/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName string) {
	var err error

	if dbDriver == "mysql" {
		dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dbUser, dbPassword, dbHost, dbPort, dbName)
		server.DB, err = gorm.Open(dbDriver, dbUrl)
		if err != nil {
			fmt.Printf("Can't connect to %s database ", dbDriver)
			log.Fatal("ERROR:", err)
		} else {
			fmt.Printf("Connected to the %s database", dbDriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{})

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port :8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

