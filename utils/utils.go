package utils

import (
	"net/http"
	"html/template"
	"crypto/sha256"
	"crypto/subtle"
	"fmt"

	"github.com/SmokierLemur51/gleaf/models"
)

func RenderTemplate(w http.ResponseWriter, data models.PageData) error {
	tmpl, err := template.ParseFiles("templates/" + data.Page)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}


func GenerateHash(prehash string) []byte {
	hasher := sha256.New()
	data := []byte(prehash)
	hasher.Write(data)
	hashbytes := hasher.Sum(nil)
	return hashbytes 
}


func CompareHash(hash1, hash2 []byte) bool {
	if subtle.ConstantTimeCompare(hash1, hash2) == 1 {
		fmt.Println(true)
		return true
	} 
	fmt.Println(false)
	return false
}