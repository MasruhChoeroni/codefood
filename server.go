package main

import (
	"codefood/config"
	"codefood/db"
	"codefood/routes"
	"fmt"
)

func main() {
	db.Init()
	e := routes.Init()
	fmt.Println("HOST : " + config.GetConfig().DB_HOST)
	fmt.Println("NAME : " + config.GetConfig().DB_NAME)
	fmt.Println("USER : " + config.GetConfig().DB_USERNAME)
	fmt.Println("PASS : " + config.GetConfig().DB_PASSWORD)
	fmt.Println("PORT : " + config.GetConfig().DB_PORT)
	e.Logger.Fatal(e.Start(":3030"))
}
