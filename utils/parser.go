package utils

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func GetInputAsSlice(fileName string) []string {
	// Read file
	absPath, _ := filepath.Abs(fileName)
	inputAsByteArray, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Panicf("File %s not found. Exiting. \n", fileName)
		panic(err)
	}
	// Convert to string slice.
	stringSlice := strings.Split(string(inputAsByteArray), "\n")
	// Remove empty line at the end.
	stringSlice = stringSlice[:len(stringSlice)-1]
	return stringSlice
}
