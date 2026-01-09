package event_service_api

import (
	"context"

	"github.com/Timofey0880/GoProject2/event-service/internal/models"
	"github.com/Timofey0880/GoProject2/event-service/internal/pb/event_api"
)

type eventService interface {
	CreateEvent(ctx context.Context, event *models.Event) (uint64, error)
	GetEvent(ctx context.Context, id uint64) (*models.Event, error)
}

type EventServiceAPI struct {
	event_api.UnimplementedEventServiceServer
	eventService eventService
}

func NewEventServiceAPI(eventService eventService) *EventServiceAPI {
	return &EventServiceAPI{
		eventService: eventService,
	}
}
