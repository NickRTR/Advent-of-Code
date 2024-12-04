// Exercise: https://adventofcode.com/2024/day/4

package main

import (
	"fmt"
	"log"
	"os"
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

func checkCoordinate(y int, x int, grid [140][]string, search string) bool {
	if y < 0 || y > len(grid[0])-1 || x < 0 || x > len(grid)-1 {
		return false
	}
	if grid[y][x] == search {
		return true
	}
	return false
}

func checkXMAS(y int, x int, grid [140][]string) int {
	coordinates := [][][]int{
		{ // LEFT
			{y, x - 1},
			{y, x - 2},
			{y, x - 3},
		},
		{ // RIGHT
			{y, x + 1},
			{y, x + 2},
			{y, x + 3},
		},
		{ // UP
			{y + 1, x},
			{y + 2, x},
			{y + 3, x},
		},
		{ // UP-Left
			{y + 1, x - 1},
			{y + 2, x - 2},
			{y + 3, x - 3},
		},
		{ // UP-RIGHT
			{y + 1, x + 1},
			{y + 2, x + 2},
			{y + 3, x + 3},
		},
		{ // DOWN
			{y - 1, x},
			{y - 2, x},
			{y - 3, x},
		},
		{ // DOWN-LEFT
			{y - 1, x - 1},
			{y - 2, x - 2},
			{y - 3, x - 3},
		},
		{ // DOWN-RIGHT
			{y - 1, x + 1},
			{y - 2, x + 2},
			{y - 3, x + 3},
		},
	}

	counter := 0
	for _, directions := range coordinates {
		if checkCoordinate(directions[0][0], directions[0][1], grid, "M") && checkCoordinate(directions[1][0], directions[1][1], grid, "A") && checkCoordinate(directions[2][0], directions[2][1], grid, "S") {
			counter++
		}
	}
	return counter
}

func checkCrossMAS(y int, x int, grid [140][]string) bool {
	corners := [][]int{
		{y + 1, x + 1},
		{y + 1, x - 1},
		{y - 1, x + 1},
		{y - 1, x - 1},
	}

	if checkCoordinate(corners[0][0], corners[0][1], grid, "M") && checkCoordinate(corners[3][0], corners[3][1], grid, "S") ||
		checkCoordinate(corners[0][0], corners[0][1], grid, "S") && checkCoordinate(corners[3][0], corners[3][1], grid, "M") {
		if checkCoordinate(corners[1][0], corners[1][1], grid, "M") && checkCoordinate(corners[2][0], corners[2][1], grid, "S") ||
			checkCoordinate(corners[1][0], corners[1][1], grid, "S") && checkCoordinate(corners[2][0], corners[2][1], grid, "M") {
			return true

		}
	}
	return false
}

func part1(grid [140][]string) (int, time.Duration) {
	start := time.Now()

	counter := 0
	for y, row := range grid {
		for x, column := range row {
			if column == "X" {
				counter += checkXMAS(y, x, grid)
			}
		}
	}

	duration := time.Since(start)

	return counter, duration
}

func part2(grid [140][]string) (int, time.Duration) {
	start := time.Now()

	counter := 0
	for y, row := range grid {
		for x, column := range row {
			if column == "A" {
				if checkCrossMAS(y, x, grid) {
					counter++
				}
			}
		}
	}

	duration := time.Since(start)

	return counter, duration
}

func main() {
	start := time.Now()

	file := readInput("input")
	lines := strings.Split(file, "\n")

	var grid [140][]string

	for y, line := range lines {
		characters := strings.Split(line, "")
		grid[y] = characters
	}

	part1, part1Duration := part1(grid)
	part2, part2Duration := part2(grid)

	duration := time.Since(start)

	fmt.Printf("✅ Part 1: 'XMAS' appears %v times in the word search (Execution time: %s)\n", part1, part1Duration)
	fmt.Printf("✅ Part 2: 'X-MAS' appears %v times in the word search (Execution time: %s)\n", part2, part2Duration)
	fmt.Printf("🚀 Total execution time: %s\n", duration)
}
