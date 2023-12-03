package storage_test

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/fischettij/delivery-advertisement/internal/storage"
)

type MemorySuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
}

func TestMemorySuite(t *testing.T) {
	suite.Run(t, new(MemorySuite))
}

func (suite *MemorySuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
}

func (suite *MemorySuite) TearDownTest() {
}

func (suite *MemorySuite) TestLoadFromFile() {

	suite.Run("given_a_file_path_with_expected_format_when_load_it_then_return_no_error", func() {
		fileName := "testfile.csv"
		tempFile, err := os.Create(fileName)
		suite.Require().NoError(err)
		defer os.Remove(tempFile.Name())

		header := "id,latitude,longitude,availability_radius,open_hour,close_hour,rating"
		data := "1,51.194253600000003,6.455508,5,14:00:00,23:00:00,4.7"
		testData := []byte(fmt.Sprintf("%s\n%s\n", header, data))
		_, err = tempFile.Write(testData)
		suite.Require().NoError(err)
		tempFile.Close()

		memory, err := storage.NewMemoryStorage(zap.NewNop())
		suite.Require().NoError(err)

		err = memory.LoadFromFile(fileName)
		suite.Require().NoError(err)
	})

	suite.Run("given_a_file_path_with_unexpected_headers_when_load_it_then_return_error", func() {
		fileName := "testfile.csv"
		tempFile, err := os.Create(fileName)
		suite.Require().NoError(err)
		defer os.Remove(tempFile.Name())

		header := "id,latitude,UNEXPECTED,availability_radius,open_hour,close_hour,rating"
		data := "1,51.194253600000003,6.455508,5,14:00:00,23:00:00,4.7"
		testData := []byte(fmt.Sprintf("%s\n%s\n", header, data))
		_, err = tempFile.Write(testData)
		suite.Require().NoError(err)
		tempFile.Close()

		memory, err := storage.NewMemoryStorage(zap.NewNop())
		suite.Require().NoError(err)

		err = memory.LoadFromFile(fileName)
		suite.Require().ErrorIs(err, storage.ErrUnexpectedHeaders)
	})

	suite.Run("given_a_file_path_with_invalid_latitude_when_load_it_then_skip_record", func() {
		fileName := "testfile.csv"
		tempFile, err := os.Create(fileName)
		suite.Require().NoError(err)
		defer os.Remove(tempFile.Name())

		header := "id,latitude,longitude,availability_radius,open_hour,close_hour,rating"
		data := "1,INVALID,6.455508,5,14:00:00,23:00:00,4.7"
		testData := []byte(fmt.Sprintf("%s\n%s\n", header, data))
		_, err = tempFile.Write(testData)
		suite.Require().NoError(err)
		tempFile.Close()

		memory, err := storage.NewMemoryStorage(zap.NewNop())
		suite.Require().NoError(err)

		err = memory.LoadFromFile(fileName)
		suite.Require().NoError(err)
	})

	suite.Run("given_a_file_path_with_invalid_longitude_when_load_it_then_skip_record", func() {
		fileName := "testfile.csv"
		tempFile, err := os.Create(fileName)
		suite.Require().NoError(err)
		defer os.Remove(tempFile.Name())

		header := "id,latitude,longitude,availability_radius,open_hour,close_hour,rating"
		data := "1,6.455508,INVALID,5,14:00:00,23:00:00,4.7"
		testData := []byte(fmt.Sprintf("%s\n%s\n", header, data))
		_, err = tempFile.Write(testData)
		suite.Require().NoError(err)
		tempFile.Close()

		memory, err := storage.NewMemoryStorage(zap.NewNop())
		suite.Require().NoError(err)

		err = memory.LoadFromFile(fileName)
		suite.Require().NoError(err)
	})
}
