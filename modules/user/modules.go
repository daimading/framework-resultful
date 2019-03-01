package user

import "testProject/core/database"

type ForumAdmin struct {
	database.BaseModel
	UserName string `gorm:"user_name"`
	Password string `gorm:"password"`
	NickName string `gorm:"nick_name"`
	Phone    string `gorm:"phone"`
	Email    string `gorm:"email"`
	Status   int    `gorm:"status"` //1：表示可用，2：表示不可用
}

func (t *ForumAdmin) TableName() string {
	return "forum_admin"
}

func init() {
	database.RegisterTable(&ForumAdmin{})
}
