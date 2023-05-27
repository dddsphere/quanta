package main

import (
	"os"

	a "github.com/dddsphere/quanta/internal/app"
	l "github.com/dddsphere/quanta/internal/log"
)

const (
	name = "quanta"
	env  = "qta"
)

var (
	log l.Logger = l.NewLogger(l.Level.Info, false)
)

func main() {
	app := a.NewApp(name, env, log)

	err := app.Run()
	if err != nil {
		log.Errorf("%s exit error: %s", name, err.Error())
		os.Exit(1)
	}
}
