package main

import (
	"flag"
	"log"
	"net/http"
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

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(cfg.staticDir))

	// Handlers
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Logger
	log.Print("starting server on: ", cfg.addr, " & static dir ", cfg.staticDir)

	// Error handler
	err := http.ListenAndServe(cfg.addr, mux)
	log.Fatal(err)
}