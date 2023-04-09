package utils

import "strconv"

func NotNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err != nil
}

func IsMatch(arr1, arr2 []string) bool {
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func GetParts (lenChunk, lenMsg int) int {
	parts := lenMsg/lenChunk
	if lenMsg%lenChunk != 0 {
		parts++
	}
	return parts
}