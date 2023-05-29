package store

import (
	"context"

	"github.com/dddsphere/quanta/internal/event"
)

type (
	Write interface {
		WriteEvent(ctx context.Context, ev event.Event) error
	}
)
