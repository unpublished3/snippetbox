package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

type config struct {
	addr string
	staticDir string
}

func main()  {
	var cfg config

	// Command-line arguments
	flag.StringVar(&cfg.addr ,"addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "staticDir", "./ui/static/", "Static directory path")
	flag.Parse();

	// Logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		AddSource: true,
	}))

	// Initialize application
	app := &application{logger: logger}

	// Logger
	logger.Info("starting server on: ", slog.String("addr ", cfg.addr), slog.String("static", cfg.staticDir))

	// Error handler
	err := http.ListenAndServe(cfg.addr, app.routes(cfg.staticDir) )
	logger.Error(err.Error())
	os.Exit(1)
}