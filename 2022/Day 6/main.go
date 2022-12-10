// Exercise: https://adventofcode.com/2022/day/6

package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

func check(input []string) bool {
	// check for marker

	counter := make(map[string]bool)
	for _, c := range input {
		if counter[c] == true {
			return false
		}
		counter[c] = true
	}
	return true
}

func main() {
	input := strings.Split(readInput("input.txt"), "")

	if len(input) < 4 {
		fmt.Println("Input too short!")
		return
	}

	characters := input[0:4]

	for i := 0; i < len(input)-3; i++ {
		characters = input[i : i+4]
		if check(characters) {
			fmt.Printf("%v characters have to be processed to find the first start-of-packet marker.\n", i+4)
			return
		}
	}

	fmt.Println("No markers found.")
}
