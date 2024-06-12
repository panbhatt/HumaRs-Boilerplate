package routers

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

/*
This is going to initial Go Router with Go 122
*/
func initWtihGo122Router(humaApiConfig huma.Config) (api huma.API, rtr *http.ServeMux) {

	router := http.NewServeMux()
	api = humago.New(router, humaApiConfig)

	return api, router
}
