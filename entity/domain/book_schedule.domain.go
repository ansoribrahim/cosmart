package domain

import (
	"cosmart/entity/dto"
)

type GetBooksFromAPI struct {
	Works []Works
}

type Works struct {
	Authors      []Author `json:"authors"`
	Title        string   `json:"title"`
	EditionCount int      `json:"edition_count"`
}

type Author struct {
	Name string `json:"name"`
}

func GetResponse(data GetBooksFromAPI) []dto.GetBookListResp {
	var result []dto.GetBookListResp

	for _, work := range data.Works {
		resp := dto.GetBookListResp{}
		resp.Title = work.Title
		resp.EditionNumber = work.EditionCount
		authors := ""
		for i, author := range work.Authors {
			if i != 0 {
				authors = authors + " ; " + author.Name
			} else {
				authors = author.Name
			}
		}
		resp.Author = authors

		result = append(result, resp)
	}
	return result
}
