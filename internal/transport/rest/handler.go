package rest

import (
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"

	"RunLengthEncoding/internal/models"
	"RunLengthEncoding/internal/services"
)

func EncodeHandler(c *fiber.Ctx) error {
	bodyByte := c.Body()
	msg := models.Msg{}
	jsoniter.Unmarshal(bodyByte, &msg)
	res := services.RunLengthEncode(msg.Data)
	resByte, err := jsoniter.Marshal(res)
	if err != nil {
		return err // TODO обернуть
	}
	c.Write(resByte)
	return nil
}

func DecodeHandler(c *fiber.Ctx) error {
	bodyByte := c.Body()
	msg := models.Msg{}
	jsoniter.Unmarshal(bodyByte, &msg)
	res := services.RunLengthDecode(msg.Data)
	resByte, err := jsoniter.Marshal(res)
	if err != nil {
		return err // TODO обернуть
	}
	c.Write(resByte)
	return nil
}
