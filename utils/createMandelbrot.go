package utils

import (
	"math/cmplx"
)

const (
	mandelbrotScale = 2
)

type (
	Value struct {
		limit  float64
		symbol byte
	}
	ValueSymbols [5]Value
)

var (
	symbol = ValueSymbols{
		Value{0.45, 'A'},
		Value{0.3, 'B'},
		Value{0.2, 'C'},
		Value{0.0, 'D'},
	}
)

func (values ValueSymbols) From(value float64) byte {
	for _, v := range values {
		if value > v.limit {
			return v.symbol
		}
	}
	return ' '
}

func mandelbrot(a complex128) (z complex128) {
	for i := 0; i < 50; i++ {
		z = z*z + a
	}
	return
}

func CreateMandelbrot() []string {
	var lines []string
	var value float64
	for y := 1.0; y >= -1.0; y -= 0.05 / mandelbrotScale {
		var line []byte
		for x := -2.0; x <= 0.5; x += 0.0315 / mandelbrotScale {
			value = cmplx.Abs(mandelbrot(complex(x, y)))
			line = append(line, symbol.From(value))
		}
		lines = append(lines, string(line))
	}
	return lines
}
