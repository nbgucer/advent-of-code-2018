package main

import (
	"fmt"
	"github.com/nbgucer/advent-of-code-2018/utils"
)

func main() {

	fileName := "days\\day-2\\input"
	stringSlice := utils.GetInputAsSlice(fileName)
	totalTwoCount := 0
	totalThreeCount := 0
	for _, item := range stringSlice {
		twos, threes := LetterCounter(item)
		totalTwoCount += twos
		totalThreeCount += threes
	}

	fmt.Printf("Result is %v", totalTwoCount*totalThreeCount)
}

func LetterCounter(input string) (twos int, threes int) {
	letterMap := make(map[rune]int)
	twoCount := 0
	threeCount := 0
	for _, char := range input {
		_, ok := letterMap[char]
		if ok {
			letterMap[char] += 1
			if letterMap[char] == 1 {
				// do nothing.
			} else if letterMap[char] == 2 {
				twoCount++
			} else if letterMap[char] == 3 {
				twoCount--
				threeCount++
			} else if letterMap[char] == 4 {
				threeCount--
			} else {
				// do nothing.
			}
		} else {
			letterMap[char] = 1
		}
	}

	hasTwo := 0
	hasThree := 0

	if twoCount > 0 {
		hasTwo = 1
	}
	if threeCount > 0 {
		hasThree = 1
	}

	return hasTwo, hasThree
}
