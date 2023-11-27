package models

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/surrealdb/surrealdb.go"
)

type User struct {
	BaseModel      surrealdb.Basemodel `table:"users" json:"-"` // does not need to be serialized
	Id             string              `json:"id"`
	Username       string              `json:"username"`
	AuthHash       string              `json:"auth_hash"`
	Email          string              `json:"email"`
	DatePrefrence  int                 `json:"date_prefrence"`
	Bio            string              `json:"bio"`
	Avatar         string              `json:"avatar"`
	CreatedAt      string              `json:"created_at"`
	Premium        bool                `json:"premium"`
	PremiumExpires string              `json:"premium_expires"`
}

func (u *User) GetTableName() string {
	return "users"
}

func (u *User) CreateHash(password string) {
	// create a hash of the username and Password
	concat := u.Username + password
	u.AuthHash = fmt.Sprintf("%x", sha256.Sum256([]byte(concat)))
	println(u.AuthHash)
}

func (u *User) ValidateHash(password string) bool {
	// create a hash of the username and Password
	concat := u.Username + password
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(concat)))
	return u.AuthHash == hash
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

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
