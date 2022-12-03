package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	buf := bufio.NewScanner(f)

	var priority int32

	for buf.Scan() {
		line := buf.Text()
		word1 := line[:len(line)/2]
		word2 := line[len(line)/2:]

		var common rune
	out:
		for _, a := range word1 {
			for _, b := range word2 {
				if a == b {
					common = a
					break out
				}
			}
		}

		priority += prio(common)
	}

	fmt.Println("Part 1:", priority)
}

func prio(a rune) rune {
	if a >= 'A' && a <= 'Z' {
		a = a - 'A' + 27
		return a
	}

	if a >= 'a' && a <= 'z' {
		a = a - 'a' + 1
		return a
	}

	return -1
}
