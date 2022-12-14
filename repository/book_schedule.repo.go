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

func (rp *Repository) GetScheduleListByName(req dto.GetScheduleListByNameRP) ([]dto.GetScheduleListResp, error) {
	dataDb := []model.BookScheduleReturn{}
	result := []dto.GetScheduleListResp{}
	err := rp.db.Raw("select * from book_schedules where time=?", req.Time).Scan(&dataDb).Error

	for _, ddb := range dataDb {
		result = append(result, dto.GetScheduleListResp{
			Id:            ddb.ID,
			Title:         ddb.Title,
			Author:        ddb.Author,
			EditionNumber: ddb.EditionNumber,
			Time:          ddb.Time,
		})
	}

	return result, err
}
