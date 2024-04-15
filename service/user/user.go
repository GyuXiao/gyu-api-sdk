package user

type User struct {
	Username string `json:"username"`
}

func NewUser(username string) *User {
	return &User{
		Username: username,
	}
}
