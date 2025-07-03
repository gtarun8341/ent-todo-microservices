package main

import (
	"log"

	"github.com/gtarun8341/ent-todo-microservices/user-service/db"
	"github.com/gtarun8341/ent-todo-microservices/user-service/router"
)

func main() {
	log.Println("Starting User Service...")

	client:=db.InitENT()
	defer client.Close()

	log.Println("Database connection established.")

	r := router.SetupRouter(*client)

	log.Println("User service running at http://localhost:8001")
	if err := r.Run(":8001"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
