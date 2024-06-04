package user

type User struct {
	Username string `json:"username"`
}

const ServiceUrl = "http://127.0.0.1:8090/api/user"
const MethodPost = "POST"

var TestUsername = "user_test1"
var AccessKey = "gyu"
var SecretKey = "abcdefg"

func NewUser(username string) *User {
	return &User{
		Username: username,
	}
}
