package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

const establishmentsTableName = "establishments"

// Postgres contains resources to interact with a postgres db.
type Postgres struct {
	db      *sql.DB
	builder sq.StatementBuilderType
	logger  Logger
}

// NewPostgresDatabase creates a new postgres database.
func NewPostgresDatabase(logger Logger, db *sql.DB) (*Postgres, error) {
	if logger == nil {
		return nil, errors.New("logger cannot be nil")
	}
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}

	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(db)

	return &Postgres{
		db:      db,
		builder: builder,
		logger:  logger,
	}, nil
}

func (p *Postgres) LoadFromFile(path string) error {
	_, err := p.db.Exec("TRUNCATE TABLE $1", establishmentsTableName)
	if err != nil {
		return fmt.Errorf("error executing insert establishment: %w", err)
	}

	p.logger.Info("db population from file started")
	err = loadFromFile(p.logger, path, func(establishment *Establishment) error {
		builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(p.db)
		_, err := builder.Insert(establishmentsTableName).
			Columns("establishment_id", "latitude", "longitude", "availability_radius").
			Values(establishment.ID, establishment.Latitude, establishment.Longitude, establishment.AvailabilityRadios).
			Exec()
		if err != nil {
			return fmt.Errorf("error executing insert establishment: %w", err)
		}
		return nil
	})

	p.logger.Info("db population finished")
	return err
}

func (p *Postgres) DeliveryServicesNearLocation(_ context.Context, latitude, longitude float64) ([]string, error) {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).RunWith(p.db)

	// Select the locations approximately 20 km away to analyze a smaller subset.
	rows, err := builder.Select("establishment_id", "latitude", "longitude", "availability_radius").
		From(establishmentsTableName).
		Where(sq.And{
			sq.Expr("\"latitude\" BETWEEN $1 AND $2", latitude-0.25, latitude+0.25),
			sq.Expr("\"longitude\" BETWEEN $3 AND $4", longitude-0.25, longitude+0.25),
		}).
		Query()
	defer rows.Close()

	if err != nil {
		return []string{}, fmt.Errorf("error in query select of DeliveryServicesNearLocation: %w", err)
	}

	var result []string
	for rows.Next() {
		var id string
		var establismentLatitude, establismentLongitude, availabilityRadios float64
		if err := rows.Scan(&id, &establismentLatitude, &establismentLongitude, &availabilityRadios); err != nil {
			return []string{}, fmt.Errorf("error scanning query result: %w", err)
		}

		if IsInRangeToRadius(&Establishment{
			ID:                 id,
			Latitude:           establismentLatitude,
			Longitude:          establismentLongitude,
			AvailabilityRadios: availabilityRadios,
		}, latitude, longitude) {
			result = append(result, id)
		}
	}

	return result, nil
}
