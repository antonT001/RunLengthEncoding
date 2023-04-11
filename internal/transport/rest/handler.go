package rest

import (
	"RunLengthEncoding/internal/services"

	"github.com/gofiber/fiber/v2"
)

const TIMEOUT = 2

type RleHandler interface {
	Encode(c *fiber.Ctx) error
	Decode(c *fiber.Ctx) error
}

type rleHandler struct {
	rleService services.RleService
}

func NewRleHandler(rleService services.RleService) *rleHandler {
	return &rleHandler{
		rleService: rleService,
	}
}
