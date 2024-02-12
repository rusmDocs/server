package user

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/rusmDocs/rusmDocs/pkg/database"
)

func (user *User) createUser(body RegisterBody) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := database.UseCollection("users")

	result, err := coll.InsertOne(ctx, bson.D{
		{"login", body.Login},
		{"password", body.Password},
		{"email", body.Email},
	})

	if err != nil {
		return errors.New("email conflict")
	}

	err = coll.FindOne(ctx, bson.D{
		{"_id", result.InsertedID.(primitive.ObjectID)},
	}).Decode(user)

	return nil
}

func (user *User) checkUser(body LoginBody) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := database.UseCollection("users")
	err := coll.FindOne(ctx, bson.D{
		{"login", body.Login},
	}).Decode(user)
	if err != nil {
		return errors.New("user not found")
	}

	if user.Password != body.Password {
		return errors.New("incorrect password")
	}

	return nil
}
