package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pswd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pswd), 14)
	return string(bytes), err
}

func CheckPasswordHash(pswd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pswd))
	return err == nil
}
