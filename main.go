package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/miko089/mipm-server/endpoints"
	"os"
)

func main() {
	app := fiber.New()
	endpoints.RegisterHandlers(app)
	err := app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
	log.Fatal(err)
	panic("server is dead")
}
