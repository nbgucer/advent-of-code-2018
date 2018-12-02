package main

import (
	"fmt"
	"github.com/nbgucer/advent-of-code-2018/utils"
)

func main() {

	fileName := "days\\day-2\\input"
	stringSlice := utils.GetInputAsSlice(fileName)
	commonIdPart := FindCommonIdPart(stringSlice)

	fmt.Printf("\n Result is %v with a closeness of %v within %v \n", commonIdPart, len(commonIdPart), len(stringSlice[0]))
}

func FindCommonIdPart(stringSlice []string) string {
	closeness := 0
	commonIdPart := ""
	for i := 0; i < len(stringSlice); i++ {
		for j := i + 1; j < len(stringSlice); j++ {
			ResultCommonIdPart := FindCloseness(stringSlice[i], stringSlice[j])
			resultCloseness := len(ResultCommonIdPart)
			if resultCloseness > closeness {
				commonIdPart = ResultCommonIdPart
				closeness = resultCloseness
			}
		}
	}
	return commonIdPart
}

func FindCloseness(a string, b string) string {
	output := ""
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			output += string(a[i])
		}
	}

	return output
}
