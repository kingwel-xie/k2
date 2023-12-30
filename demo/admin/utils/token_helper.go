package utils

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

const tokenKey string = "kobh"

// GenerateToken returns a unique token based on the provided email string
func GenerateToken() string {
	hash, err := bcrypt.GenerateFromPassword([]byte(tokenKey), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return base64.URLEncoding.EncodeToString(hash)
}

func CompareToken(token string) bool {
	hash, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return false
	}
	return bcrypt.CompareHashAndPassword(hash, []byte(tokenKey)) == nil
}
