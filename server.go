package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupServer(serverName string) *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Msg from "+serverName+"  : received a request from "+c.Request.RemoteAddr)
	})

	log.Println("Setting up server:", serverName)
	return r
}
