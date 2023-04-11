package services

import (
	"strconv"
	"strings"
)

func runLengthEncode(msg []string) ([]string, error) {
	res := make([]string, len(msg))
	sb := strings.Builder{}

	for i := range msg {
		var (
			firsrElement rune
			count        int
		)

		for _, elem := range msg[i] {
			if elem == firsrElement {
				count++
				continue
			} else if elem != firsrElement && count > 0 {
				sb.WriteString(strconv.Itoa(count + 1))
				sb.WriteRune(firsrElement)
				count = 0
			} else if elem != firsrElement && firsrElement != 0 && count == 0 {
				sb.WriteRune(firsrElement)
			}
			firsrElement = elem
		}

		if count > 0 {
			sb.WriteString(strconv.Itoa(count + 1))
			sb.WriteRune(firsrElement)
		} else if firsrElement == 0 {
			sb.WriteString("")
		} else {
			sb.WriteRune(firsrElement)
		}

		res[i] = sb.String()
		sb.Reset()
	}

	return res, nil
}
