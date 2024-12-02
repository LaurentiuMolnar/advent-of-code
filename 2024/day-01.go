package solutions2024

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func getSortedLists(lines []string) ([]int, []int) {
	var left, right []int

	for _, line := range lines {
		parts := strings.Split(line, " ")
		num, _ := strconv.Atoi(parts[0])
		left = append(left, num)
		num, _ = strconv.Atoi(parts[len(parts)-1])
		right = append(right, num)
	}

	slices.Sort(left)
	slices.Sort(right)

	return left, right
}

func Day01Part1() {
	input := utils.ReadInputLines()
	left, right := getSortedLists(input)

	distance := 0
	for i := range left {
		distance += int(math.Abs(float64(left[i] - right[i])))
	}
	fmt.Println("Result:", distance)
}

func Day01Part2() {
	input := utils.ReadInputLines()
	left, right := getSortedLists(input)
	var alreadyCounted map[int]int = make(map[int]int)

	similarity := 0
	lastPos := 0
	for i := range left {
		if val, ok := alreadyCounted[left[i]]; ok {
			similarity += left[i] * val
			continue
		}

		count := 0
		for j := lastPos; left[i] >= right[j]; j++ {
			if left[i] == right[j] {
				count++
			}
			lastPos = j
		}
		alreadyCounted[left[i]] = count
		similarity += left[i] * count
	}
	fmt.Println("Result:", similarity)
}
