package usecase

import (
	"cosmart/entity/dto"
)

type bookScheduleRepository interface {
	SchedulePickup(req dto.PostSchedulePickRP) error
}
