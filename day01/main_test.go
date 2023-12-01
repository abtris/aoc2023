package main

import "testing"

func TestParseLine(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result string
	}{
		{name: "eighthree", input: "eighthree", result: "83"},
		{name: "sevenine", input: "sevenine", result: "79"},
		{name: "eightwothree", input: "eightwothree", result: "83"},
		{name: "abcone2threexyz", input: "abcone2threexyz", result: "13"},
		{name: "xtwone3four", input: "xtwone3four", result: "24"},
		{name: "4nineeightseven2", input: "4nineeightseven2", result: "42"},
		{name: "zoneight234", input: "zoneight234", result: "14"},
		{name: "7pqrstsixteen", input: "7pqrstsixteen", result: "76"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := parseLine_part2(test.input)
			if result != test.result {
				t.Errorf("Expected %v and real %v)", test.result, result)
			}
		})
	}
}
