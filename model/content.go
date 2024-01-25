package model

import "time"

type Content struct {
	ID        uint   `gorm:"primarykey"`
	MetaID    uint   `gorm:"not null;unique_index"`
	Article   string `gorm:"mediumtext"`
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"index"`
}
