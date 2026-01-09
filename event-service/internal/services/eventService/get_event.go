package eventService

import (
	"context"

	"github.com/Timofey0880/GoProject2/event-service/internal/models"
)

func (s *EventService) GetEvent(ctx context.Context, id uint64) (*models.Event, error) {
	return s.eventStorage.GetEvent(ctx, id)
}
