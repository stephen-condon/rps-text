package game

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var winMap map[string]string

func Start() {
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
			computerSelection := computerTurn(generateRawComputerChoice)
			result := evaluateGame(selection, computerSelection)
			fmt.Println(evaluateResult(result))
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

func computerTurn(genFn generateChoice) string {
	var selection string

	rawRand := genFn()
	if rawRand == 0 {
		selection = "rock"
	} else if rawRand == 1 {
		selection = "paper"
	} else {
		selection = "scissors"
	}

	return selection
}

func generateRawComputerChoice() int {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	return generator.Intn(3)
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

func evaluateResult(result int) string {
	str := "Tie!"
	if result == 1 {
		str = "Player Won!"
	} else if result == -1 {
		str = "Player Lost!"
	}

	return str
}

func setWinMap() {
	winMap = map[string]string{
		"rock":     "scissors",
		"paper":    "rock",
		"scissors": "paper",
	}
}

type generateChoice func() int
