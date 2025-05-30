package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func setupLoadBalancer() *gin.Engine {
	targetURL := "http://localhost:8081"

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			// Modify the request as needed, e.g., update headers
			req.URL, _ = url.Parse(targetURL)
			req.Host = req.URL.Host
			req.Header.Add("X-Forwarded-For", req.RemoteAddr) // Forward the client's IP
		},
		ModifyResponse: func(res *http.Response) error {
			// Modify the response headers as needed, e.g., remove unwanted headers
			return nil
		},
		Transport: http.DefaultTransport,
	}

	r := gin.Default()
	r.GET("/lb", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from Load Balancer!")
	})

	r.GET("/forward/server1", func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	log.Println("\nSetting up load balancer...\n")

	return r
}
