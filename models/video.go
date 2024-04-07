package models

import "github.com/beego/beego/orm"

type Video struct {
	Id    int    `orm:"column(id);auto"`
	Video string `orm:"column(video);"`
	Audio string `orm:"column(audio);"`
	Text  string `orm:"column(text);"`
}

func (t *Video) TableName() string {
	return "video_details"
}

func init() {
	orm.RegisterModel(new(Video))
}
