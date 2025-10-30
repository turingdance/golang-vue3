package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(plainPwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
func ComparePassword(hashedPwd string, plainPwd string) (bool, error) {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, errors.New("密码不正确")
		} else {
			return false, err
		}
	}
	return true, nil
}
