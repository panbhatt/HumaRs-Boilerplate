package routers

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"

	"github.com/panbhatt/HumaRs-Boilerplate/src/internal/config"
	"github.com/panbhatt/HumaRs-Boilerplate/src/internal/handlers"
)

func InitRouterAndStartServer(humaApiConfig huma.Config) (humaApi huma.API, server *http.Server) {

	switch config.Cfg.ROUTER {
	case "chi":
		api, rtr := initWithChiRouter(humaApiConfig)
		handlers.RegisterHandlers(api)
		//		http.ListenAndServe(":"+config.Cfg.API_PORT, rtr)
		server = &http.Server{
			Handler: rtr,
			Addr:    ":" + config.Cfg.API_PORT,
		}
		humaApi = api
		slog.Info("Server has been started ", "PORT", config.Cfg.API_PORT)
	case "go122":
		api, rtr := initWtihGo122Router(humaApiConfig)
		handlers.RegisterHandlers(api)
		//	http.ListenAndServe(":"+config.Cfg.API_PORT, rtr)
		server = &http.Server{
			Handler: rtr,
			Addr:    ":" + config.Cfg.API_PORT,
		}
		humaApi = api
		slog.Info("Server has been started ", "PORT", config.Cfg.API_PORT)

	default:
		slog.Info("No Correct Router has been provided. Exiting.....")
		os.Exit(1)
	}

	return humaApi, server
}
