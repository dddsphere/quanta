package entity

import "time"

type (
	Metadata interface {
		Version() int
		Timestamp() time.Time
		Props() map[string]any
	}

	Meta struct {
		version   int64
		timestamp time.Time
		props     map[string]any
	}
)
