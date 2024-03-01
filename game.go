package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var winMap map[string]string

func main() {
	log.Println("Welcome to the game")
	fmt.Println("Press s to start or q to quit")

	var startInput string
	fmt.Scanln(&startInput)
	if startInput == "q" {
		os.Exit(0)
	} else if startInput == "s" {
		fmt.Println("Pressed s")
		executeGame()
	}
}

func executeGame() {
	quit := false
	setWinMap()

	// main game loop
	for !quit {
		selection := playerTurn()

		if selection == "quit" {
			quit = true
		} else {
			computerSelection := computerTurn()
			result := evaluateGame(selection, computerSelection)
			printResult(result)
		}

	}
}

func playerTurn() string {
	selection := ""
	fmt.Println("Make your selection")
	fmt.Println("r: Rock | s: Scissors | p: Paper | q: Quit the Game")
	var input string

	fmt.Scanln(&input)
	if input == "q" {
		selection = "quit"
	} else if input == "r" {
		selection = "rock"
	} else if input == "p" {
		selection = "paper"
	} else if input == "s" {
		selection = "scissors"
	}

	return selection
}

func computerTurn() string {
	// computer picks
	// do random, convert to choice
	var selection string
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	rawRand := generator.Intn(3)
	if rawRand == 0 {
		selection = "rock"
	} else if rawRand == 1 {
		selection = "paper"
	} else {
		selection = "scissors"
	}

	return selection
}

func evaluateGame(player, computer string) int {
	result := 0

	if player != computer {
		playerWinValue := winMap[player]
		if computer == playerWinValue {
			result = 1
		} else {
			result = -1
		}
	}

	return result
}

func printResult(result int) {
	if result == 1 {
		fmt.Println("Player Won!")
	} else if result == -1 {
		fmt.Println("Player Lost!")
	} else {
		fmt.Println("Tie!")
	}
}

func setWinMap() {
	winMap = map[string]string{
		"rock":     "scissors",
		"paper":    "rock",
		"scissors": "paper",
	}
}
