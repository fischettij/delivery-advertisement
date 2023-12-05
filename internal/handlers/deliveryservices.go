package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeliveryServiceManager interface {
	DeliveryServicesNearLocation(ctx context.Context, latitude, longitude float64) ([]string, error)
}

type deliveryServicesResponse struct {
	IDs []string `json:"ids"`
}

type DeliveryServicesHandler struct {
	manager DeliveryServiceManager
}

func NewDeliveryServicesHandler(manager DeliveryServiceManager) (*DeliveryServicesHandler, error) {
	if manager == nil {
		return nil, errors.New("delivery service manager can not be nil")
	}
	return &DeliveryServicesHandler{manager: manager}, nil
}

func (d DeliveryServicesHandler) FindNearLocation(c *gin.Context) {
	latitude, longitude, err := d.parseCoordinatesFromQuery(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	deliveryServicesIDs, err := d.manager.DeliveryServicesNearLocation(context.Background(), latitude, longitude)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, deliveryServicesResponse{IDs: deliveryServicesIDs})
}

func (d DeliveryServicesHandler) parseCoordinatesFromQuery(c *gin.Context) (float64, float64, error) {
	latitudeStr := c.Query("lat")
	longitudeStr := c.Query("lon")
	latitude, err := parseDegree(latitudeStr)
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing latitude: %w", err)
	}

	longitude, err := parseDegree(longitudeStr)
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing longitude: %w", err)
	}
	return latitude, longitude, nil
}

func parseDegree(latitudeStr string) (float64, error) {
	// The bit size 64 good choice because it provides higher precision
	return strconv.ParseFloat(latitudeStr, 64)
}
