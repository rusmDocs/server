package user

import (
	"context"
	"errors"
	"github.com/rusmDocs/rusmDocs/pkg/passwords"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/rusmDocs/rusmDocs/pkg/database"
	"github.com/rusmDocs/rusmDocs/pkg/exceptionCodes"
)

func (user *User) createUser(body RegisterBody) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := database.UseCollection("users")

	body.Password = passwords.HashPassword(body.Password, body.Login)

	result, err := coll.InsertOne(ctx, body)

	if err != nil {
		return errors.New(
			exceptionCodes.MakeException(exceptionCodes.EntityExists, "user"),
		)
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
		return errors.New(
			exceptionCodes.MakeException(exceptionCodes.EntityNotFound, "user"),
		)
	}

	if !passwords.ComparePassword(user.Password, body.Password, body.Login) {
		return errors.New(
			exceptionCodes.MakeException(exceptionCodes.EntityInvalid, "user"),
		)
	}

	return nil
}
