package bootstrap

import (
	"context"

	"github.com/Timofey0880/GoProject2/event-service/config"
	"github.com/Timofey0880/GoProject2/event-service/internal/producers/event_created_producer"
	"github.com/Timofey0880/GoProject2/event-service/internal/services/eventService"
	"github.com/Timofey0880/GoProject2/event-service/internal/storage/pgstorage"
)

func InitEventService(storage *pgstorage.PGstorage, producer *event_created_producer.EventCreatedProducer, cfg *config.Config) *eventService.EventService {

	return eventService.NewEventService(context.Background(), storage, producer, cfg.EventServiceSettings.MinNameLen, cfg.EventServiceSettings.MaxNameLen, cfg.EventServiceSettings.MinCityLen, cfg.EventServiceSettings.MaxCityLen)
}
