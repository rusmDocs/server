package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Login    string             `bson:"login" json:"login"`
	Password string             `bson:"password" json:"password"`
	Email    string             `bson:"email" json:"email"`
}
