package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	stacks := make([]ppSlice, 0)
	buf := bufio.NewScanner(f)
	var stackInit bool

	for buf.Scan() {
		line := buf.Text()
		if len(line) > 0 && strings.TrimSpace(line)[0] == byte('[') {
			ss := transfLine(line)

			if !stackInit {
				for range ss {
					stacks = append(stacks, ppSlice{})
				}
				stackInit = true
			}

			var stack int
			for _, crate := range ss {
				stacks[stack] = append(stacks[stack], crate)
				stack++
			}
		}
	}

	for n, stack := range stacks {
		stacks[n] = cleanupSlice(stack)
		rev(stacks[n])
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		panic(err)
	}

	buf = bufio.NewScanner(f)

	for buf.Scan() {
		line := buf.Text()
		splitLine := strings.Split(line, " ")

		if len(line) > 0 && splitLine[0] == "move" {
			howMany, _ := strconv.Atoi(splitLine[1])
			fromStack, _ := strconv.Atoi(splitLine[3])
			toStack, _ := strconv.Atoi(splitLine[5])
			fromStack--
			toStack--

			for i := 0; i < howMany; i++ {
				outSlice, popped := stacks[fromStack].Pop(1)
				if popped == nil {
					panic("popped is nil")
				}
				stacks[fromStack] = outSlice
				stacks[toStack] = append(stacks[toStack], popped...)
			}
		}
	}

	final := new(strings.Builder)

	for i := range stacks {
		_, popped := stacks[i].Pop(1)
		if popped == nil {
			panic("popped is nil")
		}

		final.WriteString(popped[0])
	}

	fmt.Println("Part 1:", final.String())
}

func cleanupSlice(pp ppSlice) ppSlice {
	out := make(ppSlice, 0)
	for _, i := range pp {
		if i != "" {
			out = append(out, i)
		}
	}
	return out
}

func rev(a ppSlice) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func transfLine(l string) []string {
	ss := make([]string, 0, 9)
	var numOfFoundSpaces int
	numOfSpaces := 4

	for _, i := range l {
		if i != 32 && i != 91 && i != 93 {
			spacesToappend := numOfFoundSpaces / numOfSpaces
			for i := 0; i < spacesToappend; i++ {
				ss = append(ss, "")
			}

			ss = append(ss, string(i))
			numOfFoundSpaces = 0
		}

		if i == 32 {
			numOfFoundSpaces++
		}
	}

	if numOfFoundSpaces == numOfSpaces {
		ss = append(ss, "")
	}

	return ss
}

type ppSlice []string

func (s ppSlice) Push(v ...string) ppSlice {
	return append(s, v...)
}

func (s ppSlice) Pop(n int) (ppSlice, ppSlice) {
	l := len(s)
	if n > l {
		return s, nil
	}

	return s[:l-n], s[l-n:]
}
