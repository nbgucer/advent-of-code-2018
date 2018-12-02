package main

import (
	"fmt"
	"github.com/nbgucer/advent-of-code-2018/utils"
	"log"
	"strconv"
)

func main() {
	sum := 0
	result := 0
	found := false
	seenFrequencies := make(map[int]int)
	seenFrequencies[0] = 1
	fileName := "days\\day-1\\input"

	stringSlice := utils.GetInputAsSlice(fileName)

	// Calculate sum
	for found == false {
		for i, item := range stringSlice {
			intItem, err := strconv.Atoi(item)
			if err == nil {
				sum += intItem
				_, ok := seenFrequencies[sum]
				if ok {
					result = sum
					found = true
					break
				} else {
					seenFrequencies[sum] = 1
				}
			} else if len(item) == 0 {
				//Empty string, skip
				log.Printf("Empty string. Skipping at line %d \n", i)
			} else {
				log.Printf("Unable to convert %s to int. at line %d \n", item, i)
			}
		}
	}

	// Print result
	fmt.Printf("Result is %d \n", result)
}
