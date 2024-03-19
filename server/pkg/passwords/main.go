package passwords

import (
	"crypto/sha512"
	"encoding/hex"
	"github.com/rusmDocs/rusmDocs/pkg/config"
)

func HashPassword(password string, login string) string {
	var w = sha512.New()
	// основа для хэширования - пароль и логин (чтобы избежать коллизий при хэшировании одинаковых паролей)
	hash := append([]byte(password), []byte(login)...)
	passwordByte := append(hash, []byte(config.Config.App.Salt)...)

	w.Write(passwordByte)
	passwordHash := w.Sum(nil)
	passwordHashHex := hex.EncodeToString(passwordHash)

	return passwordHashHex
}

func ComparePassword(passwordHash string, passwordUser string, loginUser string) bool {
	passwordUser1 := HashPassword(passwordUser, loginUser)

	return passwordUser1 == passwordHash
}
