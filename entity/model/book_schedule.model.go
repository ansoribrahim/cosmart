package model

import (
	"time"
)

type BookSchedule struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"index"`
	Author        string
	EditionNumber int
	Time          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type BookScheduleReturn struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"index"`
	Author        string
	EditionNumber int
	Time          string
}
