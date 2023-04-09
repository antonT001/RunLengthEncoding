package rest

import (
	"RunLengthEncoding/internal/models"
	"RunLengthEncoding/internal/services"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

func (r rle) Decode(c *fiber.Ctx) error {
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
