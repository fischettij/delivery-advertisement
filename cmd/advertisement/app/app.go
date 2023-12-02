package app

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/fischettij/delivery-advertisement/internal/handlers"
)

func Start() {
	r := gin.Default()

	deliveryServicesHandler, err := handlers.NewDeliveryServicesHandler(nil)
	if err != nil {
		log.Fatal(err)
	}
	handlers.ConfigureRoutes(r, deliveryServicesHandler)

	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
