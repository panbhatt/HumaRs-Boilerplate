package routers

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

func initWithChiRouter(humaApiConfig huma.Config) (api huma.API, rtr *chi.Mux) {

	router := chi.NewMux()
	api = humachi.New(router, humaApiConfig)

	return api, router
}
