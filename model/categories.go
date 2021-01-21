package model

import "gorm.io/gorm"

type Categories struct {
	Id   int
	Name string
}

type Cates struct {
	gorm.Model
	Id   int
	Name string
	Sort int
}
