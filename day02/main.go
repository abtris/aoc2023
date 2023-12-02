package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type play struct {
	blue  int
	red   int
	green int
}

type gameset struct {
	name string
	sets []play
}

type games struct {
	gamesets []gameset
}

func parseGame(s string) gameset {
	var g gameset
	parts := strings.Split(s, ":")
	g.name = parts[0]
	plays := strings.Split(parts[1], ";")
	for _, p := range plays {
		var ourset play
		items := strings.Split(p, ",")
		for _, v := range items {
			if strings.HasSuffix(v, "blue") {
				s, _ := strings.CutSuffix(v, "blue")
				s = strings.Trim(s, " ")
				v, _ := strconv.Atoi(s)
				ourset.blue = v
			}
			if strings.HasSuffix(v, "green") {
				s, _ := strings.CutSuffix(v, "green")
				s = strings.Trim(s, " ")
				v, _ := strconv.Atoi(s)
				ourset.green = v
			}
			if strings.HasSuffix(v, "red") {
				s, _ := strings.CutSuffix(v, "red")
				s = strings.Trim(s, " ")
				v, _ := strconv.Atoi(s)
				ourset.red = v
			}
		}
		g.sets = append(g.sets, ourset)
	}
	return g
}

func prepareData(filename string) games {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data games
	for scanner.Scan() {
		s := scanner.Text()
		data.gamesets = append(data.gamesets, parseGame(s))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func part1() {

}

func main() {
	prepareData("input_test")
}
