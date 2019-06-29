package services

import (
	repositoriesI "github.com/ruspatrick/stan-svc/domain/interfaces/repositories"
	"github.com/ruspatrick/stan-svc/domain/models"
	"github.com/ruspatrick/stan-svc/infrastructure/repositories"
)

var repository repositoriesI.SendToStaner

func init() {
	repository = repositories.InitRepo()
}

func PostNews(news models.News) error {
	return repository.SendNews(news.ToEntity())
}

func GetNews(channelName string, n int) ([]models.News, error) {
	return repository.GetNews(channelName, n)
}
