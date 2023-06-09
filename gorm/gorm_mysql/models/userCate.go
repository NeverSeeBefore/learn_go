package models

type UserCate struct {
	ID    int    `gorm:"primaryKey;" json:"id"` // 主键
	Name  string `json:"name"`                  // 分类名称
	State int    `json:"state"`
}

func (u UserCate) TableName() string {
	return "user_cate"
}
