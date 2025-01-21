package models

type Post struct {
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	UserId int    `json:"userId"`
}
