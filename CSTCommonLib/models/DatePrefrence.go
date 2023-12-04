package models

import "github.com/surrealdb/surrealdb.go"

type DatePreference struct {
	BaseModel surrealdb.Basemodel `table:"users" json:"-"` // does not need to be serialized
	Id        string              `json:"id"`
	Number    int                 `json:"pref_number"`
    Title     string              `json:"title"`
}

func (dp *DatePreference) GetTableName() string {
    return "date_prefrences"
}
