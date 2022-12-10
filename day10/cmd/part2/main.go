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
		cycle    int
		register int = 1
	)

	buffer := make([]string, 40)
	clearBuffer(buffer, true)

	crtScanning := make([]string, 40)
	clearBuffer(crtScanning, false)

	for buf.Scan() {
		lineSlice := strings.Split(buf.Text(), " ")
		switch len(lineSlice) {
		case 1:
			crtScanning[cycle%40] = buffer[cycle%40]
			cycle++
			if cycle%40 == 0 {
				fmt.Println(crtScanning)
			}
		case 2:
			for i := 0; i < 2; i++ {
				crtScanning[cycle%40] = buffer[cycle%40]
				cycle++
				if cycle%40 == 0 {
					fmt.Println(crtScanning)
				}
			}
			num, _ := strconv.Atoi(lineSlice[1])
			register += num
			clearBuffer(buffer, false)
			drawBuffer(buffer, register)
		}
	}

}

func drawBuffer(b []string, r int) {
	if r >= 0 && r <= len(b)-1 {
		b[r] = "#"
	}

	if r+1 >= 0 && r+1 <= len(b)-1 {
		b[r+1] = "#"
	}

	if r-1 >= 0 && r-1 <= len(b)-1 {
		b[r-1] = "#"
	}
}

func clearBuffer(b []string, addChar bool) {
	for n := range b {
		switch n {
		case 0, 1, 2:
			b[n] = "."
			if addChar {
				b[n] = "#"
			}
		default:
			b[n] = "."
		}
	}
}
