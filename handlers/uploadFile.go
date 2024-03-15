package endpoints

import "github.com/gofiber/fiber/v2"

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status()
	}
}
