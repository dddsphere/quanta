package pg

import (
	"context"
	"fmt"

	"github.com/dddsphere/quanta/internal/core/errors"
	db "github.com/dddsphere/quanta/internal/infra"
	"github.com/dddsphere/quanta/internal/system"

	_ "github.com/lib/pq"
)

type (
	WriteStore struct {
		system.Worker
		db *db.DB
	}
)

const (
	name = "write-store"
)

func NewWriteStore(db *db.DB, opts ...system.Option) *WriteStore {
	return &WriteStore{
		Worker: system.NewWorker(name, opts...),
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
