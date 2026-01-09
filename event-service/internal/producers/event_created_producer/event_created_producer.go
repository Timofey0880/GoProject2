package event_created_producer

type EventCreatedProducer struct {
	kafkaBrokers []string
	topicName    string
}

func NewEventCreatedProducer(kafkaBrokers []string, topicName string) *EventCreatedProducer {
	return &EventCreatedProducer{
		kafkaBrokers: kafkaBrokers,
		topicName:    topicName,
	}
}
