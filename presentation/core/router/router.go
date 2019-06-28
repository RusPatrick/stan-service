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
	routes := []Route{
		{
			Name:        "post-news",
			Method:      http.MethodPost,
			Path:        "/news",
			HandlerFunc: controllers.PostNews,
		}, {
			Name:        "get-news",
			Method:      http.MethodGet,
			Path:        "/news/{news}",
			HandlerFunc: controllers.GetNews,
		},
	}

	r := mux.NewRouter()
	// r.Name("post-news").
	// 	Methods(http.MethodPost).
	// 	Path("/news").
	// 	HandlerFunc(controllers.PostNews)

	apiV1 := r.PathPrefix(apiV1Prefix).Subrouter()

	for _, route := range routes {
		apiV1.Name(route.Name).
			Methods(route.Method).
			Path(route.Path).
			HandlerFunc(route.HandlerFunc)
	}

	return r
}
