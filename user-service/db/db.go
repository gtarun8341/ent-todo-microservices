package db

import (
	"log"

	"ent-todo-microservices/shared/config"
	"ent-todo-microservices/user-service/ent"

	_ "github.com/lib/pq"
)

func InitENT() *ent.Client{
	log.Println("Connecting to database...")

	client, err := ent.Open("postgres", config.DB_DSN)
	if err != nil {
		log.Fatal("failed to connect db", err)
	}
	log.Println("Successfully connected to database.")
  
	// Do not use this in prod as it will create tables and columns automatically
	// if err := client.Schema.Create(context.Background()); err != nil{
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }
	// log.Println("Successfully connected and migrated using Ent.")

	return client
}
