package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/dddsphere/quanta/internal/core/errors"
	"github.com/dddsphere/quanta/internal/system"
)

type (
	DB struct {
		system.Worker
		*sqlx.DB
	}
)

const (
	name = "pg-db"
)

func NewDB(opts ...system.Option) *DB {
	return &DB{
		Worker: system.NewWorker(name, opts...),
		DB:     nil,
	}
}

func (db *DB) Setup(ctx context.Context) error {
	return db.Connect()
}

func (db *DB) Connect() error {
	pgdb, err := sqlx.Open("postgres", db.connString())
	if err != nil {
		msg := fmt.Sprintf("%s connection error", db.Name())
		return errors.Wrap(msg, err)
	}

	err = pgdb.Ping()
	if err != nil {
		msg := fmt.Sprintf("%s ping connection error", db.Name())
		return errors.Wrap(msg, err)
	}

	db.Log().Infof("%s database connected!", db.Name())
	return nil
}

func (db *DB) connString() (connString string) {
	cfg := db.Cfg()
	user := cfg.GetString("store.write.db.user")
	pass := cfg.GetString("store.write.db.pass")
	name := cfg.GetString("store.write.db.db")
	host := cfg.GetString("store.write.db.host")
	port := cfg.GetInt("store.write.db.port")
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=require", user, pass, name, host, port)
}
