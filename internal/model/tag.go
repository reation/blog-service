package model

import "github.com/go-programming-tour-book/blog-service/pkg/app"

type Tag struct {
	*Model
	Name		string 	`json:"name"`
	State		uint32	`json:"state"`
}

type TagSwagger struct {
	List	[]*Tag
	Pager	*app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}
