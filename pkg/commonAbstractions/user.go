package commonAbstractions

type UserModel struct {
	Username string
	Password string
}

type User struct {
	Username string `json:"username"`
}

type UserWithToken struct {
	User
	Token string `json:"token"`
}
