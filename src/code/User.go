package code

import (
	"fmt"
)

type User struct {
	Email    string
	Password string
}

func NewUser(email string, password string) *User {
	var user User
	user.Email = email
	user.Password = password
	return &user
}

func (user User) String() string {
	return fmt.Sprintf("%v: %v", user.Email, user.Password)
}