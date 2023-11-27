package models

import (
	"crypto/sha256"
	"time"

	"github.com/surrealdb/surrealdb.go"
)

type User struct {
	BaseModel      surrealdb.Basemodel
	Id             string `json:"id"`
	Username       string `json:"username"`
	AuthHash       string `json:"auth_hash"`
	Email          string `json:"email"`
	DatePrefrence  int    `json:"date_prefrence"`
	Bio            string `json:"bio"`
	Avatar         string `json:"avatar"`
	CreatedAt      string `json:"created_at"`
	Premium        bool   `json:"premium"`
	PremiumExpires string `json:"premium_expires"`
}

func (u *User) GetTableName() string {
	return "users"
}

func (u *User) CreateHash(password string) {
	concat := u.Username + password
	u.AuthHash = string(sha256.New().Sum([]byte(concat)))
}

func (u *User) ValidateHash(password string) bool {
	concat := u.Username + password
	return u.AuthHash == string(sha256.New().Sum([]byte(concat)))
}

type UserRequest struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	Bio           string `json:"bio"`
	DatePrefrence int    `json:"date_prefrence"`
}

func (u *UserRequest) ToUser() *User {
	user := &User{
		Username:       u.Username,
		Email:          u.Email,
		Bio:            u.Bio,
		CreatedAt:      time.Now().Format(time.RFC3339),
		Premium:        false,
		PremiumExpires: "",
		Avatar:         "",
		DatePrefrence:  u.DatePrefrence,
	}
	user.CreateHash(u.Password)
    user.Id = user.GetTableName() + ":" + user.Username
	return user
}
