package abstractions

import "github.com/LDugdale/Dropper/pkg/commonAbstractions"

type AuthenticationService interface {
	CreateToken(userAuthenticationDetails *commonAbstractions.ClaimsDetails) (string, error)
	ValidateToken(token string) (bool, error)
	RefreshToken(token string) (*string, error)
}
