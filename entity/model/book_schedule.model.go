package model

import (
	"time"
)

type BookSchedule struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"index"`
	Author        string
	EditionNumber int
	Time          time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
