package repositories

import (
	"github.com/ruspatrick/stan-svc/domain/models"
	pbnews "github.com/ruspatrick/stan-svc/domain/models/news"
)

type SendToStaner interface {
	SendNews(news pbnews.News) error
	GetNews(channelName string) (*models.News, error)
}
