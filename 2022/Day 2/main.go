// Exercise: https://adventofcode.com/2022/day/2

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// 		 Opponent | Player
// Rock: 		A | X
// Paper: 		B | Y
// Scissors: 	C | Z

func readInput(file string) string {
	// read input
	rawFile, err := ioutil.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func decode(action string) string {
	if action == "A" || action == "X" {
		return "Rock"
	} else if action == "B" || action == "Y" {
		return "Paper"
	} else {
		return "Scissors"
	}
}

func encode(action string) string {
	switch action {
	case "Rock":
		return "X"
	case "Paper":
		return "Y"
	default:
		return "Z"
	}
}

func selectionScore(selection string) int {
	var score int

	// set score for selection
	switch selection {
	case "X":
		score = 1
	case "Y":
		score = 2
	case "Z":
		score = 3
	}

	return score
}

func play(opponent, player string) int {
	opponent = decode(opponent)
	player = decode(player)

	score := selectionScore(player)

	// 0 for loss, 3 for draw, 6 for win
	if opponent == player {
		// draw
		score += 3
	} else if opponent == "Rock" && player == "Paper" || opponent == "Paper" && player == "Scissors" || opponent == "Scissors" && player == "Rock" {
		// win
		score += 6
	} else {
		// loss
		score += 0
	}

	return score
}

func playWithResult(opponent, result string) int {
	// X: loss
	// Y: draw
	// Z: win

	opponent = decode(opponent)

	var selection string

	// select right action for result
	switch result {
	case "X":
		switch opponent {
		case "Rock":
			selection = encode("Scissors")
		case "Paper":
			selection = encode("Rock")
		case "Scissors":
			selection = encode("Paper")
		}
	case "Y":
		selection = encode(opponent)
	case "Z":
		switch opponent {
		case "Rock":
			selection = encode("Paper")
		case "Paper":
			selection = encode("Scissors")
		case "Scissors":
			selection = encode("Rock")
		}
	}

	score := selectionScore(selection)

	switch result {
	case "Y":
		score += 3
	case "Z":
		score += 6
	}

	return score
}

func main() {
	input := readInput("input.txt")
	rounds := strings.Split(input, "\n")

	totalScore1 := 0
	totalScore2 := 0

	for _, r := range rounds {
		if len(r) == 3 {
			opponent := r[0]
			player := r[2]
			totalScore1 += play(string(opponent), string(player))
			totalScore2 += playWithResult(string(opponent), string(player))
		}
	}

	fmt.Printf("The total score with the second row as the player's action will be %v.\n", totalScore1)
	fmt.Printf("The total score with the second row as the round's result will be %v.\n", totalScore2)
}
