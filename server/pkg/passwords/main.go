package passwords

import (
	"crypto/sha512"
	"encoding/hex"
	"github.com/rusmDocs/rusmDocs/pkg/config"
)

func HashPassword(password string, login string) string {
	var w = sha512.New() // создание сетера для хэширования
	// основа для хэширования - пароль и логин (чтобы избежать коллизий при хэшировании одинаковых паролей)
	hash := append([]byte(password), []byte(login)...)
	passwordByte := append(hash, []byte(config.Config.App.Salt)...) // добавление "соли" к паролю

	w.Write(passwordByte)                               // установил данные в сеттер
	passwordHash := w.Sum(nil)                          // создал хэш
	passwordHashHex := hex.EncodeToString(passwordHash) // перевел хэш в строку

	return passwordHashHex
}

func ComparePassword(passwordHash string, passwordUser string, loginUser string) bool {
	passwordUser1 := HashPassword(passwordUser, loginUser) // создал хэш из пароля и логина

	return passwordUser1 == passwordHash // сравнил хэши
}
