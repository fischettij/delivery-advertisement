package storage_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fischettij/delivery-advertisement/internal/storage"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"testing"
)

type PostgresSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
}

func TestPostgresSuite(t *testing.T) {
	suite.Run(t, new(PostgresSuite))
}

func (suite *PostgresSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
}

func (suite *PostgresSuite) TearDownTest() {
}

func (suite *PostgresSuite) TestNewPostgresDatabase() {
	suite.Run("given_a_db_when_creates_new_postgres_then_return_postgres_and_no_error", func() {
		db, _, err := sqlmock.New()
		suite.Require().NoError(err)

		manager, err := storage.NewPostgresDatabase(zap.NewNop(), db)
		suite.Require().NoError(err)
		suite.Require().NotNil(manager)
	})

	suite.Run("given_a_nil_db_when_creates_new_postgres_then_return_error", func() {
		manager, err := storage.NewPostgresDatabase(zap.NewNop(), nil)
		suite.Require().Error(err)
		suite.Require().ErrorContains(err, "db cannot be nil")
		suite.Require().Nil(manager)
	})

	suite.Run("given_a_nil_logger_when_creates_new_postgres_then_return_error", func() {
		db, _, err := sqlmock.New()
		suite.Require().NoError(err)

		manager, err := storage.NewPostgresDatabase(nil, db)
		suite.Require().Error(err)
		suite.Require().ErrorContains(err, "logger cannot be nil")
		suite.Require().Nil(manager)
	})
}
