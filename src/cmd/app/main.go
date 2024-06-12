package main

import (
	"log/slog"

	"github.com/danielgtaylor/huma/v2"

	"github.com/panbhatt/HumaRs-Boilerplate/src/internal/config"
	"github.com/panbhatt/HumaRs-Boilerplate/src/internal/config/routers"
)

func main() {

	config.Init()
	apiConfig := huma.DefaultConfig(config.Cfg.API_NAME, config.Cfg.API_VERSION)
	api := routers.InitRouter(apiConfig)

	slog.Info(api.OpenAPI().OpenAPI)

}
