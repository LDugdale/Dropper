package authenticationService

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
	pb "github.com/LDugdale/Dropper/proto"
	"golang.org/x/net/context"
)



var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
 
type AuthenticationService struct {}

func NewAuthenticationService() *AuthenticationService{
	return &AuthenticationService{}
}

func (authenticationService AuthenticationService)CreateToken(context context.Context, userAuthenticationDetails *pb.UserAuthenticationDetails)  (*pb.CreateTokenResult, error) {

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: userAuthenticationDetails.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	result := &pb.CreateTokenResult {
		IsSuccess: true,
		Token: tokenString,
	}
	return result, err
}

func (authenticationService AuthenticationService)ValidateToken(context context.Context, token *pb.Token) (*pb.ValidateTokenResult, error){
	
	claims := &Claims{}

	return isValidToken(claims, token.Token)
}

func (authenticationService AuthenticationService)RefreshToken(context context.Context, tokenDto *pb.Token) (*pb.CreateTokenResult, error){

	claims := &Claims{}


	validateTokenResult, err := isValidToken(claims, tokenDto.Token)

	createTokenResult := &pb.CreateTokenResult {
		IsSuccess: validateTokenResult.IsSuccess,
		StatusCode: validateTokenResult.StatusCode,
	}

	if err != nil {
		return createTokenResult, err
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		createTokenResult.StatusCode = http.StatusBadRequest
		createTokenResult.IsSuccess = false
		return createTokenResult, nil
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		createTokenResult.StatusCode = http.StatusInternalServerError
		createTokenResult.IsSuccess = false
		return createTokenResult, err
	}

	createTokenResult.StatusCode = http.StatusOK
	createTokenResult.Token = tokenString
	return createTokenResult, nil
}

func isValidToken(claims *Claims, token string) (*pb.ValidateTokenResult, error){


	validateTokenResult := &pb.ValidateTokenResult{
		IsSuccess: true,
		StatusCode: http.StatusOK,
	}
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			validateTokenResult.StatusCode = http.StatusUnauthorized
			validateTokenResult.IsSuccess = false
			return validateTokenResult, err
		}
		validateTokenResult.StatusCode = http.StatusBadRequest
		validateTokenResult.IsSuccess = false
		return validateTokenResult, err
	}
	if !tkn.Valid {
		validateTokenResult.StatusCode = http.StatusUnauthorized
		validateTokenResult.IsSuccess = false
		return validateTokenResult, err
	}

	return validateTokenResult, err

}