package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_"github.com/go-sql-driver/mysql"
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

		// Database
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "Mariadb data source name")

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	// Initialize application
	app := &application{logger: logger}

	// Logger
	logger.Info("starting server on: ", slog.String("addr ", cfg.addr), slog.String("static", cfg.staticDir))

	// Error handler
	err = http.ListenAndServe(cfg.addr, app.routes(cfg.staticDir) )
	logger.Error(err.Error())
	os.Exit(1)
}

func  openDB(dsn string) (*sql.DB, error)  {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return  nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return  nil, err
	}
	return  db, nil
}