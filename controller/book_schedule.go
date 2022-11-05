package controller

import (
	"cosmart/entity/dto"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Controller struct {
	bsUC bookScheduleUseCase
}

func New(bsUc bookScheduleUseCase) *Controller {
	return &Controller{
		bsUC: bsUc,
	}
}

func (ctrl *Controller) GetListOfBook(c echo.Context) error {
	req := &dto.GetBookListReq{}
	err := c.Bind(req)
	if err != nil {
		return err
	}

	resp, err := ctrl.bsUC.GetListOfBook(dto.GetBookListUC{Subject: req.Subject})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.DataResp{Data: resp})
}

func (ctrl *Controller) SchedulePickup(c echo.Context) error {
	p := &dto.PostSchedulePickReq{}
	err := c.Bind(p)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, err = time.Parse("2006-01-02", p.Time)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ctrl.bsUC.SchedulePickup(dto.PostSchedulePickUC{
		Title:         p.Title,
		Author:        p.Author,
		EditionNumber: p.EditionNumber,
		Time:          p.Time,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

func (ctrl *Controller) ScheduleList(c echo.Context) error {
	p := &dto.GetScheduleListReq{}
	err := c.Bind(p)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	//just for validation
	_, err = time.Parse("2006-01-02", p.Time)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := ctrl.bsUC.ScheduleList(dto.GetScheduleListUC{
		Time: p.Time,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, dto.DataResp{Data: resp})
}
