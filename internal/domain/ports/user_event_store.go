package ports

type UserEvent struct {
	Event
	UserID          string
}

type UserEventStore interface {
	GetUserEvents() ([]UserEvent, error)
	StoreEvent(userEvent UserEvent) error
}
