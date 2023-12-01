package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

func parseLine(line string) []string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAllString(line, -1)
}

var numbersInText = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

var replacement = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9", "1": "1", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7", "8": "8", "9": "9"}

func getIndex(line string, value string) int {
	return strings.Index(line, value)
}

func prepareData(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := 0
	for scanner.Scan() {
		s := scanner.Text()
		data += parseLine_part2(s)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func parseLine_part2(line string) int {
	searches := []struct {
		name string
		val  string
	}{{name: "one", val: "1"},
		{name: "two", val: "2"},
		{name: "three", val: "3"},
		{name: "four", val: "4"},
		{name: "five", val: "5"},
		{name: "six", val: "6"},
		{name: "seven", val: "7"},
		{name: "eight", val: "8"},
		{name: "nine", val: "9"},
		{name: "1", val: "1"},
		{name: "2", val: "2"},
		{name: "3", val: "3"},
		{name: "4", val: "4"},
		{name: "5", val: "5"},
		{name: "6", val: "6"},
		{name: "7", val: "7"},
		{name: "8", val: "8"},
		{name: "9", val: "9"}}

	var first, last string
	for i := 0; i < len(line); i++ {
		for _, search := range searches {
			if strings.HasPrefix(line[i:], search.name) {
				if first == "" {
					first = search.val
					last = search.val
				} else {
					last = search.val
				}
			}
		}
	}

	calibration, _ := strconv.Atoi(string(first) + string(last))
	return calibration
}

func main() {
	// 53221 valid
	// 53188 to low
	fmt.Println(prepareData("input"))
}
