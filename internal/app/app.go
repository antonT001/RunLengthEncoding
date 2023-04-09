package app

import (
	"RunLengthEncoding/internal/transport/rest"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run() {
	rleHandler := rest.NewRleHandler()

	app := fiber.New()

	app.Use(recover.New())

	apiRoute := app.Route("/api", func(router fiber.Router) {})

	rleRoute := apiRoute.Route("/rle", func(router fiber.Router) {})
	rleRoute.Post("/encode", rleHandler.Encode)
	rleRoute.Post("/decode", rleHandler.Decode)

	log.Fatal(app.Listen(":3000"))
}
