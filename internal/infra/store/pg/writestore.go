package pg

import (
	"context"
	"fmt"

	"github.com/dddsphere/quanta/internal/core/errors"
	"github.com/dddsphere/quanta/internal/event"
	db "github.com/dddsphere/quanta/internal/infra"
	"github.com/dddsphere/quanta/internal/sys"

	_ "github.com/lib/pq"
)

type (
	WriteStore struct {
		sys.Worker
		db *db.DB
	}
)

const (
	name = "write-store"
)

func NewWriteStore(db *db.DB, opts ...sys.Option) *WriteStore {
	return &WriteStore{
		Worker: sys.NewWorker(name, opts...),
		db:     db,
	}
}

func (ws *WriteStore) Setup(ctx context.Context) error {
	err := ws.db.Setup(ctx)
	if err != nil {
		msg := fmt.Sprintf("%s setup error", err)
		return errors.Wrap(msg, err)
	}

	return nil
}

func (ws *WriteStore) WriteEvent(ctx context.Context, ev event.Event) (err error) {
	events := []event.Event{ev}

	_, err = ws.db.NamedExec(`INSERT INTO events (stream_name, stream_id, stream_version, event_name, event_id, payload, timestamp)
        VALUES (:stream_name, :stream_id, :stream_version, :event_name, :event_id, :payload, :timestamp)`, events)

	if err != nil {
		return errors.Wrap("write event error", err)
	}

	return nil
}
