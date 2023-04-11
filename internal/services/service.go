package services

import (
	"context"
)

type RleService interface {
	Encode(ctx context.Context, msg []string) ([]string, error)
	Decode(ctx context.Context, msg []string) ([]string, error)
}

type rleService struct {
}

func NewRleService() *rleService {
	return &rleService{}
}

func (s rleService) Encode(ctx context.Context, msg []string) ([]string, error) {
	return poolWorkers(ctx, msg, runLengthEncode)
}

func (s rleService) Decode(ctx context.Context, msg []string) ([]string, error) {
	return poolWorkers(ctx, msg, runLengthDecode)
}
