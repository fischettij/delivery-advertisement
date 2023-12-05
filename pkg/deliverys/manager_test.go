package deliverys_test

import (
	"errors"
	"testing"
	"time"

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
	filePollingInterval := 1 * time.Second
	suite.Run("given_a_storage_when_creates_new_manager_then_return_manager_and_no_error", func() {
		manager, err := deliverys.NewManager(mocks.NewMockStorage(suite.mockCtrl), mocks.NewMockFileDownloader(suite.mockCtrl), filePollingInterval)
		suite.Require().NoError(err)
		suite.Require().NotNil(manager)
	})

	suite.Run("given_a_nil_storage_when_creates_new_manager_then_return_error", func() {
		manager, err := deliverys.NewManager(nil, mocks.NewMockFileDownloader(suite.mockCtrl), filePollingInterval)
		suite.Require().Error(err)
		suite.Require().ErrorContains(err, "storage can not be nil")
		suite.Require().Nil(manager)
	})

	suite.Run("given_a_nil_storage_when_creates_new_manager_then_return_error", func() {
		manager, err := deliverys.NewManager(mocks.NewMockStorage(suite.mockCtrl), nil, filePollingInterval)
		suite.Require().Error(err)
		suite.Require().ErrorContains(err, "downloader can not be nil")
		suite.Require().Nil(manager)
	})
}

func (suite *ManagerSuite) TestStart() {
	filePollingInterval := 2 * time.Second

	suite.Run("given_a_two_seconds_interval_when_start_manager_then_download_file_two_times_in_four_seconds", func() {
		mockStorage := mocks.NewMockStorage(suite.mockCtrl)
		mockFileDownloader := mocks.NewMockFileDownloader(suite.mockCtrl)
		manager, err := deliverys.NewManager(mockStorage, mockFileDownloader, filePollingInterval)
		suite.Require().NoError(err)
		suite.Require().NotNil(manager)

		tcChan := make(chan bool)
		mockFileDownloader.EXPECT().DownloadFile(gomock.Any()).DoAndReturn(func(_ string) (string, error) {
			tcChan <- true
			return "someMD5", nil
		}).AnyTimes()
		mockFileDownloader.EXPECT().RemoveFile(gomock.Any()).Return(nil).AnyTimes()
		mockStorage.EXPECT().LoadFromFile(gomock.Any()).Return(nil).AnyTimes()

		done := make(chan error)
		manager.Start(done)

		pollingTimesLeft := 2
		timeout := time.After(filePollingInterval * 2)
		// Expect two call to DownloadFile
		for pollingTimesLeft > 0 {
			select {
			case <-timeout:
				suite.Fail("test didn't finish in time")
			case <-tcChan:
				pollingTimesLeft--
			}
		}
	})

	suite.Run("when_file_downloader_returns_error_then_done_channel_receive_an_error", func() {
		mockStorage := mocks.NewMockStorage(suite.mockCtrl)
		mockFileDownloader := mocks.NewMockFileDownloader(suite.mockCtrl)
		manager, err := deliverys.NewManager(mockStorage, mockFileDownloader, filePollingInterval)
		suite.Require().NoError(err)
		suite.Require().NotNil(manager)
		mockFileDownloader.EXPECT().DownloadFile(gomock.Any()).Return("", errors.New("some-error")).Times(1)
		mockFileDownloader.EXPECT().RemoveFile(gomock.Any()).Return(nil).AnyTimes()

		done := make(chan error)
		manager.Start(done)

		timeout := time.After(filePollingInterval * 2)
		select {
		case <-timeout:
			suite.Fail("test didn't finish in time")
		case err = <-done:
			suite.Require().Error(err)
		}
	})

	suite.Run("when_file_downloader_returns_error_then_done_channel_receive_an_error", func() {
		mockStorage := mocks.NewMockStorage(suite.mockCtrl)
		mockFileDownloader := mocks.NewMockFileDownloader(suite.mockCtrl)
		manager, err := deliverys.NewManager(mockStorage, mockFileDownloader, filePollingInterval)
		suite.Require().NoError(err)
		suite.Require().NotNil(manager)
		mockFileDownloader.EXPECT().DownloadFile(gomock.Any()).Return("md5", nil).Times(1)
		mockFileDownloader.EXPECT().RemoveFile(gomock.Any()).Return(nil).AnyTimes()
		mockStorage.EXPECT().LoadFromFile(gomock.Any()).Return(errors.New("some-error")).Times(1)

		done := make(chan error)
		manager.Start(done)

		timeout := time.After(filePollingInterval * 2)
		select {
		case <-timeout:
			suite.Fail("test didn't finish in time")
		case err = <-done:
			suite.Require().Error(err)
		}
	})
}
