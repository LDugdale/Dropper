package controller

import (
	"github.com/LDugdale/Dropper/pkg/commonAbstractions"
	"github.com/LDugdale/Dropper/pkg/services/authenticationService/abstractions"
	pb "github.com/LDugdale/Dropper/proto"
	"golang.org/x/net/context"
)

type AuthenticationServiceServer struct {
	authenticationService abstractions.AuthenticationService
}

func NewAuthenticationServiceServer(authenticationService abstractions.AuthenticationService) *AuthenticationServiceServer {
	return &AuthenticationServiceServer{
		authenticationService: authenticationService,
	}
}

func (as AuthenticationServiceServer) CreateToken(context context.Context, claimsDetails *pb.ClaimsDetails) (*pb.TokenResult, error) {

	claims := &commonAbstractions.ClaimsDetails{
		Username: claimsDetails.Username,
	}

	token, err := as.authenticationService.CreateToken(claims)
	if err != nil {
		return nil, err
	}

	tokenResult := &pb.TokenResult{
		Token: token,
	}

	return tokenResult, nil
}

func (as AuthenticationServiceServer) ValidateToken(context context.Context, tokenDto *pb.Token) (*pb.ValidateTokenResult, error) {

	isValid, err := as.authenticationService.ValidateToken(tokenDto.Token)
	if err != nil {
		return nil, err
	}

	validateTokenResult := &pb.ValidateTokenResult{
		IsValid: isValid,
	}

	return validateTokenResult, nil
}

func (as AuthenticationServiceServer) RefreshToken(context context.Context, tokenDto *pb.Token) (*pb.TokenResult, error) {

	token, err := as.authenticationService.RefreshToken(tokenDto.Token)
	if err != nil {
		return nil, err
	}

	tokenResult := &pb.TokenResult{
		Token: *token,
	}

	return tokenResult, nil
}
