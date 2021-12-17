package database

import (
	"context"
	"gitlab.com/lyra/backend/user/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func (db MongoDB) GetAllUsernames() (interface{}, error) {
	var results []models.UserModel
	var err error

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := db.Users.Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem models.UserModel
		err = cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	if err = cur.Err(); err != nil {
		return nil, err
	}
	err = cur.Close(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}


func (db MongoDB) AddUsername(username string) (interface{}, error) {
	var err error
	var results models.UserModel

	results.ID = primitive.NewObjectID()
	results.Username = username
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_, err = db.Users.InsertOne(ctx, results)
	if err != nil {
		return nil, err
	}
	return results, nil
}
