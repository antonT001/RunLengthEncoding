package services

import (
	"RunLengthEncoding/internal/utils"
	"sync"

	"github.com/gofiber/fiber/v2"
)

const (
	WORKERS   = 8
	LEN_CHUNK = 3
)

type partData struct {
	part int
	msg  []string
}

// TODO возвращать и обрабатывать ошибки
func poolWorkers(c *fiber.Ctx, msg []string, fn func([]string) []string) []string {
	var ofset int

	parts := utils.GetParts(LEN_CHUNK, len(msg))
	wg := sync.WaitGroup{}
	jobs := make(chan partData, WORKERS)
	results := make(chan partData, WORKERS)

	for w := 0; w < WORKERS; w++ {
		go worker(c, &wg, w, jobs, results, fn)
	}

	for i := 0; i < parts; i++ {
		limit := ofset + LEN_CHUNK
		if limit > len(msg) {
			limit = len(msg)
		}
		wg.Add(1)
		jobs <- partData{
			part: i,
			msg:  msg[ofset:limit],
		}

		ofset += LEN_CHUNK
		limit += LEN_CHUNK
	}

	close(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()
	storage := make(map[int][]string, parts)
	for res := range results {
		storage[res.part] = res.msg
	}

	res := make([]string, 0, len(msg))
	for i := 0; i < parts; i++ {
		res = append(res, storage[i]...)
	}
	return res
}

func worker(c *fiber.Ctx, wg *sync.WaitGroup, id int, jobs <-chan partData, results chan<- partData, fn func([]string) []string) {
	for {
		select {
		case j, ok := <-jobs:
			if !ok {
				return
			}
			res := fn(j.msg)
			results <- partData{
				part: j.part,
				msg:  res,
			}
			wg.Done()
		case <-c.Context().Done():
			return
		}
	}
}
