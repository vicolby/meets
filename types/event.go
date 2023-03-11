package types

type Event struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     string `json:"owner_id"`
}

func NewEvent(name string, description string, ownerID string) *Event {
	return &Event{
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
	}
}
