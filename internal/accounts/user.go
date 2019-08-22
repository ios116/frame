package accounts

// User model for accounts
type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// GetSomeUsers some logic
func GetSomeUsers() *User {
	return &User{
		Login:    "admin",
		Password: "134567",
	}
}
