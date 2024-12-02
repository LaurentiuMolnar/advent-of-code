package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadInputLines() []string {
	year := os.Args[1]
	day, _ := strconv.Atoi(os.Args[2])
	lastArg := os.Args[len(os.Args)-1]
	useSample := lastArg == "--sample" || lastArg == "-s"

	var inputFileName string
	if useSample {
		inputFileName = "sample"
	} else {
		inputFileName = "input"
	}

	file, err := os.Open(fmt.Sprintf("./%s/day-%02d/%s", year, day, inputFileName))

	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
