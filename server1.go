package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupServer1() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Msg from Server 1 : received a request from "+c.Request.RemoteAddr)
	})

	log.Println("Setting Server 1...")
	return r
}
