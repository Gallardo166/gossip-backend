package config

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func Compare(password string, hash string) bool {
	match := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return match == nil
}
