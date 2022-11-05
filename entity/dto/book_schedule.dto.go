package dto

import "time"

type DataResp struct {
	Data interface{} `json:"data"`
}

type GetBookListReq struct {
	Subject string `query:"subject"`
}

type GetBookListUC struct {
	Subject string
}

type GetBookListResp struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	EditionNumber int    `json:"editionNumber"`
}

type PostSchedulePickReq struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	EditionNumber int    `json:"editionNumber"`
	Time          string `json:"time"`
}

type PostSchedulePickUC struct {
	Title         string
	Author        string
	EditionNumber int
	Time          time.Time
}

type PostSchedulePickRP struct {
	Title         string
	Author        string
	EditionNumber int
	Time          time.Time
}
