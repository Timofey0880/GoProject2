package event_created_producer

import (
	"context"
	"encoding/json"

	"github.com/Timofey0880/GoProject2/event-service/internal/models"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
)

func (p *EventCreatedProducer) ProduceEventCreated(ctx context.Context, event *models.Event) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: p.kafkaBrokers,
		Topic:   p.topicName,
	})
	defer writer.Close()

	msg, err := json.Marshal(event)
	if err != nil {
		return errors.Wrap(err, "marshal event")
	}
	err = writer.WriteMessages(ctx, kafka.Message{Value: msg})
	if err != nil {
		return errors.Wrap(err, "write message")
	}
	return nil
}
