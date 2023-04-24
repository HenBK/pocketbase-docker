package main

import (
	"log"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/custom", func(c echo.Context) error {
			log.Println("pinged from custom route")
			return c.String(200, "Hello world!\n")
		})
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
