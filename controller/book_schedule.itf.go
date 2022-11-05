package controller

import (
	"cosmart/entity/dto"
)

type bookScheduleUseCase interface {
	SchedulePickup(req dto.PostSchedulePickUC) error
	GetListOfBook(gbl dto.GetBookListUC) ([]dto.GetBookListResp, error)
}
