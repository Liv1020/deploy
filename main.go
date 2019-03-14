package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", DeployHandler)
	r.POST("/", DeployHandler)
	r.GET("/log", LogHandler)

	log.Fatalln(r.Run(":4321"))
}
