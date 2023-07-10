package main

import (
	_ "github.com/lib/pq"
)

func main() {
	app := NewApp()

	app.ErrorLog.Fatal(app.Run())
}
