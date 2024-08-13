package main

import (
	"log"
	"time"

	"github.com/tinygodsdev/datasdk/pkg/logger"
	"github.com/tinygodsdev/tinycooksweb/internal/app"
	"github.com/tinygodsdev/tinycooksweb/internal/config"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage"
	"github.com/tinygodsdev/tinycooksweb/pkg/storage/entstorage"
)

func main() {
	log.Println("Starting seed data save")
	start := time.Now()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logger := logger.NewStdLogger()

	if !cfg.Dev() {
		logger.Fatal("seed data can only be saved in dev mode")
	}

	recipeStore, err := entstorage.NewEntStorage(entstorage.Config{
		StorageDriver: cfg.StorageDriver,
		StorageDSN:    cfg.StorageDSN,
		LogQueries:    cfg.LogDBQueries,
		Migrate:       true,
	}, logger)
	if err != nil {
		logger.Fatal("failed to init recipe store", "err", err)
	}
	defer recipeStore.Close()
	var store storage.Storage = recipeStore

	logger.Info("store initialized")

	cfg.CreateJobs = false
	a, err := app.New(cfg, logger, store, nil)
	if err != nil {
		logger.Fatal("failed to init app", "err", err)
	}
	logger.Info("app initialized")

	err = a.SaveSeedData()
	if err != nil {
		logger.Fatal("failed to save seed data", "err", err)
	}

	log.Println("Seed data saved in", time.Since(start))
}
