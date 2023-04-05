package main

import (
	"RunLengthEncoding/utils"
	"testing"
)

func TestRunLengthEncode(t *testing.T) {

	msg := []string{"AAAAA", "AAA BBB", "ABC DDD", "     ", "A B C", "ABC"}
	ecpected := []string{"5A", "3A 3B", "ABC 3D", "5 ", "A B C", "ABC"}

	result := RunLengthEncode(msg)

	if len(result) != len(ecpected) {
		t.Error("Incorrect result")
	}

	if !utils.IsCompare(result, ecpected) {
		t.Error("Incorrect result")
	}
}

func TestRunLengthDecode(t *testing.T) {
	msg := []string{"5A", "3A 3B", "ABC 3D", "5 ", "A B C", "ABC"}
	ecpected := []string{"AAAAA", "AAA BBB", "ABC DDD", "     ", "A B C", "ABC"}

	result := RunLengthDecode(msg)

	if len(result) != len(ecpected) {
		t.Error("Incorrect result")
	}

	if !utils.IsCompare(result, ecpected) {
		t.Error("Incorrect result")
	}
}
