package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	lb := setupLoadBalancer()
	server1 := setupServer1()
	go lb.Run(":8080")
	go server1.Run(":8081")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down gracefully...")

}
