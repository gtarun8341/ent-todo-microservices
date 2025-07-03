package main

import (
	"log"

	"ent-todo-microservices/api-gateway/router"
)


func main() {
	r := router.SetupRouter()

	log.Println("API Gateway starting on :8000")
	if err := r.Run(":8000"); err != nil {
		log.Fatalf("Failed to run API Gateway: %v", err)
	}
}

