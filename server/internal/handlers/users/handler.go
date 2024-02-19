package user

import (
	"encoding/json"
	"github.com/rusmDocs/rusmDocs/pkg/exceptionCodes"
	"net/http"
)

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
	var user User
	var userBody RegisterBody
	err := json.NewDecoder(r.Body).Decode(&userBody)
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

	_ = json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var userBody LoginBody

	err := json.NewDecoder(r.Body).Decode(&userBody)
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
		}
	}

	_ = json.NewEncoder(w).Encode(user)
}
