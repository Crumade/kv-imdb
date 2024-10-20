package main

import (
	"kv-imdb/internal/compute"
	"kv-imdb/internal/database"
	"kv-imdb/internal/storage"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Fatal("logger init error", zap.Error(err))
	}

	parser := compute.NewCompute(logger)
	storage := storage.NewStorage()

	db := database.NewDatabase(logger, parser, storage)
	logger.Debug("application starting...")
	db.Start()
}
