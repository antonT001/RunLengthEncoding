package services

import (
	"strconv"
	"strings"
)

// TODO
// возвращать и обрабатывать ошибки
// вынести контстанты
// оптимизировать логику
// пооборачивать в функции
func runLengthEncode(msg []string) []string {
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

	return res
}
