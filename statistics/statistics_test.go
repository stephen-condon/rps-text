package statistics

import (
	"reflect"
	"testing"
)

func TestCanary(t *testing.T) {

}

func TestNew(t *testing.T) {
	want := GameStatistics{
		win:  0,
		loss: 0,
		tie:  0,
	}
	New()

	if !reflect.DeepEqual(gameStats, want) {
		t.Errorf(`Got: %v, received: %v`, gameStats, want)
	}
}

func TestGet(t *testing.T) {
	want := &GameStatistics{
		win:  0,
		loss: 0,
		tie:  0,
	}
	New()

	if !reflect.DeepEqual(Get(), want) {
		t.Errorf(`Got: %v, received: %v`, Get(), want)
	}
}

func TestGameStatistics(t *testing.T) {
	type test struct {
		result string
		want   GameStatistics
	}

	tests := []test{
		{"win", GameStatistics{1, 0, 0}},
		{"win", GameStatistics{2, 0, 0}},
		{"loss", GameStatistics{2, 1, 0}},
		{"loss", GameStatistics{2, 2, 0}},
		{"tie", GameStatistics{2, 2, 1}},
		{"abcdefg", GameStatistics{2, 2, 1}},
		{"win", GameStatistics{3, 2, 1}},
	}

	New()

	for _, tc := range tests {
		gameStats.Increment(tc.result)
		if !reflect.DeepEqual(gameStats, tc.want) {
			t.Errorf(`Got: %v, received: %v`, gameStats, tc.want)
		}
	}
}
