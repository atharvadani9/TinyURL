package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"tinyurl/internal/api"
	"tinyurl/internal/migrations"
	"tinyurl/internal/store"
)

type Application struct {
	Logger         *log.Logger
	TinyURLHandler *api.TinyURLHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	tinyURLStore := store.NewPostgresTinyURLStore(pgDB)

	tinyURLHandler := api.NewTinyURLHandler(tinyURLStore, logger)

	app := &Application{
		Logger:         logger,
		TinyURLHandler: tinyURLHandler,
		DB:             pgDB,
	}

	return app, nil
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status: OK\n")
	app.Logger.Println("INFO: Health check passed")
}
