package db

// import (
// 	"context"
// 	"log"

// 	"entgo.io/ent/dialect"
// 	"github.com/gtarun8341/ent-todo-microservicesg68/shared/config"
// )

// func InitENT() *models.client{
// 	log.Println("Connecting to database...")

// 	client, err := models.Open(dialect.Postgres,config.DB_DSN)
// 	if err != nil{
// 		log.Fatal("failed to connect db in gorm",err)
// 	}
// 	log.Println("Successfully connected to database.")

// 	if err := client.schema.Create(context.Background()); err != nil{
// 		log.Fatalf("failed creating schema resources: %v", err)
// 	}
// 	log.Println("Successfully connected and migrated using Ent.")

// 	return client
// }
