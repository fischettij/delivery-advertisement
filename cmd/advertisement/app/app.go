package app

import (
	"log"

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

	config := LoadConfig()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	dbstorage, err := storage.NewMemoryStorage(logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	restClient := resty.New()

	fileDownloader, err := downloader.NewDownloader(config.FileDownloaderResourceURL, restClient)
	if err != nil {
		logger.Fatal(err.Error())
	}

	manager, err := deliverys.NewManager(dbstorage, fileDownloader)
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = manager.Start()
	if err != nil {
		logger.Fatal(err.Error())
	}

	deliveryServicesHandler, err := handlers.NewDeliveryServicesHandler(manager)
	if err != nil {
		logger.Fatal(err.Error())
	}
	handlers.ConfigureRoutes(r, deliveryServicesHandler)

	err = r.Run()
	if err != nil {
		logger.Fatal(err.Error())
	}
}
