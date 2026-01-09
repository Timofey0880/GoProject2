package pgstorage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type PGStorage struct {
	db *pgxpool.Pool
}

func NewPGStorage(connString string) (*PGStorage, error) {

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, errors.Wrap(err, "parse config error")
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, errors.Wrap(err, "connect error")
	}
	storage := &PGStorage{
		db: db,
	}
	err = storage.initTables()
	if err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *PGStorage) initTables() error {
	sql := fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %v (
        %v SERIAL PRIMARY KEY,
        %v VARCHAR(100) NOT NULL,
        %v VARCHAR(100) NOT NULL,
        %v VARCHAR(100) NOT NULL
    )`, tableName, idColumnName, nameColumnName, cityColumnName, userColumnName)
	_, err := s.db.Exec(context.Background(), sql)
	if err != nil {
		return errors.Wrap(err, "init tables")
	}
	return nil
}
