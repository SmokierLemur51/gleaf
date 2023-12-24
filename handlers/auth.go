/*
Some useful notes for JWT

Ok to put user credentials in token? - Not sensitive info
https://stackoverflow.com/questions/42652695/is-it-ok-to-store-user-credentials-in-the-jwt
*/

package handlers

import (
	"crypto/rand"
	"math/big"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var (
	key []byte
)


type CustomClaims struct {
	Username string
	Password string
	jwt.RegisteredClaims
}

func CreateToken(sessionID, username string)

func GenerateSessionID(length int) string {}

func GenerateKey(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLength := big.NewInt(int64(len(charset)))

	randomString := make([]byte, length)
	for i := 0; i < length; i ++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}

		randomString[i] = charset[randomIndex.Int64()]
	}
	fmt.Printf("\n\t%s\n\n",string(randomString))
	return string(randomString), nil
}

func SetKeyEnvironment(key string, length int) error {
	val, err := GenerateKey(length)
	if err != nil {return err}

	if err := os.Setenv(key, val); err != nil {return err}

	return nil
}

func LoadEnvKey(key string) string {
	return string(os.Getenv(key))
}