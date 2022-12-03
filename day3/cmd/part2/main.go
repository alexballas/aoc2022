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

	bufSlice := make([]string, 3)
	var counter int

	for buf.Scan() {
		var common rune

		switch counter % 3 {
		case 0:
			bufSlice[0] = buf.Text()
		case 1:
			bufSlice[1] = buf.Text()
		case 2:
			bufSlice[2] = buf.Text()

		out:
			for _, a := range bufSlice[0] {
				for _, b := range bufSlice[1] {
					for _, c := range bufSlice[2] {
						if a == b && a == c {
							common = a
							break out
						}
					}
				}
			}

			priority += prio(common)
		}
		counter++
	}

	fmt.Println("Part 2:", priority)
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
