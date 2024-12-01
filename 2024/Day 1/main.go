// Exercise: https://adventofcode.com/2024/day/1

package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
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

func part1(leftIDs []int, rightIDs []int) (int, time.Duration) {
	start := time.Now()

	slices.Sort(leftIDs)
	slices.Sort(rightIDs)

	totalDistance := 0

	for i := 0; i < len(leftIDs); i++ {
		distance := math.Abs(float64(leftIDs[i]) - float64(rightIDs[i]))
		totalDistance += int(distance)
	}

	duration := time.Since(start)

	return totalDistance, duration
}

func part2(leftIDs []int, rightIDs []int) (int, time.Duration) {
	start := time.Now()

	slices.Sort(leftIDs)
	slices.Sort(rightIDs)

	similarityScore := 0

	for _, id := range leftIDs {
		index := slices.Index(rightIDs, id)
		if index != -1 {
			count := 0
			for rightIDs[index] == id {
				count++
				index++
			}
			similarityScore += id * count
		}
	}

	duration := time.Since(start)

	return similarityScore, duration
}

func main() {
	start := time.Now()

	file := readInput("input")

	var leftIDs []int
	var rightIDs []int

	lines := strings.Split(file, "\n")

	for _, line := range lines {
		ids := strings.Split(line, "   ")
		left, _ := strconv.Atoi(ids[0])
		right, _ := strconv.Atoi(ids[1])
		leftIDs = append(leftIDs, left)
		rightIDs = append(rightIDs, right)
	}

	part1, part1Duration := part1(leftIDs, rightIDs)
	part2, part2Duration := part2(leftIDs, rightIDs)

	duration := time.Since(start)

	fmt.Printf("✅ Part 1: The total distance between both lists is: %v (Execution time: %s)\n", part1, part1Duration)
	fmt.Printf("✅ Part 2: The similarity score of both lists is: %v (Execution time: %s)\n", part2, part2Duration)
	fmt.Printf("🚀 Total execution time: %s\n", duration)
}
