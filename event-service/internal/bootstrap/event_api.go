package bootstrap

import (
	api "github.com/Timofey0880/GoProject2/event-service/internal/api/event_service_api"
	"github.com/Timofey0880/GoProject2/event-service/internal/producers/event_created_producer"
	"github.com/Timofey0880/GoProject2/event-service/internal/services/eventService"
)

func InitEventServiceAPI(eventService *eventService.EventService, producer *event_created_producer.EventCreatedProducer) *api.EventServiceAPI {
	return api.NewEventServiceAPI(eventService) // Producer inject in service if needed, but in usecase.
}
