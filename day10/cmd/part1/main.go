package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	buf := bufio.NewScanner(f)

	var (
		signal   int
		cycle    int
		register int = 1
	)

	for buf.Scan() {
		lineSlice := strings.Split(buf.Text(), " ")
		switch len(lineSlice) {
		case 1:
			signal += checkSignal(1, &cycle, register)
		case 2:
			signal += checkSignal(2, &cycle, register)

			num, _ := strconv.Atoi(lineSlice[1])
			register += num
		}
	}

	fmt.Println("Part 1:", signal)
}

func checkSignal(times int, cycle *int, register int) int {
	var signal int
	for i := 0; i < times; i++ {
		signal += func() int {
			*cycle++
			switch *cycle {
			case 20, 60, 100, 140, 180, 220:
				return *cycle * register
			default:
				return 0
			}
		}()
	}
	return signal
}
