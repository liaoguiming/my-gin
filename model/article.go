package model

import (
	//"github.com/Gre-Z/common/jtime"
)

type Article struct {
	Id          int
	CategoryId  int
	Title       string
	Description string
	//CreatedAt   jtime.JsonTime
}

type ArticleDetail struct {
	Id          int
	Title       string
	Description string
	CreatedAt   string
	Content     string
	Click       int
}

type ArticleByType struct {
	PageResult
	CatId int
}
