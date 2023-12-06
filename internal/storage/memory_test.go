package storage_test

import (
	"context"
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

func (suite *MemorySuite) TestDeliveryServicesNearLocation() {
	suite.Run("given_a_establishments_repository_when_search_deliveries_services_then_return_a_filtered_slice_of_ids", func() {
		memory, err := storage.NewMemoryStorage(zap.NewNop())
		suite.Require().NoError(err)
		suite.loadData(memory)

		deliveries, err := memory.DeliveryServicesNearLocation(context.Background(), 50.45, 3.94)
		expectedID := "14"
		suite.Require().NoError(err)
		suite.Require().Equal(1, len(deliveries))
		suite.Require().Equal(expectedID, deliveries[0])
	})
}

func (suite *MemorySuite) loadData(memoryStorage *storage.Memory) {
	content := []string{
		"id,latitude,longitude,availability_radius,open_hour,close_hour,rating",
		"1,51.194253600000003,6.455508,5,14:00:00,23:00:00,4.7",
		"2,50.132921000000003,19.638506100000001,2,14:00:00,23:00:00,4.8",
		"3,52.501866800000002,13.3254556,3,09:00:00,20:00:00,4.0",
		"4,50.911882200000001,4.4350510999999999,1,08:00:00,23:00:00,4.9",
		"5,46.994769099999999,8.6082292999999996,1,12:00:00,23:00:00,4.7",
		"6,50.835918300000003,4.3066259999999996,1,09:00:00,20:00:00,4.1",
		"7,52.0504818,4.2724795000000002,1,14:00:00,23:00:00,4.5",
		"8,50.834134599999999,4.3451088999999996,3,12:00:00,23:00:00,3.6",
		"9,48.399190900000001,10.8844431,1,14:00:00,23:00:00,4.3",
		"10,50.079341900000003,8.2053432999999991,2,14:00:00,23:00:00,4.0",
		"11,50.042026700000001,20.000539199999999,3,14:00:00,23:00:00,4.8",
		"12,47.391980199999999,8.0455124999999992,3,14:00:00,23:00:00,4.3",
		"13,52.416715699999997,4.6474712,2,14:00:00,23:00:00,3.6",
		"14,50.454039700000003,3.9527196,2,14:00:00,23:00:00,3.9",
		"15,52.504231799999999,13.307130900000001,3,12:00:00,23:00:00,4.8",
		"16,53.250390000000003,10.410920000000001,4,09:00:00,20:00:00,4.4",
		"17,50.082827299999998,8.2375963999999993,2,14:00:00,23:00:00,4.6",
		"18,49.323677000000004,8.4285720000000008,1,10:00:00,18:00:00,4.9",
		"19,50.861949000000003,5.6249083999999998,3,10:00:00,18:00:00,3.8",
		"20,54.312081999999997,10.131498000000001,5,12:00:00,23:00:00,4.4",
		"21,53.117998,23.1510839,3,14:00:00,23:00:00,4.0",
		"22,50.053419699999999,8.6705214000000002,3,12:00:00,23:00:00,3.6",
		"23,53.142735500000001,7.0393866000000003,5,12:00:00,23:00:00,4.6",
		"24,50.944586600000001,3.1274829,4,10:00:00,18:00:00,4.5",
		"25,51.905175499999999,10.4278589,2,09:00:00,20:00:00,3.9",
		"26,52.396435599999997,16.955506799999998,1,10:00:00,18:00:00,5.0",
		"27,51.924979899999997,10.431238799999999,4,08:00:00,23:00:00,4.5",
		"28,51.794403099999997,11.744137800000001,3,14:00:00,23:00:00,4.7",
		"29,52.401077000000001,16.92812,5,10:00:00,18:00:00,3.5",
		"30,51.56138,5.0825480000000001,2,09:00:00,20:00:00,4.4",
	}
	// Crear o truncar el archivo
	fileName := "test_file.csv"
	file, err := os.Create(fileName)
	suite.Require().NoError(err)
	defer file.Close()
	defer os.Remove(fileName)

	for _, line := range content {
		_, err = fmt.Fprintln(file, line)
		suite.Require().NoError(err)
	}

	err = memoryStorage.LoadFromFile(fileName)
	suite.Require().NoError(err)

}
