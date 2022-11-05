package dto

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
	Time          string
}

type PostSchedulePickRP struct {
	Title         string
	Author        string
	EditionNumber int
	Time          string
}

type GetScheduleListReq struct {
	Time string `query:"time"`
}

type GetScheduleListUC struct {
	Time string
}

type GetScheduleListByNameRP struct {
	Time string
}

type GetScheduleListResp struct {
	Id            uint   `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	EditionNumber int    `json:"editionNumber"`
	Time          string `json:"time"`
}
