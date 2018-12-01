package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	// Read file
	fileName := "input"
	inputAsByteArray, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("File %s not found. Exiting. \n", fileName)
		panic(err)
	}

	// Convert to string slice.
	stringSlice := strings.Split(string(inputAsByteArray), "\n")

	sum := 0

	// Calculate sum
	for i, item := range stringSlice {
		intItem, err := strconv.Atoi(item)
		if err == nil {
			sum += intItem
		} else {
			log.Printf("Unable to convert %s to int. at line %d \n", item, i)
		}
	}

	// Print result
	fmt.Printf("Result is %v \n", sum)
}
