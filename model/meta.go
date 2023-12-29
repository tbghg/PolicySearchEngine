package model

import "time"

type Meta struct {
	ID        uint `gorm:"primarykey"`
	ContentID uint `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"index"`
	Date      time.Time `gorm:"index"`
	Title     string    `gorm:"title"`
	Url       string    `gorm:"url"`
}
