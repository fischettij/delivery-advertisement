package downloader_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"

	"github.com/fischettij/delivery-advertisement/internal/downloader"
)

type DownloaderSuite struct {
	suite.Suite
}

func TestDownloaderSuite(t *testing.T) {
	suite.Run(t, new(DownloaderSuite))
}

func (suite *DownloaderSuite) SetupTest() {
}

func (suite *DownloaderSuite) TearDownTest() {
}

func (suite *DownloaderSuite) TestNewDownloader() {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())

	suite.Run("given_an_url_and_resty_client_when_create_new_downloader_then_return_downloader_and_no_error", func() {
		url := "localhost:8080.file.csv"
		dw, err := downloader.NewDownloader(url, client)
		suite.Require().NoError(err)
		suite.Require().NotNil(dw)
	})

	suite.Run("given_an_empty_url_and_resty_client_when_create_new_downloader_then_return_error", func() {
		url := ""
		dw, err := downloader.NewDownloader(url, client)
		suite.Require().Error(err)
		suite.Require().ErrorContains(err, "resource url cannot be empty string")
		suite.Require().Nil(dw)
	})

	suite.Run("given_an_url_and_nil_client_when_create_new_downloader_then_return_error", func() {
		url := "localhost:8080.file.csv"
		dw, err := downloader.NewDownloader(url, nil)
		suite.Require().Error(err)
		suite.Require().ErrorContains(err, "resty client can not be nil")
		suite.Require().Nil(dw)
	})
}

func (suite *DownloaderSuite) TestLoadFromFile() {
	suite.Run("given_a_file_path_with_expected_format_when_load_it_then_return_no_error", func() {
		client := resty.New()
		httpmock.ActivateNonDefault(client.GetClient())
		url := "localhost:8080.file.csv"
		testFilePath := "test_file.csv"
		defer os.Remove(testFilePath)

		csvheader := "id,latitude,longitude,availability_radius,open_hour,close_hour,rating"
		csvcontent := "1,51.194253600000003,6.455508,5,14:00:00,23:00:00,4.7"
		httpmock.RegisterResponder("GET", url,
			func(request *http.Request) (*http.Response, error) {
				body := fmt.Sprintf("%s\n%s", csvheader, csvcontent)
				return httpmock.NewStringResponse(200, body), nil
			})

		dw, err := downloader.NewDownloader(url, client)
		suite.Require().NoError(err)
		suite.Require().NotNil(dw)

		err = dw.DownloadFile(testFilePath)
		suite.Require().NoError(err)

		// Validate file content
		downloadedFile, err := os.ReadFile(testFilePath)
		suite.Require().NoError(err)
		fileContent := string(downloadedFile)
		suite.Require().Equal(fileContent, fmt.Sprintf("%s\n%s", csvheader, csvcontent))
	})
}
