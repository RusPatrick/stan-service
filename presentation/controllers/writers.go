package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ruspatrick/stan-svc/infrastructure/errors"
)

func writeSuccess(w http.ResponseWriter, statusCode int, body []byte, headers ...http.Header) {
	if len(headers) != 0 {
		setHeaders(w, headers[0])
	}
	w.WriteHeader(statusCode)
	w.Write(body)
}

func writeError(w http.ResponseWriter, err error) {
	switch _err := err.(type) {
	case errors.Error:
		errResponse, err := json.Marshal(_err)
		if err != nil {
			log.Println(err)
			return
		}
		w.WriteHeader(_err.StatusCode)
		w.Write(errResponse)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func setHeaders(w http.ResponseWriter, headers http.Header) {
	for k, v := range headers {
		for _, header := range v {
			w.Header().Add(k, header)
		}
	}
}
