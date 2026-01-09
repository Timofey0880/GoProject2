package event_service_api

import (
	"context"
	"log"

	"github.com/Timofey0880/GoProject2/event-service/internal/models"
	"github.com/Timofey0880/GoProject2/event-service/internal/pb/event_api"
)

func (s *EventServiceAPI) CreateEvent(ctx context.Context, req *event_api.CreateEventRequest) (*event_api.CreateEventResponse, error) {
	log.Printf("Received create event: %v", req)

	event := &models.Event{Name: req.Name, City: req.City, User: req.User}
	id, err := s.eventService.CreateEvent(ctx, event)
	if err != nil {
		return &event_api.CreateEventResponse{}, err
	}
	return &event_api.CreateEventResponse{EventId: id}, nil
}
