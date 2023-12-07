package app

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/fischettij/delivery-advertisement/internal/downloader"
	"github.com/fischettij/delivery-advertisement/internal/handlers"
	"github.com/fischettij/delivery-advertisement/internal/storage"
	"github.com/fischettij/delivery-advertisement/pkg/deliverys"
)

func Start() {
	r := gin.Default()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	config, err := LoadConfig()
	if err != nil {
		logger.Fatal(err.Error())
	}

	dbstorage, err := storage.NewMemoryStorage(logger)
	if err != nil {
		logger.Fatal(err.Error())
	}
	restClient := resty.New()

	fileDownloader, err := downloader.NewDownloader(config.FileDownloaderResourceURL, restClient)
	if err != nil {
		logger.Fatal(err.Error())
	}

	manager, err := deliverys.NewManager(dbstorage, fileDownloader, time.Duration(config.FilePollingIntervalMinutes)*time.Minute)
	if err != nil {
		logger.Fatal(err.Error())
	}
	done := make(chan error)
	manager.Start(done)

	deliveryServicesHandler, err := handlers.NewDeliveryServicesHandler(manager)
	if err != nil {
		logger.Fatal(err.Error())
	}
	handlers.ConfigureRoutes(r, deliveryServicesHandler)

	err = r.Run(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info(fmt.Sprintf("application started and listening on port %s", config.Port))

	select {
	case err = <-done:
		logger.Fatal(err.Error())
	}
}
