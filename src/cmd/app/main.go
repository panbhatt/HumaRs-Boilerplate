package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/danielgtaylor/huma/v2"

	"github.com/panbhatt/HumaRs-Boilerplate/src/internal/config"
	"github.com/panbhatt/HumaRs-Boilerplate/src/internal/config/routers"
	mdls "github.com/panbhatt/HumaRs-Boilerplate/src/internal/models/db"
)

func main() {

	config.Init()
	apiConfig := huma.DefaultConfig(config.Cfg.API_NAME, config.Cfg.API_VERSION)
	_, server := routers.InitRouterAndStartServer(apiConfig)

	checkDatabaseConnectivity()
	handleGracefulShutdown(server)

}

func handleGracefulShutdown(server *http.Server) {

	done := make(chan os.Signal)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Erro Occured. ", "Details", err)
		}
	}()

	<-done

	slog.Warn("Shutting Down the Server Now ... ")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Error while shutting down the server. Inititating Force Shutdown.. ")
	} else {
		slog.Error("Server Exiting ...")
	}

}

/*
*
This function will check the database connectivity by fetching first block on the Blocks Table
*/
func checkDatabaseConnectivity() {
	var db = config.Cfg.DB
	var blk mdls.Block
	db.Where("number = 0").First(&blk) // THis will search the block by Block Number.
	if blk.Hash != "" {
		slog.Info("First Block in Chain -> ", "Block = ", blk)
	} else {
		slog.Info("Unable to find the first block in the chain.")
	}
}
