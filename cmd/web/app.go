package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lucianocorreia/go-starter/config"
	database "github.com/lucianocorreia/go-starter/database/sqlc"
	"github.com/lucianocorreia/go-starter/handlers"
)

const (
	version    = "0.0.1"
	cssVersion = "1"
)

// App is the application struct
type App struct {
	Config   config.Config
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Version  string
	Handlers *handlers.Handlers
}

func NewApp() *App {
	cfg := config.NewConfig()

	infoLog := log.New(log.Writer(), "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(log.Writer(), "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	connPool, err := pgxpool.New(context.Background(), cfg.GetDBSource())
	if err != nil {
		errorLog.Fatal("cannot connect to db:", err)
	}

	store := database.NewStore(connPool)
	handlers := handlers.NewHandlers(store)

	return &App{
		Config:   cfg,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Version:  version,
		Handlers: handlers,
	}
}

func (a *App) Run() error {
	engine := html.New("./views", ".gohtml")

	server := fiber.New(
		fiber.Config{
			Views:       engine,
			ViewsLayout: "layouts/main",
		},
	)

	a.setupRoutes(server)

	server.Static("/", "./public")

	return server.Listen(fmt.Sprintf(":%s", a.Config.App.Port))
}
