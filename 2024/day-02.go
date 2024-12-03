package solutions2024

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
)

var regex = regexp.MustCompile(`(\d+)`)

func getNums(input []string) [][]int {
	var numbers = make([][]int, len(input))
	for i, line := range input {
		matches := regex.FindAllString(line, -1)
		numbers[i] = make([]int, len(matches))

		for j, match := range matches {
			num, _ := strconv.Atoi(match)
			numbers[i][j] = num
		}
	}

	return numbers
}

func isValidReport(report []int) bool {
	ascendingCmp := func(a, b int) int {
		dif := math.Abs(float64(a - b))

		if dif < 1 || dif > 3 || a > b {
			return -1
		}

		if a == b {
			return 0
		}

		return 1
	}

	descendingCmp := func(a, b int) int {
		dif := math.Abs(float64(a - b))

		if dif < 1 || dif > 3 || a < b {
			return -1
		}

		if a == b {
			return 0
		}

		return 1
	}

	return slices.IsSortedFunc(report, ascendingCmp) || slices.IsSortedFunc(report, descendingCmp)
}

func Day02Part1() {
	input := utils.ReadInputLines()
	reports := getNums(input)

	count := 0
	for _, report := range reports {
		if isValidReport(report) {
			count++
		}
	}

	fmt.Println("Result:", count)
}

func Day02Part2() {
	input := utils.ReadInputLines()
	reports := getNums(input)

	fmt.Println(reports)
}
