package rest

import (
	"github.com/gofiber/fiber/v2"
)

type Rle interface {
	Encode(c *fiber.Ctx) error
	Decode(c *fiber.Ctx) error
}

type rle struct {
}

func NewRleHandler() Rle {
	return &rle{}
}
