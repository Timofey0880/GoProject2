package event_service_api

import (
	"context"
	"log"

	"github.com/Timofey0880/GoProject2/event-service/internal/pb/event_api"
	"github.com/pkg/errors"
)

func (s *EventServiceAPI) GetEvent(ctx context.Context, req *event_api.GetEventRequest) (*event_api.GetEventResponse, error) {
	log.Printf("Received get event id: %v", req.EventId)

	event, err := s.eventService.GetEvent(ctx, req.EventId)
	if err != nil {
		return &event_api.GetEventResponse{}, err
	}
	if event == nil {
		return &event_api.GetEventResponse{}, errors.New("event not found")
	}
	return &event_api.GetEventResponse{Name: event.Name, City: event.City, User: event.User}, nil
}
