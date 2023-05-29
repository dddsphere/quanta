package pg

import (
	"context"
	"fmt"

	"github.com/dddsphere/quanta/internal/core/errors"
	"github.com/dddsphere/quanta/internal/system"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type (
	WriteStore struct {
		system.Worker
		db *sqlx.DB
	}
)

const (
	name = "write-store"
)

func NewWriteStore(opts ...system.Option) *WriteStore {
	return &WriteStore{
		Worker: system.NewWorker(name, opts...),
	}
}

func (ws *WriteStore) Setup(ctx context.Context) error {
	db, err := ws.getConn()
	if err != nil {
		msg := fmt.Sprintf("%s setup error", ws.Name())
		return errors.Wrap(msg, err)
	}

	ws.db = db

	return nil
}

// NOTE: This will be injected later from app
func (ws *WriteStore) getConn() (*sqlx.DB, error) {
	// FIX: Get conn. string values from config
	connString := "user=username password=password dbname=db host=hostname port=port sslmode=require"

	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		return nil, err // TODO: Wrap error
	}

	err = db.Ping()
	if err != nil {
		return nil, err // TODO: Wrap error
	}

	ws.Log().Infof("%s database connected!", ws.Name())
	return db, nil
}
