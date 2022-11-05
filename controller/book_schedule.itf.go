package controller

import (
	"cosmart/entity/dto"
)

type bookScheduleUseCase interface {
	SchedulePickup(req dto.PostSchedulePickUC) error
	ScheduleList(req dto.GetScheduleListUC) ([]dto.GetScheduleListResp, error)
	GetListOfBook(gbl dto.GetBookListUC) ([]dto.GetBookListResp, error)
}
