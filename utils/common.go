package utils

import "strconv"

func NotNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err != nil
}

func IsCompare(arr1, arr2 []string) bool {
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}