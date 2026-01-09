package eventService

import (
	"context"

	"github.com/Timofey0880/GoProject2/event-service/internal/models"
	"github.com/pkg/errors"
)

func (s *EventService) CreateEvent(ctx context.Context, event *models.Event) (uint64, error) {
	if err := s.validateEvent(event); err != nil {
		return 0, err
	}
	id, err := s.eventStorage.CreateEvent(ctx, event)
	if err != nil {
		return 0, err
	}
	event.ID = id
	err = s.eventProducer.ProduceEventCreated(ctx, event)
	if err != nil {
		return 0, errors.Wrap(err, "produce event created")
	}
	return id, nil
}
