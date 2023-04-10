package rest

import (
	"RunLengthEncoding/internal/models"
	"RunLengthEncoding/internal/utils"
	"context"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

func (h rleHandler) Encode(c *fiber.Ctx) error {
	ctx, cancel := context.WithCancel(c.Context())
	defer cancel()
	msg := models.Msg{}
	bodyByte := c.Body()
	err := jsoniter.Unmarshal(bodyByte, &msg)
	if err != nil {
		return utils.HttpResponse(c, models.OutError{
			Success: false,
			Error:   utils.StringPointer("[Encode]: can not unmarshal body: " + err.Error()),
		}, http.StatusInternalServerError)
	}

	err = validationEncode(msg)
	if err != nil {
		return utils.HttpResponse(c, models.OutError{
			Success: false,
			Error:   utils.StringPointer("[Encode]: validation error: " + err.Error()),
		}, http.StatusBadRequest)
	}

	res := h.rleService.Encode(ctx, msg.Data) // TODO возвращать и обрабатывать ошибки
	return utils.HttpResponse(c, models.RleResponse{
		Success: true,
		Data:    res,
	}, http.StatusOK)
}

func validationEncode(msg models.Msg) error {
	if len(msg.Data) == 0 {
		return fmt.Errorf("len(msg.Data) == 0")
	}
	return nil
}
