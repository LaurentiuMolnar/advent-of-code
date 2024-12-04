package solutions2024

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var multiplicationRegex = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
var part2Regex = regexp.MustCompile(`(mul\((\d+),(\d+)\))|(do\(\))|(don\'t\(\))`)

func Day03Part1() {
	input := utils.ReadInputLines()

	var sum int64 = 0
	for _, line := range input {
		matches := multiplicationRegex.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			term1, _ := strconv.Atoi(match[1])
			term2, _ := strconv.Atoi(match[2])

			sum += int64(term1) * int64(term2)
		}
	}

	fmt.Println("Result:", sum)
}

func Day03Part2() {
	input := utils.ReadInputLines()

	var sum int64 = 0
	enabled := true
	for _, line := range input {
		matches := part2Regex.FindAllStringSubmatch(line, -1)
		// fmt.Println(matches)

		for _, match := range matches {
			instruction := match[0]
			if strings.HasPrefix(instruction, "mul") {
				if enabled {
					term1, _ := strconv.Atoi(match[2])
					term2, _ := strconv.Atoi(match[3])
					sum += int64(term1) * int64(term2)
				}
			} else if strings.HasPrefix(instruction, "do()") {
				enabled = true
			} else if strings.HasPrefix(instruction, "don't()") {
				enabled = false
			}
		}
	}

	fmt.Println("Result:", sum)
}
