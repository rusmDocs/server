package user

import (
	"context"
	"encoding/json"
	"github.com/rusmDocs/rusmDocs/api/auth"
	"github.com/rusmDocs/rusmDocs/pkg/exceptionCodes"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

type SignResponse struct {
	Login string `json:"login"`
	Email string `json:"email"`
}

type RegisterBody struct {
	Login    string `bson:"login" json:"login"`
	Password string `bson:"password" json:"password"`
	Email    string `bson:"email" json:"email"`
}

type LoginBody struct {
	Login    string `bson:"login" json:"login"`
	Password string `bson:"password" json:"password"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user User
	var userBody RegisterBody

	conn, err := grpc.Dial(":50051")
	defer conn.Close()

	err = json.NewDecoder(r.Body).Decode(&userBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = user.createUser(userBody)
	if err != nil {
		switch err.Error() {
		case exceptionCodes.MakeException(exceptionCodes.EntityExists, "user"):
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

	c := auth.NewAuthServiceClient(conn)
	tokens, err := c.CreateTokens(ctx, &auth.User{Id: user.ID.Hex()})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", tokens.AccessToken)
	w.Header().Set("Refresh", tokens.RefreshToken)

	_ = json.NewEncoder(w).Encode(SignResponse{Login: user.Login, Email: user.Email})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user User
	var userBody LoginBody

	conn, err := grpc.Dial(":50051")
	defer conn.Close()

	err = json.NewDecoder(r.Body).Decode(&userBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = user.checkUser(userBody)
	if err != nil {
		switch err.Error() {
		case exceptionCodes.MakeException(exceptionCodes.EntityNotFound, "user"):
			w.WriteHeader(http.StatusNotFound)
			return
		case exceptionCodes.MakeException(exceptionCodes.EntityInvalid, "user"):
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	c := auth.NewAuthServiceClient(conn)
	tokens, err := c.CreateTokens(ctx, &auth.User{Id: user.ID.Hex()})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", tokens.AccessToken)
	w.Header().Set("Refresh", tokens.RefreshToken)

	_ = json.NewEncoder(w).Encode(SignResponse{Login: user.Login, Email: user.Login})
}
