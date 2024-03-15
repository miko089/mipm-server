package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miko089/mipm-server/handlers"
)

func RegisterHandlers(app *fiber.App) {
	app.Get("/ping", handlers.PingHandler)
	pkgs := app.Group("/pkgs")
	pkgs.Post("/upload")

}
