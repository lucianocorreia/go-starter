package main

import (
	"fmt"
	"log"

	"github.com/lucianocorreia/go-starter/config"

	_ "github.com/lib/pq"
)

func main() {
	app := config.NewApp()

	fmt.Println(app)

	log.Println("Starting the application...")
}
