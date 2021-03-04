package models

type User struct {
	ID      string `json:"id"`
	MaxTodo int    `json:"maxTodo"`
	Pwd     string `json:"password,-"`
}
