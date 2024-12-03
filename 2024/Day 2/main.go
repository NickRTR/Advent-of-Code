// Exercise: https://adventofcode.com/2024/day/2

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput(file string) string {
	rawFile, err := os.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func checkReport(report []int) bool {
	increasing := report[1] > report[0]

	for i := 1; i < len(report); i++ {
		a := report[i-1]
		b := report[i]

		if increasing {
			if b <= a || b > a+3 {
				return false
			}
		} else {
			if b >= a || b < a-3 {
				return false
			}
		}

	}
	return true
}

func reportVariations(report []int) bool {
	for i := 0; i < len(report); i++ {
		modifiedReport := make([]int, len(report))
		copy(modifiedReport, report)
		modifiedReport = append(modifiedReport[:i], modifiedReport[i+1:]...)
		if checkReport(modifiedReport) {
			return true
		}
	}
	return false
}

func part1(reports [][]int) (int, time.Duration) {
	start := time.Now()

	saveReports := 0

	for _, report := range reports {
		if checkReport(report) {
			saveReports++
		}
	}

	duration := time.Since(start)

	return saveReports, duration
}

func part2(reports [][]int) (int, time.Duration) {
	start := time.Now()

	saveReports := 0

	for _, report := range reports {
		if checkReport(report) {
			saveReports++
		} else {
			if reportVariations(report) {
				saveReports++
			}
		}
	}

	duration := time.Since(start)

	return saveReports, duration
}

func main() {
	start := time.Now()

	file := readInput("input")
	lines := strings.Split(file, "\n")
	reports := make([][]int, len(lines))
	for i, line := range lines {
		levels := strings.Split(line, " ")
		reports[i] = make([]int, len(levels))
		for z, number := range levels {
			converted, _ := strconv.Atoi(number)
			reports[i][z] = converted
		}
	}

	part1, part1Duration := part1(reports)
	part2, part2Duration := part2(reports)

	duration := time.Since(start)

	fmt.Printf("✅ Part 1: The number of safe nuclear reports is: %v (Execution time: %s)\n", part1, part1Duration)
	fmt.Printf("✅ Part 2: The number of safe nuclear reports considering the problem dampener is: %v (Execution time: %s)\n", part2, part2Duration)
	fmt.Printf("🚀 Total execution time: %s\n", duration)
}
