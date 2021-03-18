package model

//import "github.com/Gre-Z/common/jtime"

type Publish struct {
	Id        int
	Content   string
	Position  string
	UserId    int
	Pics      []PublishPic `gorm:"foreignKey:PubId"`
	//CreatedAt jtime.JsonTime
}

type PublishInfo struct {
	Id       int
	Content  string
	Position string
	UserId   int
	Files    []string
}

type PublishPic struct {
	PubId int
	Pic   string
}

func (Publish) TableName() string {
	return "publish"
}

func (PublishPic) TableName() string {
	return "publish_pic"
}
