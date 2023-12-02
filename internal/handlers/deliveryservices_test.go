package handlers_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fischettij/delivery-advertisement/internal/handlers"
	"github.com/fischettij/delivery-advertisement/internal/handlers/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks -source=deliveryservices.go

type DeliveryServicesHandlerSuite struct {
	suite.Suite
	mockCtrl               *gomock.Controller
	deliveryServiceManager *mocks.MockDeliveryServiceManager
	handler                *handlers.DeliveryServicesHandler
}

func TestHandlersSuite(t *testing.T) {
	suite.Run(t, new(DeliveryServicesHandlerSuite))
}

func (suite *DeliveryServicesHandlerSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.deliveryServiceManager = mocks.NewMockDeliveryServiceManager(suite.mockCtrl)
	var err error
	suite.handler, err = handlers.NewDeliveryServicesHandler(suite.deliveryServiceManager)
	suite.Require().NoError(err)
	gin.SetMode(gin.TestMode)
}

func (suite *DeliveryServicesHandlerSuite) TearDownTest() {
}

func (suite *DeliveryServicesHandlerSuite) TestFindNearLocation() {
	router := gin.New()
	endpoint := "/delivery-services"
	router.GET(endpoint, suite.handler.FindNearLocation)

	suite.Run("given_a_coordinates_when_manager_return_return_ids_then_got_response_with_that_ids", func() {
		expectedIDs := []string{"1"}
		latitude := 25.420861
		longitude := 51.490388

		suite.deliveryServiceManager.EXPECT().DeliveryServicesNearLocation(gomock.Any(), latitude, longitude).Return(expectedIDs, nil)

		queryParameters := fmt.Sprintf("?lat=%f&lon=%f", latitude, longitude)
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", endpoint, queryParameters), nil)
		suite.Require().NoError(err)

		responseRecorder := httptest.NewRecorder()
		router.ServeHTTP(responseRecorder, req)

		suite.Require().Equal(http.StatusOK, responseRecorder.Code)

		type responseBody struct {
			IDs []string `json:"ids"`
		}
		var gotBody *responseBody
		err = json.Unmarshal(responseRecorder.Body.Bytes(), &gotBody)
		suite.Require().NoError(err)
		expectedBody := &responseBody{IDs: expectedIDs}
		suite.Require().Equal(expectedBody, gotBody)
	})

	suite.Run("given_a_coordinates_when_manager_return_return_no_ids_then_got_an_empty_response", func() {
		expectedIDs := []string{}
		latitude := 25.420861
		longitude := 51.490388

		suite.deliveryServiceManager.EXPECT().DeliveryServicesNearLocation(gomock.Any(), latitude, longitude).Return(expectedIDs, nil)

		queryParameters := fmt.Sprintf("?lat=%f&lon=%f", latitude, longitude)
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", endpoint, queryParameters), nil)
		suite.Require().NoError(err)

		responseRecorder := httptest.NewRecorder()
		router.ServeHTTP(responseRecorder, req)

		suite.Require().Equal(http.StatusOK, responseRecorder.Code)

		type responseBody struct {
			IDs []string `json:"ids"`
		}
		var gotBody *responseBody
		err = json.Unmarshal(responseRecorder.Body.Bytes(), &gotBody)
		suite.Require().NoError(err)
		expectedBody := &responseBody{IDs: expectedIDs}
		suite.Require().Equal(expectedBody, gotBody)
	})

	suite.Run("given_bad_formatted_latitude_when_parse_return_an_error_then_got_bad_request_response", func() {
		queryParameters := fmt.Sprintf("?lat=%s&lon=%f", "BADFORMAT", 10.10)
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", endpoint, queryParameters), nil)
		suite.Require().NoError(err)

		responseRecorder := httptest.NewRecorder()
		router.ServeHTTP(responseRecorder, req)

		suite.Require().Equal(http.StatusBadRequest, responseRecorder.Code)
	})

	suite.Run("given_bad_formatted_longitude_when_parse_return_an_error_then_got_bad_request_response", func() {
		queryParameters := fmt.Sprintf("?lat=%f&lon=%s", 10.10, "BADFORMAT")
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", endpoint, queryParameters), nil)
		suite.Require().NoError(err)

		responseRecorder := httptest.NewRecorder()
		router.ServeHTTP(responseRecorder, req)

		suite.Require().Equal(http.StatusBadRequest, responseRecorder.Code)
	})

	suite.Run("given_a_coordinate_when_manager_return_an_error_then_got_internal_server_error", func() {
		latitude := 25.420861
		longitude := 51.490388

		suite.deliveryServiceManager.EXPECT().DeliveryServicesNearLocation(gomock.Any(), latitude, longitude).Return([]string{}, errors.New("some-error"))

		queryParameters := fmt.Sprintf("?lat=%f&lon=%f", latitude, longitude)
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", endpoint, queryParameters), nil)
		suite.Require().NoError(err)

		responseRecorder := httptest.NewRecorder()
		router.ServeHTTP(responseRecorder, req)

		suite.Require().Equal(http.StatusInternalServerError, responseRecorder.Code)
	})

}
