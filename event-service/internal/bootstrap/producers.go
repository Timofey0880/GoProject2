package bootstrap

import (
	"fmt"

	"github.com/Timofey0880/GoProject2/event-service/config"
	"github.com/Timofey0880/GoProject2/event-service/internal/producers/event_created_producer"
)

func InitEventCreatedProducer(cfg *config.Config) *event_created_producer.EventCreatedProducer {
	kafkaBrokers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return event_created_producer.NewEventCreatedProducer(kafkaBrokers, cfg.Kafka.EventCreatedTopicName)
}
