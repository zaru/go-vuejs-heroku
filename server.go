package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.Static("/vue", "vue-app/dist")
	e.Static("/vue-org", "vue-app")
	e.GET("/", func(c echo.Context) error {
		db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
		defer db.Close()
		_, _ = db.Exec("INSERT INTO post(title, body) VALUES($1, $2);", "foo", "bar")

		return c.String(http.StatusOK, "Hello, World!3")
	})
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
