package main

import (
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
		2024: {},
	}

	fmt.Printf("Running script for AOC %v Day %v Part %v\n", year, day, part)
	(func() {
		start := time.Now()
		functions[year][day][part]()
		fmt.Printf("Time: %v\n", time.Since(start))
	})()
}
