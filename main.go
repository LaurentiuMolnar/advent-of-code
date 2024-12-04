package main

import (
	solutions2024 "advent-of-code/2024"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Usage: main year day part [options]. Options: --sample or -s - Use sample input instead of the real input")
	}

	year, _ := strconv.Atoi(os.Args[1])
	day, _ := strconv.Atoi(os.Args[2])
	part, _ := strconv.Atoi(os.Args[3])

	functions := map[int]map[int]map[int]func(){
		2024: {
			1: {
				1: solutions2024.Day01Part1,
				2: solutions2024.Day01Part2,
			},
			2: {
				1: solutions2024.Day02Part1,
				2: solutions2024.Day02Part2,
			},
			3: {
				1: solutions2024.Day03Part1,
				2: solutions2024.Day03Part2,
			},
		},
	}

	fmt.Printf("Running script for AOC %v Day %v Part %v\n", year, day, part)
	(func() {
		start := time.Now()
		functions[year][day][part]()
		fmt.Printf("Time: %v\n", time.Since(start))
	})()
}
