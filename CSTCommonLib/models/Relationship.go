package models

import "github.com/surrealdb/surrealdb.go"

type Relationship struct {
	BaseModel surrealdb.Basemodel `json:"-"`
	ID        string `json:"id"`
    UserID1   string `json:"userId1"`
    UserID2   string `json:"userId2"`
    Status    string `json:"status"`
    number    int    `json:"number"`
}

func (r *Relationship) GetTableName() string {
    return "relationships"
}

func (r *Relationship) StatusFromNumber(number int) string {
    return r.ID
}
