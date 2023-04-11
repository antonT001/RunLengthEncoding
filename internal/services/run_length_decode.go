package services

import (
	"RunLengthEncoding/internal/utils"
	"strconv"
	"strings"
)

func runLengthDecode(msg []string) ([]string, error) {
	res := make([]string, len(msg))
	sb := strings.Builder{}
	for i := range msg {
		var (
			numberStr     string
			secondElement string
		)

		for _, elem := range msg[i] {
			if utils.IsRuneNumber(elem) && secondElement == "" {
				numberStr += string(elem)
			} else if secondElement == "" && numberStr != "" {
				secondElement = string(elem)
			} else if utils.NotNumber(secondElement) && numberStr != "" {
				number, err := strconv.Atoi(numberStr)
				if err != nil {
					return nil, err
				}

				for j := 0; j < number; j++ {
					_, err := sb.WriteString(secondElement)
					if err != nil {
						return nil, err
					}
				}
				if !utils.IsRuneNumber(elem) {
					_, err := sb.WriteRune(elem)
					if err != nil {
						return nil, err
					}
					numberStr = ""
				} else {
					numberStr = string(elem)
				}
				secondElement = ""
			} else {
				_, err := sb.WriteRune(elem)
				if err != nil {
					return nil, err
				}
			}
		}

		if numberStr != "" {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				return nil, err
			}
			for j := 0; j < number; j++ {
				_, err = sb.WriteString(secondElement)
				if err != nil {
					return nil, err
				}
			}
		}

		res[i] = sb.String()
		sb.Reset()
	}

	return res, nil
}
