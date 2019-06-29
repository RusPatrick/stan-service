package repositories

import (
	"errors"
	"log"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/nats-io/stan.go"
	"github.com/ruspatrick/stan-svc/domain/models"
	pbnews "github.com/ruspatrick/stan-svc/domain/models/news"
	customerrors "github.com/ruspatrick/stan-svc/infrastructure/errors"
)

type StanConnect struct {
	stan.Conn
}

const (
	stanNewsSubject = "news"
	natsUrl         = "nats://stan:4222"
)

var (
	sc           *StanConnect
	ErrNoNewNews = errors.New("новые новости отсутсвуют")
)

func InitRepo() *StanConnect {
	scTmp, err := stan.Connect("news_stan", "news_client", stan.NatsURL(natsUrl))
	if err != nil {
		log.Fatalln(err)
	}
	sc = &StanConnect{
		Conn: scTmp,
	}

	return sc
}

func (s StanConnect) SendNews(news pbnews.News) error {
	data, err := news.Marshal()
	if err != nil {
		return customerrors.CreateServerError(err, "ошибка", "ошибка при работе c хранилищем")
	}
	if err := s.Publish(stanNewsSubject, data); err != nil {
		return customerrors.CreateServerError(err, "ошибка", "ошибка при работе c хранилищем")
	}
	return nil
}

func (s StanConnect) GetNews(durableName string, numberMessages int) ([]models.News, error) {
	news := make([]models.News, 0)
	timeout := make(chan bool)
	go func() {
		time.Sleep(time.Second)
		timeout <- true
	}()
	getMessages := func(msg *stan.Msg) {
		newspb := pbnews.News{}
		if err := proto.Unmarshal(msg.Data, &newspb); err != nil {
			log.Println("bad unmarhal: " + err.Error())
		}

		news = append(news, models.News{
			Title: newspb.Title,
			Date:  newspb.Date,
		})

		timeout <- false
	}

	sub, err := s.Subscribe(stanNewsSubject, getMessages, stan.DeliverAllAvailable(), stan.DurableName(durableName), stan.MaxInflight(1))
	if err != nil {
		return nil, customerrors.CreateServerError(err, "ошибка", "ошибка при работе c хранилищем")
	}
	defer sub.Close()

	if <-timeout {
		return nil, customerrors.CreateServerError(ErrNoNewNews, "ошибка", ErrNoNewNews.Error())
	}

	return news, nil
}
