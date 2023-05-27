package entity

import "time"

type (
	Audit struct {
		createdAt time.Time
		updatedAt time.Time
	}
)
