package services

import (
	"RunLengthEncoding/internal/models"
	"RunLengthEncoding/internal/utils"
	"strconv"
	"strings"
)

func RunLengthEncode(msg []string) []string {
	var ofset int
	parts := utils.GetParts(LEN_CHUNK, len(msg))
	storage := models.Storage{M: make(map[int][]string, parts)}
	for i := 0; i < parts; i++ {
		limit := ofset + LEN_CHUNK
		if limit > len(msg) {
			limit = len(msg)
		}
		storage.Wg.Add(1)
		go encode(msg[ofset:limit], i, &storage)
		ofset += LEN_CHUNK
		limit += LEN_CHUNK
	}
	storage.Wg.Wait()
	res := make([]string, 0, len(msg))
	for i := 0; i < parts; i++ {
		res = append(res, storage.M[i]...)
	}
	return res
}

func encode(msg []string, part int, storage *models.Storage) {
	var firsrElement rune
	count := 1
	res := make([]string, len(msg))
	sb := strings.Builder{}
	for i := range msg {
		for _, elem := range msg[i] {
			if firsrElement == 0 {
				firsrElement = elem
			} else if elem == firsrElement {
				count++
				continue
			} else if elem != firsrElement && count > 1 {
				sb.WriteString(strconv.Itoa(count))
				sb.WriteRune(firsrElement)
				count = 1
			} else if elem != firsrElement && count == 1 {
				sb.WriteRune(firsrElement)
			}
			firsrElement = elem
		}
		if count > 1 {
			sb.WriteString(strconv.Itoa(count))
			sb.WriteRune(firsrElement)
			count = 1
		} else if firsrElement == 0 {
			sb.WriteString("")
		} else {
			sb.WriteRune(firsrElement)
		}
		firsrElement = 0
		res[i] = sb.String()
		sb.Reset()
	}
	storage.Mu.Lock()
	defer storage.Mu.Unlock()
	defer storage.Wg.Done()
	storage.M[part] = res
}
