package repository

import (
	"cosmart/entity/dto"
	"cosmart/entity/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (rp *Repository) SchedulePickup(req dto.PostSchedulePickRP) error {
	return rp.db.Save(&model.BookSchedule{
		Title:         req.Title,
		Author:        req.Author,
		EditionNumber: req.EditionNumber,
		Time:          req.Time,
	}).Error
}
