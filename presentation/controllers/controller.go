package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ruspatrick/stan-svc/domain/models"
)

func PostMessage(w http.ResponseWriter, r *http.Request) {
	message := new(models.Message)
	if err := json.NewDecoder(r.Body).Decode(message); err != nil {
		return
	}

}

func GetMessages(w http.ResponseWriter, r *http.Request) {

}
