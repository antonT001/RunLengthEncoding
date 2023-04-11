package services

import (
	"RunLengthEncoding/internal/utils"
	"context"
	"fmt"
	"sync"
)

const (
	WORKERS   = 8
	LEN_CHUNK = 3
)

type partData struct {
	part int
	msg  []string
}

func poolWorkers(ctx context.Context, msg []string, fn func([]string) ([]string, error)) ([]string, error) {
	var ofset int

	parts := utils.GetParts(LEN_CHUNK, len(msg))
	wg := sync.WaitGroup{}
	jobs := make(chan partData, WORKERS)
	results := make(chan partData, WORKERS)
	errorCh := make(chan error)

	for w := 0; w < WORKERS; w++ {
		wg.Add(1)
		go worker(ctx, &wg, w, jobs, results, errorCh, fn)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	go func() {
		defer close(jobs)
		for i := 0; i < parts; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				limit := ofset + LEN_CHUNK
				if limit > len(msg) {
					limit = len(msg)
				}

				jobs <- partData{
					part: i,
					msg:  msg[ofset:limit],
				}

				ofset += LEN_CHUNK
				limit += LEN_CHUNK
			}
		}
	}()

	storage := make(map[int][]string, parts)
	for res := range results {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("timeout")
		case err := <-errorCh:
			return nil, err
		default:
			storage[res.part] = res.msg
		}
	}

	res := make([]string, 0, len(msg))
	for i := 0; i < parts; i++ {
		res = append(res, storage[i]...)
	}
	return res, nil
}

func worker(
	ctx context.Context,
	wg *sync.WaitGroup,
	id int,
	jobs <-chan partData,
	results chan<- partData,
	errorCh chan<- error,
	fn func([]string) ([]string, error)) {
	defer wg.Done()
	for {
		select {
		case j, ok := <-jobs:
			if !ok {
				return
			}
			res, err := fn(j.msg)
			if err != nil {
				errorCh <- err
				return
			}
			results <- partData{
				part: j.part,
				msg:  res,
			}
		case <-ctx.Done():
			return
		}
	}
}
