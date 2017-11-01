package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/vue", "vue-app/dist")
	e.Static("/vue-org", "vue-app")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!2")
	})
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
