package services

import (
	"context"
)

type RleService interface {
	Encode(ctx context.Context, msg []string) []string
	Decode(ctx context.Context, msg []string) []string
}

type rleService struct {
}

func NewRleService() *rleService {
	return &rleService{}
}

func (s rleService) Encode(ctx context.Context, msg []string) []string { // TODO возвращать и обрабатывать ошибки
	return poolWorkers(ctx, msg, runLengthEncode)
}

func (s rleService) Decode(ctx context.Context, msg []string) []string { // TODO возвращать и обрабатывать ошибки
	return poolWorkers(ctx, msg, runLengthDecode)
}
