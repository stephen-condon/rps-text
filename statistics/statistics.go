package statistics

import "fmt"

var gameStats GameStatistics

func New() {
	gameStats = GameStatistics{
		win:  0,
		loss: 0,
		tie:  0,
	}
}

func Get() *GameStatistics {
	return &gameStats
}

func PrintStatistics() {
	fmt.Println("Current Statistics")
	fmt.Println("Wins: ", Get().win)
	fmt.Println("Losses: ", Get().loss)
	fmt.Println("Ties: ", Get().tie)
}

func (gs *GameStatistics) Increment(result string) {
	if result == "win" {
		gs.win = gs.win + 1
	} else if result == "loss" {
		gs.loss = gs.loss + 1
	} else if result == "tie" {
		gs.tie = gs.tie + 1
	}
}

type GameStatistics struct {
	win, loss, tie int
}
