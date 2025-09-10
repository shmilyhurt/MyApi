package utils

import (
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func GenerateSalt() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes), nil
}

func HashPassword(password, salt string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, salt, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	return err == nil
}
