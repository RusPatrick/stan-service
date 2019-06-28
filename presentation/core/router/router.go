package router

import (
	"net/http"

	"github.com/ruspatrick/stan-svc/presentation/controllers"

	"github.com/gorilla/mux"
)

type Route struct {
	Method      string
	Name        string
	Path        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

const (
	apiV1Prefix = "/api/v1"
)

func NewRouter() *mux.Router {
	routes := Routes{
		{
			Name:        "post-message",
			Method:      http.MethodPost,
			Path:        "/messages",
			HandlerFunc: controllers.PostMessage,
		}, {
			Name:        "get-message",
			Method:      http.MethodGet,
			Path:        "/messages",
			HandlerFunc: controllers.GetMessages,
		},
	}

	r := mux.NewRouter()
	apiV1 := r.Path(apiV1Prefix).Subrouter()

	for _, route := range routes {
		apiV1.Name(route.Name).
			Methods(route.Method).
			Path(route.Method).
			HandlerFunc(route.HandlerFunc)
	}

	return r
}
