package main

import (
	"testing"
)

var (
	inputSlice []string = []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
)

func TestFindCloseness(t *testing.T) {
	output := FindCloseness(inputSlice[0], inputSlice[5])

	if output != "ace" {
		t.Fatalf("Output expected ace actual %s", output)
	}

	if len(output) != 3 {
		t.Fatalf("Output length expected 3 actual %v", len(output))
	}
}

func TestFindCommonIdPart(t *testing.T) {
	commonIdPart := FindCommonIdPart(inputSlice)

	if commonIdPart != "fgij" {
		t.Fatalf("Output expected fgij actual %s", commonIdPart)
	}

	if len(commonIdPart) != 4 {
		t.Fatalf("Output length expected 4 actual %v", len(commonIdPart))
	}

}
