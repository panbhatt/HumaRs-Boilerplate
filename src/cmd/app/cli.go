package main

import (
	"fmt"
	"log/slog"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/humacli"

	"github.com/panbhatt/HumaRs-Boilerplate/src/internal/config"
	"github.com/panbhatt/HumaRs-Boilerplate/src/internal/config/routers"
)

/*
*
This Repreesnts the CLI option that we can give while running the program from CLI.
*/
type CliOptions struct {
	Debug bool   `doc:"Enable debug logging`
	Host  string `doc:"HOST ON WHICH TO LISTEN`
	Port  string `doc:"Port to Listen to"`
}

func main() {

	cli := humacli.New(func(hooks humacli.Hooks, options *CliOptions) {
		fmt.Printf("CLI was started with the following options Debug=%v Host=%v Port=%v\n", options.Debug, options.Host, options.Port)

		hooks.OnStart(func() {
			config.Init()

			// Similarly we need to update every single CONFIG as ENV Variable takes precendence over it.
			if options.Port != "" {
				config.Cfg.API_PORT = options.Port
			}

			apiConfig := huma.DefaultConfig(config.Cfg.API_NAME, config.Cfg.API_VERSION)
			api := routers.InitRouter(apiConfig)

			slog.Info(api.OpenAPI().OpenAPI)
		})

	})

	cli.Run()
}
