package handlers

import (
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine, deliveryServicesHandler *DeliveryServicesHandler) {
	router.GET("/delivery-services", deliveryServicesHandler.FindNearLocation)
}
