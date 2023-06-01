package handler

import (
	"context"

	"github.com/dddsphere/quanta/internal/dto"
	"github.com/dddsphere/quanta/internal/event"
	"github.com/dddsphere/quanta/internal/infra/store"
	"github.com/dddsphere/quanta/internal/sys"
)

type Handler struct {
	sys.Worker
	store store.Write // Inject the write store dependency
}

const (
	name = "app-event-handler"
)

func NewHandler(store store.Write, opts ...sys.Option) *Handler {
	return &Handler{
		Worker: sys.NewWorker(name, opts...),
		store:  store,
	}
}

func (h *Handler) CreateListEvent(ctx context.Context, data dto.CreateList) error {
	// TODO: Validate evt data

	evt := event.Event{
		StreamName:    event.ListStream,
		StreamID:      event.GenID(),
		StreamVersion: "0.0.1", // NOTE: Will see later where to get this value
		Name:          event.CreateListCmd,
		ID:            data.ReqID, // NOTE: Using request ID for now
		Payload:       data.Payload,
	}

	err := h.store.WriteEvent(ctx, evt)
	if err != nil {
		h.Log().Errorf("cannot write event into the store", err)
		return err
	}

	return nil
}
