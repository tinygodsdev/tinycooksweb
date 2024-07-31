package main

import (
	"fmt"
	"log"

	"github.com/tinygodsdev/datasdk/pkg/logger"
	"github.com/tinygodsdev/tinycooksweb/cmd/prepare"
	"github.com/tinygodsdev/tinycooksweb/internal/app"
	"github.com/tinygodsdev/tinycooksweb/internal/config"
	"github.com/tinygodsdev/tinycooksweb/internal/handler"
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

	store := prepare.Store(cfg)

	a, err := app.New(cfg, logger)
	if err != nil {
		logger.Fatal("failed to init app", "err", err)
	}

	h := handler.NewHandler(a, logger, "templates/")
	r := prepare.Mux(cfg, store, h)

	httpServer := prepare.Server(cfg, r)
	httpServer.Addr = cfg.GetAddress()
	logger.Info("starting http server", cfg.GetAddress())
	log.Fatal(httpServer.ListenAndServe())

	logger.Info("Starting application.", "env", cfg.Env)
}
