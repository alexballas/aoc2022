package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	calories := make([]int, 0)
	var calcounter int

	buf := bufio.NewScanner(f)
	for buf.Scan() {
		if buf.Text() == "" {
			calories = append(calories, calcounter)
			calcounter = 0
			continue
		}

		num, err := strconv.Atoi(buf.Text())
		if err != nil {
			panic(err)
		}

		calcounter += num
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	if len(calories) > 2 {
		fmt.Println("Part 1:", calories[0])
		fmt.Println("Part 2:", calories[0]+calories[1]+calories[2])
	}
}
