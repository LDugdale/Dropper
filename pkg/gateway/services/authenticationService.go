package services

import (
	pb "github.com/LDugdale/dropper/proto"
	"github.com/ldugdale/dropper/pkg/log"
	"golang.org/x/net/context"
)

type AuthenticationService struct {
	logger                      log.Logger
	authenticationServiceClient pb.AuthenticationServiceClient
}

func NewAuthenticationService(logger log.Logger, authenticationServiceClient pb.AuthenticationServiceClient) *AuthenticationService {
	authenticationService := &AuthenticationService{
		logger:                      logger,
		authenticationServiceClient: authenticationServiceClient,
	}

	return authenticationService
}

func (as AuthenticationService) CreateToken(username string) (*string, error) {

	claimsDetails := &pb.ClaimsDetails{
		Username: username,
	}

	tokenResult, err := as.authenticationServiceClient.CreateToken(context.Background(), claimsDetails)
	//as.logger.LogTrace("Token: ", tokenResult.Token)
	if err != nil {
		as.logger.LogError("Error occured: ", err)
		return nil, err
	}

	return &tokenResult.Token, nil
}
