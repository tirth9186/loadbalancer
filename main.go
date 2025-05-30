package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	lb := setupLoadBalancer()
	server1 := setupServer("Server 1")
	server2 := setupServer("Server 2")
	server3 := setupServer("Server 3")
	go lb.Run(":8080")
	go server1.Run(":8081")
	go server2.Run(":8082")
	go server3.Run(":8083")
	log.Println("Servers are running...")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down gracefully...")

}
