package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// define a gin go server
	var router *gin.Engine = gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	router.GET("/people", handlers.taskProgressHandler)

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("impossible to start server: %s", err)
	}
}
