package userService

import (
	"log"
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher struct {

}

func (ph *PasswordHasher)hashAndSalt(password string) (*string, error) {

    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
    if err != nil {
        return nil, err
    }
	
	hashedPassword := string(hash)

    return &hashedPassword, nil
}

func (ph *PasswordHasher)comparePasswords(hashedPwd string, paswordToCompare string) bool {
	
    byteHash := []byte(hashedPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, []byte(paswordToCompare))
    if err != nil {
        log.Println(err)
        return false
    }
    
    return true
}