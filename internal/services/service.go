package services

import "github.com/gofiber/fiber/v2"

type RleService interface {
	Encode(c *fiber.Ctx, msg []string) []string
	Decode(c *fiber.Ctx, msg []string) []string
}

type rleService struct {
}

func NewRleService() *rleService {
	return &rleService{}
}

func (s rleService) Encode(c *fiber.Ctx, msg []string) []string { // TODO возвращать и обрабатывать ошибки
	return poolWorkers(c, msg, runLengthEncode)
}

func (s rleService) Decode(c *fiber.Ctx, msg []string) []string { // TODO возвращать и обрабатывать ошибки
	return poolWorkers(c, msg, runLengthDecode)
}
