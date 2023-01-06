// Exercise: https://adventofcode.com/2022/day/10

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readInput(file string) string {
	// read input
	rawFile, err := ioutil.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func part1(instructions []string) {
	cycles := make(map[int]int)

	cycle := 0
	x := 1

	for _, i := range instructions {
		if len(i) > 0 {
			if i == "noop" {
				cycle++
				cycles[cycle] = x
			} else {
				cycles[cycle+1] = x
				cycles[cycle+2] = x
				cycle += 2
				// get value from addx instruction
				value, _ := strconv.Atoi(strings.Split(i, " ")[1])
				x += value
			}
		}
	}

	signalStrengthSum := 0

	for i := 20; i <= 220; i += 40 {
		signalStrengthSum += i * cycles[i]
	}

	fmt.Printf("The sum of all signal strengths is %v\n", signalStrengthSum)
}

func main() {
	input := readInput("input.txt")
	instructions := strings.Split(input, "\n")

	part1(instructions)
}
