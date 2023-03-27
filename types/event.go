package types

type Event struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	OwnerID      string   `json:"owner_id"`
	Participants []string `json:"participants"`
}

func NewEvent(name string, description string, owner_id string) *Event {
	return &Event{
		Name:        name,
		Description: description,
		OwnerID:     owner_id,
	}
}
