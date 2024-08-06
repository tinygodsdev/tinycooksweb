package main

import (
	"fmt"
	"log"

	"github.com/tinygodsdev/datasdk/pkg/logger"
	"github.com/tinygodsdev/tinycooksweb/cmd/prepare"
	"github.com/tinygodsdev/tinycooksweb/internal/app"
	"github.com/tinygodsdev/tinycooksweb/internal/config"
	"github.com/tinygodsdev/tinycooksweb/internal/handler"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger := logger.NewStdLogger()

	if cfg.Dev() {
		fmt.Printf("Loaded config: %+v\n", cfg)
	}

	recipeStore, closeFunc, err := entstorage.NewEntStorage(entstorage.Config{
		StorageDriver: cfg.StorageDriver,
		StorageDSN:    cfg.StorageDSN,
		LogQueries:    cfg.LogDBQueries,
		Migrate:       true,
	}, logger)
	if err != nil {
		logger.Fatal("failed to init recipe store", "err", err)
	}
	defer closeFunc()

	a, err := app.New(cfg, logger, recipeStore)
	if err != nil {
		logger.Fatal("failed to init app", "err", err)
	}

	h, err := handler.NewHandler(a, logger, "templates/")
	if err != nil {
		logger.Fatal("failed to init handler", "err", err)
	}

	r := prepare.Mux(cfg, prepare.Store(cfg), h)

	httpServer := prepare.Server(cfg, r)
	httpServer.Addr = cfg.GetAddress()
	logger.Info("starting http server", cfg.GetAddress())
	log.Fatal(httpServer.ListenAndServe())

	logger.Info("Starting application.", "env", cfg.Env)
}
