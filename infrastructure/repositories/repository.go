package repositories

import (
	"log"

	"github.com/gogo/protobuf/proto"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/ruspatrick/stan-svc/domain/models"
	pbnews "github.com/ruspatrick/stan-svc/domain/models/news"
)

type StanConnect struct {
	stan.Conn
}

const (
	stanNewsSubject = "news"
)

var sc *StanConnect

func InitRepo() *StanConnect {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatalln(err)
	}
	scTmp, err := stan.Connect("news_stan", "news_client", stan.NatsConn(nc))
	if err != nil {
		log.Fatalln(err)
	}
	sc = &StanConnect{
		Conn: scTmp,
	}

	return sc
}

func (s StanConnect) SendNews(news pbnews.News) error {
	data, _ := news.Marshal()
	return s.Publish(stanNewsSubject, data)
}

func (s StanConnect) GetNews(channelName string) (*models.News, error) {
	var news models.News
	getMessages := func(msg *stan.Msg) {
		newspb := pbnews.News{}
		if err := proto.Unmarshal(msg.Data, &newspb); err != nil {
			log.Println("FUCKED unmarhal")
		}

		news = models.News{
			Title: newspb.Title,
			Date:  newspb.Date,
		}
	}
	sub, err := s.Subscribe(channelName, getMessages, stan.DeliverAllAvailable())
	if err != nil {
		return nil, err
	}
	sub.SetPendingLimits(1, -1)

	return &news, nil
}
