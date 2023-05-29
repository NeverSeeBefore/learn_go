package models

import "time"

type User struct {
	ID       int       `gorm:"primaryKey;" json:"id"` // 主键叫ID
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Age      int       `json:"age"`
	AddTime  time.Time `json:"addTime"`
	EditTime time.Time `json:"editTime"`
	CateId   int       `json:"cateId"` // 外键
	// 使用时需要先 db.preload("UserCate");
	Cate UserCate `gorm:"foreignKey:CateId" json:"cate"` // 这个字段的值将通过查询外键填充， 规则是: 以指定的外键名或默认值(UserCateId)作为外键查询
}

func (u User) TableName() string {
	return "user"
}
