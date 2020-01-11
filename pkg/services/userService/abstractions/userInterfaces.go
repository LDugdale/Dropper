package abstractions

type UserService interface {
	SignUp(user *UserModel) (*User, error)
	SignIn(signInDetails *UserModel) (*User, error)
}

type UserRepository interface {
	CreateUser(signUpDetails *UserModel) (int64, error)
	GetUser(username string) (*UserModel, error)
}