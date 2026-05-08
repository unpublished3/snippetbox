package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

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

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(cfg.staticDir))

	// Handlers
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Logger
	logger.Info("starting server on: ", slog.String("addr ", cfg.addr), slog.String("static", cfg.staticDir))

	// Error handler
	err := http.ListenAndServe(cfg.addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}