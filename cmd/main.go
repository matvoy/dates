package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()
	app := newApp(db, e)
	app.init()
	e.Logger.Fatal(e.Start(":3333"))
}
