package eventService

import (
	"context"

	"github.com/Timofey0880/GoProject2/event-service/internal/models"
)

type EventStorage interface {
	CreateEvent(ctx context.Context, event *models.Event) (uint64, error)
	GetEvent(ctx context.Context, id uint64) (*models.Event, error)
}

type EventProducer interface {
	ProduceEventCreated(ctx context.Context, event *models.Event) error
}

type EventService struct {
	eventStorage  EventStorage
	eventProducer EventProducer
	minNameLen    int
	maxNameLen    int
	minCityLen    int
	maxCityLen    int
}

func NewEventService(ctx context.Context, eventStorage EventStorage, eventProducer EventProducer, minNameLen, maxNameLen, minCityLen, maxCityLen int) *EventService {
	return &EventService{
		eventStorage:  eventStorage,
		eventProducer: eventProducer,
		minNameLen:    minNameLen,
		maxNameLen:    maxNameLen,
		minCityLen:    minCityLen,
		maxCityLen:    maxCityLen,
	}
}
