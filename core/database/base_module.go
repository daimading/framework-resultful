package database

import "time"

type BaseModel struct {
	ID        int       `json:"id" gorm:"id"`
	CreateAt  time.Time `json:"create_at" gorm:"create_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}
