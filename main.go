package main

import (
	"test/config"
	"test/middlewares"
	"test/routes"
)

func main() {
	config.InitDb()
	e := routes.New()
	middlewares.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}
