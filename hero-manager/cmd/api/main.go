package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/rstropek/hero-manager/internal/data"
)

type config struct {
	port int
	env  string // dev, staging, prod, ...
	dsn  string // data source name
}

type application struct {
	config config
	logger *slog.Logger
	repos  *data.AppRepository
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|staging|prod)")
	flag.StringVar(&cfg.dsn, "dsn", os.Getenv("POSTGRES_DSN"), "PostgreSQL DSN")
	flag.Parse()

	fmt.Printf("Starting server on port %d\n", cfg.port)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		repos:  data.NewAppRepository(db),
	}

	srv := &http.Server{
		Addr:     fmt.Sprintf(":%d", cfg.port),
		Handler:  app.routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
