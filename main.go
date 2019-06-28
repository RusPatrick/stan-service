package main

import (
	"log"
	"net/http"

	"github.com/ruspatrick/stan-svc/presentation/core/router"
)

func main() {
	r := router.NewRouter()
	log.Println("Слушаю")
	http.ListenAndServe(":8000", r)
}
