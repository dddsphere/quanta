package pg

import (
	"github.com/dddsphere/quanta/internal/system"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type (
	WriteStore struct {
		system.Worker
		db sqlx.DB
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

// NOTE: This will be injected later from app
func (ws *WriteStore) getConn() (*sqlx.DB, error) {
	// FIX: Get conn. string values from config
	connString := "user=username password=password dbname=db host=hostname port=port sslmode=require"

	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		return nil, err // TODO: Wrap error
	}

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	ws.Log().Info("Connected!")
	return db, nil
}
