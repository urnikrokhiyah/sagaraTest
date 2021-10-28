package models

import "time"

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(30);unique;not null" json:"name" form:"name"`
	Email     string     `gorm:"type:varchar(30);unique;not null"  json:"email" form:"email"`
	Password  string     `gorm:"type:varchar(30);not null"  json:"-" form:"password"`
	Address   string     `gorm:"type:varchar(50)" json:"-" form:"address"`
	Role      string     `gorm:"type:ENUM('admin', 'customer', 'merchant')" json:"-"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
}

type LoginUser struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
