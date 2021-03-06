package cryptography

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher struct {
}

func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{}
}

func (ph *PasswordHasher) HashAndSalt(password string) (*string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	hashedPassword := string(hash)

	return &hashedPassword, nil
}

func (ph *PasswordHasher) ComparePasswords(hashedPwd string, paswordToCompare string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(paswordToCompare))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
