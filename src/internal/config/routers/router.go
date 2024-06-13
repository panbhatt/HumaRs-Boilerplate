package routers

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"

	"github.com/panbhatt/HumaRs-Boilerplate/src/internal/config"
	"github.com/panbhatt/HumaRs-Boilerplate/src/internal/handlers"
)

func InitRouter(humaApiConfig huma.Config) (humaApi huma.API) {

	switch config.Cfg.ROUTER {
	case "chi":
		api, rtr := initWithChiRouter(humaApiConfig)
		handlers.RegisterHandlers(api)
		slog.Info("Server has been started at Port ", "PORT ", config.Cfg.API_PORT)
		http.ListenAndServe(":"+config.Cfg.API_PORT, rtr)
		humaApi = api
		slog.Info("Server has been started at Port ", config.Cfg.API_PORT)
	case "go122":
		api, rtr := initWtihGo122Router(humaApiConfig)
		handlers.RegisterHandlers(api)
		slog.Info("Server has been started at Port ", config.Cfg.API_PORT)
		http.ListenAndServe(":"+config.Cfg.API_PORT, rtr)
		humaApi = api

	default:
		slog.Info("No Correct Router has been provided. Exiting.....")
		os.Exit(1)
	}

	return humaApi
}
