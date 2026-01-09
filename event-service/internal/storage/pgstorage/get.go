package pgstorage

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/Timofey0880/GoProject2/event-service/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (storage *PGStorage) GetEvent(ctx context.Context, id uint64) (*models.Event, error) {
	query := storage.getQuery(id)
	queryText, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "generate query error")
	}
	event := &models.Event{}
	err = storage.db.QueryRow(ctx, queryText, args...).Scan(&event.ID, &event.Name, &event.City, &event.User)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "querying error")
	}
	return event, nil
}

func (storage *PGStorage) getQuery(id uint64) squirrel.Sqlizer {
	q := squirrel.Select(idColumnName, nameColumnName, cityColumnName, userColumnName).From(tableName).
		Where(squirrel.Eq{idColumnName: id}).PlaceholderFormat(squirrel.Dollar)
	return q
}
