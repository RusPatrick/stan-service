package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ruspatrick/stan-svc/application/services"
	"github.com/ruspatrick/stan-svc/domain/models"
)

func PostNews(w http.ResponseWriter, r *http.Request) {
	message := new(models.News)
	if err := json.NewDecoder(r.Body).Decode(message); err != nil {
		return
	}

	services.PostMessage(*message)
	w.WriteHeader(200)
}

func GetNews(w http.ResponseWriter, r *http.Request) {
	newsParam := mux.Vars(r)["news"]

	news, err := services.GetMessage(newsParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	response, err := json.Marshal(news)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		log.Println(err)
		return
	}

	w.WriteHeader(200)
	w.Write(response)
}
