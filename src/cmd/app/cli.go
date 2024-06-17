package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

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

		var server *http.Server
		hooks.OnStart(func() {

			config.Init()

			// Similarly we need to update every single CONFIG as ENV Variable takes precendence over it.
			if options.Port != "" {
				config.Cfg.API_PORT = options.Port
			}

			apiConfig := huma.DefaultConfig(config.Cfg.API_NAME, config.Cfg.API_VERSION)
			_, server = routers.InitRouterAndStartServer(apiConfig)

			server.ListenAndServe()

		})

		hooks.OnStop(func() {
			slog.Warn(("Server is shutting down in 2 seconds"))
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			server.Shutdown(ctx)

		})

	})

	// This is to set the Root
	rootCli := cli.Root()
	rootCli.Use = "blokexplorer"
	rootCli.Version = config.Cfg.API_VERSION

	cli.Run()
}
