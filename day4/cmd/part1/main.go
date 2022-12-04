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

	var counter int
	buf := bufio.NewScanner(f)
	for buf.Scan() {
		pairs := strings.Split(buf.Text(), ",")
		if len(pairs) != 2 {
			panic("bad data")
		}

		pair1 := strings.Split(pairs[0], "-")
		n1s, n2s := pair1[0], pair1[1]

		pair2 := strings.Split(pairs[1], "-")
		n3s, n4s := pair2[0], pair2[1]

		n1, _ := strconv.Atoi(n1s)
		n2, _ := strconv.Atoi(n2s)
		n3, _ := strconv.Atoi(n3s)
		n4, _ := strconv.Atoi(n4s)

		if n1 >= n3 && n2 <= n4 || n3 >= n1 && n4 <= n2 {
			counter++
		}

	}

	fmt.Println("Part 1:", counter)
}
