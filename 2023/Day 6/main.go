// Exercise: https://adventofcode.com/2023/day/6

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInput(file string) string {
	rawFile, err := os.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func part1(lines []string) int {
	// get times and distances from input
	var distances []int
	var times []int

	timeString := strings.Split(lines[0], " ")
	for _, game := range timeString {
		time, _ := strconv.Atoi(game)
		if time != 0 {
			times = append(times, time)
		}
	}

	distanceString := strings.Split(lines[1], " ")
	for _, game := range distanceString {
		distance, _ := strconv.Atoi(game)
		if distance != 0 {
			distances = append(distances, distance)
		}
	}

	sum := 1

	for i := 0; i < len(times); i++ {
		winCounter := 0

		time := times[i]
		distance := distances[i]

		for t := 1; t < time; t++ {
			d := t * (time - t)
			if d > distance {
				winCounter++
			}
		}

		sum *= winCounter
	}
	return sum
}

func part2(lines []string) int {
	var time int
	var distance int

	for i, line := range lines {
		line = regexp.MustCompile(`\D`).ReplaceAllString(line, "")

		converted, _ := strconv.Atoi(line)
		if i == 0 {
			time = converted
		} else {
			distance = converted
		}
	}

	counter := 0

	for t := 1; t < time; t++ {
		d := t * (time - t)
		if d > distance {
			counter++
		}
	}

	return counter
}

func main() {
	file := readInput("input")

	lines := strings.Split(file, "\n")

	fmt.Println("The product of all possible ways to get the record in each race is:", part1(lines))
	fmt.Println("The possible ways to get the record in the long race is:", part2(lines))
}
