package main

import (
	"RunLengthEncoding/utils"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestRunLengthEncode(t *testing.T) {
	testTable := []struct {
		name     string
		msg      []string
		expected []string
	}{
		{
			name:     "test_1",
			msg:      []string{"AAAAA", "AAA BBB", "ABC DDD", "     ", "A B C", "ABC"},
			expected: []string{"5A", "3A 3B", "ABC 3D", "5 ", "A B C", "ABC"},
		},
		{
			name:     "test_2",
			msg:      []string{"AAA", "BC D", "EEEF"},
			expected: []string{"3A", "BC D", "3EF"},
		},
		{
			name:     "test_3",
			msg:      []string{""},
			expected: []string{""},
		},
		{
			name:     "test_4",
			msg:      []string{},
			expected: []string{},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := RunLengthEncode(testCase.msg)
			assert.Equal(t, result, testCase.expected)
		})
	}
}

func TestRunLengthDecode(t *testing.T) {
	testTable := []struct {
		name     string
		msg      []string
		expected []string
	}{
		{
			name:     "test_1",
			msg:      []string{"5A", "3A 3B", "ABC 3D", "5 ", "A B C", "ABC"},
			expected: []string{"AAAAA", "AAA BBB", "ABC DDD", "     ", "A B C", "ABC"},
		},
		{
			name:     "test_2",
			msg:      []string{"3A", "BC D", "3EF"},
			expected: []string{"AAA", "BC D", "EEEF"},
		},
		{
			name:     "test_3",
			msg:      []string{""},
			expected: []string{""},
		},
		{
			name:     "test_4",
			msg:      []string{},
			expected: []string{},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := RunLengthDecode(testCase.msg)
			assert.Equal(t, result, testCase.expected)
		})
	}
}

func TestEncodeDecodeMandelbrot(t *testing.T) {
	msg := utils.CreateMandelbrot()
	code := RunLengthEncode(msg)
	res := (RunLengthDecode(code))

	assert.Equal(t, res, msg)
}

func BenchmarkRunLengthEncode(b *testing.B) {
	msg := utils.CreateMandelbrot()
	b.StartTimer()
	RunLengthEncode(msg)
	b.StopTimer()
}
