package utils

import "strconv"

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
