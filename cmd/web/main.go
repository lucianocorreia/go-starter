package main

import (
	_ "github.com/jackc/pgx/v5"
)

func main() {
	app := NewApp()

	app.ErrorLog.Fatal(app.Run())
}
