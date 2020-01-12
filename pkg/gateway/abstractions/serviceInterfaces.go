package abstractions

import (
	"github.com/ldugdale/dropper/pkg/commonAbstractions"
)

type UserService interface {
	SignUp(user commonAbstractions.UserModel) (*commonAbstractions.User, error)
}