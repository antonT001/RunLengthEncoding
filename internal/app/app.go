package app

import (
	"RunLengthEncoding/internal/transport/rest"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run() {
	app := fiber.New()
	app.Use(recover.New())
	app.Post("/api/encode", rest.EncodeHandler)
	app.Post("/api/decode", rest.DecodeHandler)
	log.Fatal(app.Listen(":3000"))
}
