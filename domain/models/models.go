package models

import (
	"github.com/ruspatrick/stan-svc/domain/models/news"
)

type News struct {
	Title string `json:"title,omitempty"`
	Date  string `json:"date,omitempty"`
}

func (n News) ToEntity() news.News {
	return news.News{
		Title: n.Title,
		Date:  n.Date,
	}
}
