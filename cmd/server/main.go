package main

import (
	"fmt"
	"net/http"

	handler "github.com/charlallison/go-rest-api/internal/transport/http"
)

// App - struct that contains things like pointers to
// database connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println(`Setting up our APP`)

	h := handler.NewHandler()
	h.SetupRoutes()

	if err := http.ListenAndServe(":8080", h.Router); err != nil {
		fmt.Println(`Failed to setup server`)
	}

	return nil
}

func main() {
	fmt.Println("Go REST API course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
	}
}
