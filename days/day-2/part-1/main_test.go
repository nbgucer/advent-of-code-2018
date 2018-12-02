package main

import "testing"

var (
	input1 string = "abcdef"
	input2 string = "bababc"
	input3 string = "abbcde"
	input4 string = "abcccd"
	input5 string = "aabcdd"
	input6 string = "abcdee"
	input7 string = "ababab"
)

func TestLineContains0Double0Triple(t *testing.T) {
	twos, threes := LetterCounter(input1)

	if twos != 0 {
		t.Fatalf("expecting 0 for twos, got %v", twos)
	}

	if threes != 0 {
		t.Fatalf("expecting 0 for threes, got %v", threes)
	}
}

func TestLineContainsDoubleLetterOnceTripleLetterOnce(t *testing.T) {
	twos, threes := LetterCounter(input2)

	if twos != 1 {
		t.Fatalf("expecting 1 for twos, got %v", twos)
	}

	if threes != 1 {
		t.Fatalf("expecting 1 for threes, got %v", threes)
	}
}

func TestLineContainsDoubleLetterOnceTripleLetterZero(t *testing.T) {
	twos, threes := LetterCounter(input3)

	if twos != 1 {
		t.Fatalf("expecting 1 for twos, got %v", twos)
	}

	if threes != 0 {
		t.Fatalf("expecting 0 for threes, got %v", threes)
	}
}

func TestLineContainsDoubleLetterZeroTripleLetterOnce(t *testing.T) {
	twos, threes := LetterCounter(input4)

	if twos != 0 {
		t.Fatalf("expecting 0 for twos, got %v", twos)
	}

	if threes != 1 {
		t.Fatalf("expecting 1 for threes, got %v", threes)
	}
}

func TestLineContainsDoubleDoubleTwosExpectingTwosOne(t *testing.T) {
	twos, threes := LetterCounter(input5)

	if twos != 1 {
		t.Fatalf("expecting 1 for twos, got %v", twos)
	}

	if threes != 0 {
		t.Fatalf("expecting 0 for threes, got %v", threes)
	}
}

func TestLineContainsTwoTwosZeroThrees(t *testing.T) {
	twos, threes := LetterCounter(input6)

	if twos != 1 {
		t.Fatalf("expecting 1 for twos, got %v", twos)
	}

	if threes != 0 {
		t.Fatalf("expecting 0 for threes, got %v", threes)
	}
}

func TestLineContainsThreeThrees(t *testing.T) {
	twos, threes := LetterCounter(input7)

	if twos != 0 {
		t.Fatalf("expecting 0 for twos, got %v", twos)
	}

	if threes != 1 {
		t.Fatalf("expecting 1 for threes, got %v", threes)
	}
}
