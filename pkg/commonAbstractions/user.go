package commonAbstractions

type UserModel struct {
	Username string
	Password string
}

type User struct {
	Username string
}

type UserWithToken struct {
	User
	Token string
}
