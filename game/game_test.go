package game

import "testing"

func TestCanary(t *testing.T) {

}

func TestComputerTurn(t *testing.T) {
	type test struct {
		genFn generateChoice
		want  string
	}
	tests := []test{
		{func() int { return 0 }, "rock"},
		{func() int { return 1 }, "paper"},
		{func() int { return 2 }, "scissors"},
		{func() int { return 999 }, "scissors"},
	}

	for _, tc := range tests {
		got := computerTurn(tc.genFn)
		if got != tc.want {
			t.Errorf(`Received %v but expected %v`, got, tc.want)
		}
	}
}

func TestEvaluateGame(t *testing.T) {
	type test struct {
		player, computer string
		want             int
	}
	tests := []test{
		{"rock", "rock", 0},
		{"rock", "paper", -1},
		{"rock", "scissors", 1},
		{"paper", "rock", 1},
		{"paper", "paper", 0},
		{"paper", "scissors", -1},
		{"scissors", "rock", -1},
		{"scissors", "paper", 1},
		{"scissors", "scissors", 0},
	}

	setWinMap()

	for _, tc := range tests {
		got := evaluateGame(tc.player, tc.computer)
		if got != tc.want {
			t.Errorf(`player: %v, computer: %v; Received %v but expected %v`, tc.player, tc.computer, got, tc.want)
		}
	}
}

func TestEvaluateResult(t *testing.T) {
	type test struct {
		input int
		want  string
	}
	tests := []test{
		{1, "Player Won!"},
		{-1, "Player Lost!"},
		{0, "Tie!"},
		{999, "Tie!"},
	}

	for _, tc := range tests {
		got := evaluateResult(tc.input)
		if got != tc.want {
			t.Errorf(`input: %v; Received %v but expected %v`, tc.input, got, tc.want)
		}
	}
}

func TestSetWinMap(t *testing.T) {
	wantKeys := []string{"rock", "paper", "scissors"}
	wantValues := []string{"scissors", "rock", "paper"}

	setWinMap()

	for i, key := range wantKeys {
		if winMap[key] != wantValues[i] {
			t.Errorf(`key: %v; Received value %v, but expected %v`, key, winMap[key], wantValues[i])
		}
	}
}
