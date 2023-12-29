package model

import "time"

type Content struct {
	ID        uint `gorm:"primarykey"`
	MetaID    uint `gorm:"not null;index"`
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"index"`
}
