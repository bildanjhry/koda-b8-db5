package models

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type UserName struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserContact struct {
	Phone string
}
