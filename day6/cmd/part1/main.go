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

	var result int

	for buf.Scan() {
		line := buf.Text()

		for i := 0; i < len(line)-3; i++ {
			fourChars := line[i : i+4]
			uniq := make(map[rune]struct{})
			for _, letter := range fourChars {
				uniq[letter] = struct{}{}
			}

			if len(uniq) == 4 {
				result = i + 4
				fmt.Println("Part 1:", result)
				break
			}
		}
	}
}
