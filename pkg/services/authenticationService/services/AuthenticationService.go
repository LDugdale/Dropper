package services

import (
	"time"

	"github.com/LDugdale/Dropper/pkg/commonAbstractions"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type AuthenticationService struct{}

func NewAuthenticationService() *AuthenticationService {
	return &AuthenticationService{}
}

func (as AuthenticationService) CreateToken(userAuthenticationDetails *commonAbstractions.ClaimsDetails) (string, error) {

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: userAuthenticationDetails.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {

	}

	return tokenString, nil
}

func (as AuthenticationService) ValidateToken(token string) (bool, error) {

	claims := &Claims{}

	return isValidToken(claims, token)
}

func (as AuthenticationService) RefreshToken(token string) (*string, error) {

	claims := &Claims{}

	isValidToken, err := isValidToken(claims, token)
	if !isValidToken {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return &token, nil
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tkn.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func isValidToken(claims *Claims, token string) (bool, error) {

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, err
		}
		return false, err
	}
	if !tkn.Valid {
		return false, err
	}

	return true, nil

}
