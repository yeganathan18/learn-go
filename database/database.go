package database

import (
	"context"
	"github.com/learn/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Session *mongo.Client
	Users   *mongo.Collection
}

// ConnectDB connects to the database
func (db MongoDB) ConnectDB() MongoDB {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.Config.MongoUri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return MongoDB{
		Session: client,
		Users:   client.Database(config.Config.MongoDb).Collection("users"),
	}
}

// CloseDB Disconnect to MongoDB
func (db MongoDB) CloseDB() {
	err := db.Session.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
