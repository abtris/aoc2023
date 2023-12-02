package main

import (
	"testing"
)

func TestParseGame(t *testing.T) {
	tests := []struct {
		name   string
		x      string
		result gameset
	}{
		{name: "simple", x: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", result: gameset{name: "Game 1", sets: []play{
			play{blue: 3, red: 4, green: 0},
			play{blue: 6, red: 1, green: 2},
			play{blue: 0, red: 0, green: 2},
		},
		},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := parseGame(test.x)
			t.Logf("%v\n", result)
		})
	}
}
