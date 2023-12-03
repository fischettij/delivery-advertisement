package deliverys_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/fischettij/delivery-advertisement/pkg/deliverys"
	"github.com/fischettij/delivery-advertisement/pkg/deliverys/mocks"
)

//go:generate mockgen -destination=mocks/mocks.go -package=mocks -source=manager.go

type ManagerSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
}

func TestHandlersSuite(t *testing.T) {
	suite.Run(t, new(ManagerSuite))
}

func (suite *ManagerSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
}

func (suite *ManagerSuite) TearDownTest() {
}

func (suite *ManagerSuite) TestNew() {
	suite.Run("given_a_storage_when_creates_new_manager_then_return_manager_and_no_error", func() {
		manager, err := deliverys.NewManager(mocks.NewMockStorage(suite.mockCtrl), mocks.NewMockFileDownloader(suite.mockCtrl))
		suite.Require().NoError(err)
		suite.Require().NotNil(manager)
	})

	suite.Run("given_a_nil_storage_when_creates_new_manager_then_return_error", func() {
		manager, err := deliverys.NewManager(nil, mocks.NewMockFileDownloader(suite.mockCtrl))
		suite.Require().Error(err)
		suite.Require().ErrorContains(err, "storage can not be nil")
		suite.Require().Nil(manager)
	})

	suite.Run("given_a_nil_storage_when_creates_new_manager_then_return_error", func() {
		manager, err := deliverys.NewManager(mocks.NewMockStorage(suite.mockCtrl), nil)
		suite.Require().Error(err)
		suite.Require().ErrorContains(err, "downloader can not be nil")
		suite.Require().Nil(manager)
	})
}
