package event

import (
	"time"

	"github.com/google/uuid"
)

type (
	Event struct {
		StreamName    StreamName
		StreamID      string
		StreamVersion string
		Name          EventName
		ID            string
		Payload       []byte
		Timestamp     time.Time
	}
)

func GenID() (id string) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return uid.String()
	}

	return id
}

// WIP: This probably needs to be moved to a more appropriate place but for now
// it will be close to where it is used.
// Consider using a registry.

type (
	EventName string
)

const (
	CreateListCmd EventName = "cmd::create-list"
)

type (
	StreamName string
)

const (
	ListStream StreamName = "list"
)
