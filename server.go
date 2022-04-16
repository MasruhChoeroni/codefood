package main

import (
	"apps/db"
	"apps/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":3030"))
}
