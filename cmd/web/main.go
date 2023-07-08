package main

import (
	"fmt"
	"log"

	"github.com/lucianocorreia/go-starter/config"
)

func main() {
	cfg := config.New()

	fmt.Println(cfg)

	fmt.Println("Hello, World!!!!!")

	log.Println("Starting the application...")
}
