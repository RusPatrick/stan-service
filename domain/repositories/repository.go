package repositories

import "github.com/nats-io/stan.go"

type Nats struct {
	sc stan.Conn
}
