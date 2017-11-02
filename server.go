package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func db_exec(db *sql.DB, q string) {
	var _, err = db.Exec(q)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	e := echo.New()
	e.Static("/vue", "vue-app/dist")
	e.Static("/vue-org", "vue-app")
	e.GET("/", func(c echo.Context) error {
		os.Create("./data.db")

		var db *sql.DB

		db, err := sql.Open("sqlite3", "./sample.db")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		q := "INSERT INTO post "
		q += " (title, body)"
		q += " VALUES"
		q += " ('title', 'hogehoge')"
		db_exec(db, q)

		db.Close()

		return c.String(http.StatusOK, "Hello, World!2")
	})
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
