package pgstorage

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/Timofey0880/GoProject2/event-service/internal/models"
	"github.com/pkg/errors"
)

func (storage *PGStorage) CreateEvent(ctx context.Context, event *models.Event) (uint64, error) {
	query := storage.createQuery(event)
	queryText, args, err := query.ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "generate query error")
	}
	var id uint64
	err = storage.db.QueryRow(ctx, queryText, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "exec query")
	}
	return id, nil
}

func (storage *PGStorage) createQuery(event *models.Event) squirrel.Sqlizer {
	q := squirrel.Insert(tableName).Columns(nameColumnName, cityColumnName, userColumnName).
		Values(event.Name, event.City, event.User).
		Suffix("RETURNING " + idColumnName).
		PlaceholderFormat(squirrel.Dollar)
	return q
}
