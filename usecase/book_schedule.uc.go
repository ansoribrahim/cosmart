package usecase

import (
	"cosmart/entity/domain"
	"cosmart/entity/dto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UseCase struct {
	bsRP bookScheduleRepository
}

func New(bsRp bookScheduleRepository) *UseCase {
	return &UseCase{
		bsRP: bsRp,
	}
}

func (uc *UseCase) GetListOfBook(gbl dto.GetBookListUC) ([]dto.GetBookListResp, error) {
	var result []dto.GetBookListResp
	response, err := http.Get(fmt.Sprintf("http://openlibrary.org/subjects/%s.json?published_in=1500-2022", gbl.Subject))
	if err != nil {
		return result, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	var responseAPI domain.GetBooksFromAPI
	err = json.Unmarshal(responseData, &responseAPI)
	if err != nil {
		return result, err
	}

	return domain.GetResponse(responseAPI), nil
}

func (uc *UseCase) SchedulePickup(req dto.PostSchedulePickUC) error {
	err := uc.bsRP.SchedulePickup(dto.PostSchedulePickRP{
		Title:         req.Title,
		Author:        req.Author,
		EditionNumber: req.EditionNumber,
		Time:          req.Time,
	})
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) ScheduleList(req dto.GetScheduleListUC) ([]dto.GetScheduleListResp, error) {
	result, err := uc.bsRP.GetScheduleListByName(dto.GetScheduleListByNameRP{
		Time: req.Time,
	})
	if err != nil {
		return result, err
	}

	return result, nil
}
