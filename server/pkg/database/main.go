package database

import (
	"context"
	"fmt"
	"github.com/rusmDocs/rusmDocs/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MakeConnection() *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?ssl=false", config.Config.Database.Username, config.Config.Database.Password,
		config.Config.Database.Host, config.Config.Database.Port)
	con, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic("Database is unavailable")
	}

	return con
}

func MakeMigration() {
	con := MakeConnection()

	_, err := con.Database("main").Collection("users").Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    bson.D{{"email", 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		panic(err.Error())
	}
}

func UseCollection(coll string) *mongo.Collection {
	con := MakeConnection().Database("main")

	return con.Collection(coll)
}
