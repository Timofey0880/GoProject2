package main

import (
	"fmt"
	"os"

	"github.com/Timofey0880/GoProject2/event-service/config"
	"github.com/Timofey0880/GoProject2/event-service/internal/bootstrap"
)

func main() {

	cfg, err := config.LoadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		panic(fmt.Sprintf("ошибка парсинга конфига, %v", err))
	}

	eventsStorage := bootstrap.InitPGStorage(cfg)
	eventService := bootstrap.InitEventService(eventsStorage, cfg)
	eventCreatedProducer := bootstrap.InitEventCreatedProducer(cfg)
	eventApi := bootstrap.InitEventServiceAPI(eventService, eventCreatedProducer)

	bootstrap.AppRun(*eventApi)
}
