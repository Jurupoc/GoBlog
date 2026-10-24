package models

import (
	"time"
)

type Post struct {
	Id       int       `json:"-"`
	Author   string    `json:"author"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"-"`
}

func NewPost(author, content string) Post {
	createdAt := time.Now() // TODO: Contructor shouldn't know the time, only recieve it
	return Post{
		Author:   author,
		Content:  content,
		CreateAt: createdAt,
	}
}
