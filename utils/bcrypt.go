package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashingPasswordFunc(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {

		panic(err)
	}
	return string(hash)
}

func CheckPasswordHashFunc(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
