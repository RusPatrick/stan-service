package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ruspatrick/stan-svc/application/services"
	"github.com/ruspatrick/stan-svc/domain/models"
	customerrors "github.com/ruspatrick/stan-svc/infrastructure/errors"
	"github.com/ruspatrick/stan-svc/infrastructure/repositories"
)

type Param struct {
	Name       string
	IsRequired bool
}

const (
	errInvalidRequestBody = "некорректное тело запроса"
)

var (
	errEmptyRequiredParam = errors.New("отсутствует обязательный параметр")
	DurableNameParam      = Param{"durableName", true}
)

func PostNews(w http.ResponseWriter, r *http.Request) {
	message := new(models.News)
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		writeError(w, customerrors.CreateBusinessError(err, "ошибка", errInvalidRequestBody))
	}

	if err := services.PostNews(*message); err != nil {
		writeError(w, err)
	}
	writeSuccess(w, http.StatusCreated, nil)
}

func GetNews(w http.ResponseWriter, r *http.Request) {
	durableNameValue, err := getParam(r, DurableNameParam)
	if err != nil {
		writeError(w, err)
	}

	news, err := services.GetNews(durableNameValue.GasValue())
	if err != nil && err.Error() != repositories.ErrNoNewNews.Error() {
		writeError(w, err)
		return
	} else if err != nil && err.Error() == repositories.ErrNoNewNews.Error() {
		writeSuccess(w, http.StatusNoContent, nil)
		return
	}

	response, err := json.Marshal(news)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	writeSuccess(w, http.StatusOK, response)
}

func getParam(r *http.Request, param Param) (models.NullableString, error) {
	durableName := mux.Vars(r)[param.Name]
	if param.IsRequired {
		if durableName == "" {
			return models.NullableString{}, customerrors.CreateBusinessError(errEmptyRequiredParam, "ошибка", errEmptyRequiredParam.Error())
		}
		return models.NewNullableString(durableName, true), nil
	}

	if durableName == "" {
		return models.NewNullableString(durableName, false), nil
	}
	return models.NewNullableString(durableName, true), nil
}
