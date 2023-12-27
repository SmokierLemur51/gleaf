package handlers

import (
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func ConvertStrToInt(f string) (int, error) {
	conversion, err := strconv.Atoi(f)
	if err != nil {
		log.Printf("Error converting %s to int.", f)
		return 0, err
	}
	return conversion, nil
}

func ConvertStrToFloat64(f string) (float64, error) {
	conversion, err := strconv.ParseFloat(f, 64)
	if err != nil {
		log.Printf("Error converting %s to float64.", f)
		return 0.0, err
	}
	return conversion, nil
}

// remove this
type LoginForm struct {
	Email          string `json:"email"`
	Username       string `json:"username"`
	HashedPassword string `json:"pasword"`
}

func HashString(s string, c int) (string, error) {
	salt, err := bcrypt.GenerateFromPassword([]byte(s), c)
	if err != nil {
		return "", err
	}
	return string(salt), nil
}
