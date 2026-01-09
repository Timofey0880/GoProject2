package eventService

import (
	"errors"
	"fmt"

	"github.com/Timofey0880/GoProject2/event-service/internal/models"
)

func (s *EventService) validateEvent(event *models.Event) error {
	if len(event.Name) < s.minNameLen || len(event.Name) > s.maxNameLen {
		return errors.New("name invalid length")
	}
	if len(event.City) < s.minCityLen || len(event.City) > s.maxCityLen {
		return errors.New("city invalid length")
	}
	if len(event.User) == 0 {
		return fmt.Errorf("user empty")
	}
	return nil
}
