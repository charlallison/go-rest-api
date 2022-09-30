package main

import (
	"net/http"

	"github.com/charlallison/go-rest-api/internal/comment"
	"github.com/charlallison/go-rest-api/internal/database"
	handler "github.com/charlallison/go-rest-api/internal/transport/http"
	log "github.com/sirupsen/logrus"
)

// App -contains Application information
type App struct {
	Name    string
	Version string
}

// Run - sets up our application
func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{
		"AppName":    app.Name,
		"AppVersion": app.Version,
	}).Info("Setting up our App")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDb(db)
	if err != nil {
		return err
	}

	service := comment.NewService(db)

	h := handler.NewHandler(service)
	h.SetupRoutes()

	if err := http.ListenAndServe(":8080", h.Router); err != nil {
		log.Error(`Failed to setup server`)
	}

	return nil
}

func main() {
	app := App{
		Name:    "Commenting service",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		log.Error("Error starting up our REST API")
		log.Fatal(err)
	}
}
