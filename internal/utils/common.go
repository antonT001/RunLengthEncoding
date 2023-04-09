package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

const (
	SYSTEM_ERROR = "system error: "
)

func NotNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err != nil
}

func GetParts(lenChunk, lenMsg int) int {
	parts := lenMsg / lenChunk
	if lenMsg%lenChunk != 0 {
		parts++
	}
	return parts
}

func HttpResponse(c *fiber.Ctx, out interface{}, statusCode int) error {
	result, err := jsoniter.Marshal(out)
	if err != nil {
		err := c.SendString(SYSTEM_ERROR + "can not marshal response: " + err.Error())
		return err
	}

	err = c.SendStatus(statusCode)
	if err != nil {
		err := c.SendString(SYSTEM_ERROR + "can not send status code response: " + err.Error())
		return err
	}

	err = c.Send(result)
	if err != nil {
		err := c.SendString(SYSTEM_ERROR + "can not send response body: " + err.Error())
		return err
	}
	return nil
}

func StringPointer(value string) *string {
	return &value
}
