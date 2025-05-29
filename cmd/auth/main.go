package main

import (
	"database/sql"
	"log"
	"snowApp/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:")
	}

	if config.Environment == "development" {
		log.Println("Running in development mode")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Println("Cannot connect to DB:")
	}
	defer conn.Close()

	// server, err :=
}

func runGinServer(config utils.Config) {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	if err := router.Run(config.ServerAddress); err != nil {
		log.Fatal("Cannot start HTTP server:", err)
	}
}
