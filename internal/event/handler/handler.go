package handler

import (
	"context"

	"github.com/dddsphere/quanta/internal/dto"
	"github.com/dddsphere/quanta/internal/event"
	"github.com/dddsphere/quanta/internal/infra/store"
	"github.com/dddsphere/quanta/internal/system"
)

type Handler struct {
	system.Worker
	store store.Write // Inject the write store dependency
}

const (
	name = "app-event-handler"
)

func NewHandler(store store.Write, opts ...system.Option) *Handler {
	return &Handler{
		Worker: system.NewWorker(name, opts...),
		store:  store,
	}
}

func (h *Handler) CreateListEvent(ctx context.Context, data dto.CreateList) error {
	// TODO: Validate evt data

	evt := event.Event{
		ReqID: data.ReqID,
		Name:  data.Name,
	}

	err := h.store.WriteEvent(ctx, evt)
	if err != nil {
		h.Log().Errorf("cannot write event into the store", err)
		return err
	}

	return nil
}
