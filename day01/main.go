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
		line := parseLine(s)
		numbers := strings.Join(line, "")
		if len(numbers) > 1 {
			charAtString := fmt.Sprintf("%c%c", numbers[0], numbers[len(numbers)-1])
			lineNumber, err := strconv.Atoi(charAtString)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("x:", lineNumber)
			data += lineNumber
		} else {
			number := fmt.Sprintf("%c%c", numbers[0], numbers[0])
			lineNumber, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("y", lineNumber)
			data += lineNumber
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func main() {
	fmt.Println(prepareData("input"))
}
