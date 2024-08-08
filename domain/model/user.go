package model

type User struct {
	Id    int    `json:"id" param:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
