package services

import (
	"RunLengthEncoding/internal/utils"
	"strconv"
	"strings"
)

// TODO
// возвращать и обрабатывать ошибки
// вынести контстанты
// оптимизировать логику
// пооборачивать в функции
func runLengthDecode(msg []string) []string {
	var numberStr string
	var number int
	var secondElement string
	res := make([]string, len(msg))
	sb := strings.Builder{}
	for i := range msg {
		for _, elem := range msg[i] {
			if elem >= 48 && elem <= 57 && secondElement == "" {
				numberStr += string(elem)
			} else if secondElement == "" && numberStr != "" {
				secondElement = string(elem)
			} else if elem == 32 || utils.NotNumber(secondElement) && numberStr != "" {
				number, _ = strconv.Atoi(numberStr) // обработать ошибку
				sb.Grow(sb.Len() + number + len(secondElement))
				for j := 0; j < number; j++ {
					sb.WriteString(secondElement)
				}
				if utils.NotNumber(string(elem)) {
					sb.WriteRune(elem)
					numberStr = ""
				} else {
					numberStr = string(elem)
				}
				secondElement = ""
			} else {
				sb.WriteRune(elem)
			}
		}
		number, _ = strconv.Atoi(numberStr) // обработать ошибку
		sb.Grow(sb.Len() + number + len(secondElement))
		for j := 0; j < number; j++ {
			sb.WriteString(secondElement)
		}
		secondElement = ""
		numberStr = ""
		res[i] = sb.String()
		sb.Reset()
	}

	return res
}
