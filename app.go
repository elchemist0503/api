package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

type App struct {
	Port    int
	Prefix  string
	Handler http.Handler
}

func New(port int, prefix string) *App {
	return &App{
		Port:   port,
		Prefix: prefix,
	}
}

func (app *App) EnableCors(h http.Handler, opt cors.Options) {
	corsConfig := cors.New(opt)
	app.Handler = corsConfig.Handler(h)
}

func (app *App) Run() {
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", app.Port),
		Handler: app.Handler,
	}

	log.Printf("Starting server on port :%d", app.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
