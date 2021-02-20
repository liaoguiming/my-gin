package model

type ArticleId struct {
	Id int
}

type PageInfo struct {
	Page     int
	PageSize int
}

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
