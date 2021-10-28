package models

import "time"

type Product struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"type:varchar(30);not null"  json:"name" form:"name"`
	Image       string     `json:"image_path"`
	Price       int        `gorm:"not null" json:"price" form:"price"`
	Stock       int        `gorm:"not null" json:"stock" form:"stock"`
	Category    string     `gorm:"type:varchar(30);not null"  json:"category" form:"category"`
	Description string     `gorm:"type:varchar(255)" json:"description" form:"description"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"-"`
}

type ListProductResponse struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
}
