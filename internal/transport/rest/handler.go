package rest

import (
	"github.com/gofiber/fiber"
	jsoniter "github.com/json-iterator/go"

	"RunLengthEncoding/internal/models"
	"RunLengthEncoding/internal/services"
)

func EncodeHandler(c *fiber.Ctx) {
	bodyByte := c.Fasthttp.Request.Body()
	msg := models.Msg{}
	jsoniter.Unmarshal(bodyByte, &msg)
	res := services.RunLengthEncode(msg.Data)
	c.Write(res)
}

func DecodeHandler(c *fiber.Ctx) {
	bodyByte := c.Fasthttp.Request.Body()
	msg := models.Msg{}
	jsoniter.Unmarshal(bodyByte, &msg)
	res := services.RunLengthDecode(msg.Data)
	c.Write(res)
}
