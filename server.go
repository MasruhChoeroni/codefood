package main

import (
	"codefood/db"
	"codefood/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":3030"))
}
