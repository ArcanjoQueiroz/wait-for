package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDBConnection for MongoDB
type MongoDBConnection struct{}

func (s MongoDBConnection) connect(host string, port int, user string, password string, database string) bool {
	connectionInfo := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=admin", user, password, host, port, database)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionInfo))
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return false
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return false
	}
	return true
}
