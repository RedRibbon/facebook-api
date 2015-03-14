package models

type User struct {
	Id   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type UserPostView struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}
