package eventService

import (
	"context"
	"testing"

	"github.com/Timofey0880/GoProject2/event-service/internal/models"
	"github.com/Timofey0880/GoProject2/event-service/internal/services/eventService/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"gotest.tools/v3/assert"
)

type EventServiceSuite struct {
	suite.Suite
	ctx           context.Context
	eventStorage  *mocks.EventStorage
	eventProducer *mocks.EventProducer
	eventService  *EventService
}

func (s *EventServiceSuite) SetupTest() {
	s.eventStorage = mocks.NewEventStorage(s.T())
	s.eventProducer = mocks.NewEventProducer(s.T())
	s.ctx = context.Background()
	s.eventService = NewEventService(s.ctx, s.eventStorage, s.eventProducer, 1, 100, 1, 100)
}

func (s *EventServiceSuite) TestCreateSuccess() {
	event := &models.Event{Name: "Event1", City: "City1", User: "User1"}
	s.eventStorage.EXPECT().CreateEvent(s.ctx, event).Return(uint64(1), nil)
	s.eventProducer.EXPECT().ProduceEventCreated(s.ctx, event).Return(nil)

	id, err := s.eventService.CreateEvent(s.ctx, event)

	assert.NilError(s.T(), err)
	assert.Equal(s.T(), uint64(1), id)
}

func (s *EventServiceSuite) TestCreateInvalidName() {
	event := &models.Event{Name: "", City: "City1", User: "User1"}

	_, err := s.eventService.CreateEvent(s.ctx, event)

	assert.Check(s.T(), err != nil)
}

func (s *EventServiceSuite) TestCreateStorageError() {
	event := &models.Event{Name: "Event1", City: "City1", User: "User1"}
	wantErr := errors.New("error")
	s.eventStorage.EXPECT().CreateEvent(s.ctx, event).Return(uint64(0), wantErr)

	_, err := s.eventService.CreateEvent(s.ctx, event)

	assert.ErrorIs(s.T(), err, wantErr)
}

func TestEventServiceSuite(t *testing.T) {
	suite.Run(t, new(EventServiceSuite))
}
