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

	// According to our input.
	// A: Rock
	// B: Paper
	// Z: Scissors
	// X: LOSE
	// Y: DRAW
	// Z: WIN
	for buf.Scan() {
		game := strings.Split(buf.Text(), " ")
		if len(game) != 2 {
			panic("bad data")
		}

		switch game[0] {
		case "A":
			switch game[1] {
			case "X":
				score += SCISSORS + LOSE
			case "Y":
				score += ROCK + DRAW
			case "Z":
				score += PAPER + WIN
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
				score += PAPER + LOSE
			case "Y":
				score += SCISSORS + DRAW
			case "Z":
				score += ROCK + WIN
			}
		}
	}

	fmt.Println("Part 2:", score)
}
