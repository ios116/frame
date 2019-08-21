package intetrnal

// User model for users
type User struct {
	Login string `json:"login"`
    Password string `json:"password"`
}

// GetUsers some logic
func GetUsers() *User {
	return &User{
		Login: "admin",
		Password: "134567",
	}
}
