package rest

import (
	"RunLengthEncoding/internal/models"
	"RunLengthEncoding/internal/utils"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

func (h rleHandler) Decode(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), TIMEOUT*time.Second)
	defer cancel()

	msg := models.Msg{}
	bodyByte := c.Body()
	err := jsoniter.Unmarshal(bodyByte, &msg)
	if err != nil {
		return utils.HttpResponse(c, models.OutError{
			Success: false,
			Error:   utils.StringPointer("[Decode]: can not unmarshal body: " + err.Error()),
		}, http.StatusInternalServerError)
	}

	err = validationDecode(msg)
	if err != nil {
		return utils.HttpResponse(c, models.OutError{
			Success: false,
			Error:   utils.StringPointer("[Decode]: validation error: " + err.Error()),
		}, http.StatusBadRequest)
	}

	res, err := h.rleService.Decode(ctx, msg.Data)
	if err != nil {
		return utils.HttpResponse(c, models.OutError{
			Success: false,
			Error:   utils.StringPointer("[Encode]: service error: " + err.Error()),
		}, http.StatusInternalServerError)
	}

	return utils.HttpResponse(c, models.RleResponse{
		Success: true,
		Data:    res,
	}, http.StatusOK)
}

func validationDecode(msg models.Msg) error {
	if len(msg.Data) == 0 {
		return fmt.Errorf("len(msg.Data) == 0")
	}
	return nil
}
