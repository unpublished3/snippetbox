package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"snippetbox.unpublished3/internal/models"
)

type application struct {
	logger *slog.Logger
	snippets *models.SnippetModel
	templateCache map[string]*template.Template
}

type config struct {
	addr string
	staticDir string
	dsn string
}

func main()  {
	var cfg config

	// Command-line arguments
	flag.StringVar(&cfg.addr ,"addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "staticDir", "./ui/static/", "Static directory path")
	flag.StringVar(&cfg.dsn, "dsn", "web:pass@/snippetbox?parseTime=true", "Mariadb data source name")

	flag.Parse();

	// Logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		AddSource: true,
	}))

	// Database
	db, err := openDB(cfg.dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	// Template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Initialize application
	app := &application{logger: logger, snippets: &models.SnippetModel{DB: db}, templateCache: templateCache}

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