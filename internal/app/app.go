package app

import (
	"RunLengthEncoding/internal/services"
	"RunLengthEncoding/internal/transport/rest"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run() {
	rleService := services.NewRleService()
	rleHandler := rest.NewRleHandler(rleService)

	app := fiber.New()

	app.Use(recover.New())

	apiRoute := app.Route("/api", func(router fiber.Router) {})

	rleRoute := apiRoute.Route("/rle", func(router fiber.Router) {})
	rleRoute.Post("/encode", rleHandler.Encode)
	rleRoute.Post("/decode", rleHandler.Decode)

	log.Fatal(app.Listen(":3000"))
}
