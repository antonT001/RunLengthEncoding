package app

import (
	"RunLengthEncoding/internal/transport/rest"
	"log"

	"github.com/gofiber/fiber"
)

func Run() {
	app := fiber.New()
	app.Post("/encode", rest.EncodeHandler)
	app.Post("/decode", rest.DecodeHandler)
	log.Fatal(app.Listen(":3000"))
}
