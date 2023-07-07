package main

import (
	"fmt"

	"github.com/lucianocorreia/go-starter/config"
)

func main() {
	cfg := config.New()

	fmt.Println(cfg)

	fmt.Println("Hello, World!!!!!")
}
