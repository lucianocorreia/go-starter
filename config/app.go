package config

import (
	"database/sql"
	"log"

	database "github.com/lucianocorreia/go-starter/database/sqlc"
)

const (
	version    = "0.0.1"
	cssVersion = "1"
)

// App is the application struct
type App struct {
	config   Config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
	store    database.Store
}

func NewApp() *App {
	cfg := NewConfig()

	infoLog := log.New(log.Writer(), "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(log.Writer(), "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := sql.Open(cfg.GetDBDriver(), cfg.GetDBSource())
	if err != nil {
		errorLog.Fatal("cannot connect to db:", err)
	}

	store := database.NewStore(conn)

	return &App{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
		store:    store,
	}
}
