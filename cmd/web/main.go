package main

import (
	"fmt"

	"github.com/lucianocorreia/go-starter/pkg/services"
)

func main() {
	c := services.NewContainer()
	defer func() {
		if err := c.Shutdown(); err != nil {
			c.Web.Logger.Fatal(err)
		}
	}()

	fmt.Println(c)

	//  Starting the application
	fmt.Println("Strating the application...")
}
