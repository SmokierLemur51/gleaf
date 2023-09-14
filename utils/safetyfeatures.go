package utils
import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"encoding/base64"
)

func GenerateHash(prehash string) []byte {
	hasher := sha256.New()
	data := []byte(prehash)
	hasher.Write(data)
	hashbytes := hasher.Sum(nil)
	return hashbytes 
}

func ConvertHashByteSliceToString(hashBytes []byte) string {
	return base64.StdEncoding.EncodeToString(hashBytes)
}


func CompareHash(hash1, hash2 []byte) bool {
	if subtle.ConstantTimeCompare(hash1, hash2) == 1 {
		fmt.Println(true)
		return true
	} 
	fmt.Println(false)
	return false
}