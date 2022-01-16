package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserModel struct
type UserModel struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Username  string             `bson:"username" json:"username,omitempty"`
}

