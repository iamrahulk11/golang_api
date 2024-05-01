package types

import (
	"time"
)

type UserStore interface {
	GetUserByUserId(ID int) (*User, error)
	GetAllUser() (*[]User, error)
}

type RegisterUserPayload struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID          int        `json:"id"`
	USER_ID     string     `json:"user_id"`
	USERNAME    string     `json:"username"`
	INSERTED_ON *time.Time `json:"inserted_on"`
	UPDATED_ON  *time.Time `json:"updated_on"`
}
