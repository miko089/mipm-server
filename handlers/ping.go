package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func PingHandler(c *fiber.Ctx) error {
	log.Infof("ping from {%s}", c.IP())
	return c.SendString("pong")
}
