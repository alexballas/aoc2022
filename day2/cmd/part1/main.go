package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	ROCK = iota + 1
	PAPER
	SCISSORS
)

const (
	LOSE = 0
	DRAW = 3
	WIN  = 6
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	buf := bufio.NewScanner(f)
	var score int

	for buf.Scan() {
		game := strings.Split(buf.Text(), " ")
		if len(game) != 2 {
			panic("bad data")
		}

		switch game[0] {
		case "A":
			switch game[1] {
			case "X":
				score += ROCK + DRAW
			case "Y":
				score += PAPER + WIN
			case "Z":
				score += SCISSORS + LOSE
			}
		case "B":
			switch game[1] {
			case "X":
				score += ROCK + LOSE
			case "Y":
				score += PAPER + DRAW
			case "Z":
				score += SCISSORS + WIN
			}
		case "C":
			switch game[1] {
			case "X":
				score += ROCK + WIN
			case "Y":
				score += PAPER + LOSE
			case "Z":
				score += SCISSORS + DRAW
			}
		}
	}

	fmt.Println("Part 1:", score)
}
