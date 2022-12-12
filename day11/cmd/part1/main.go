package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Item struct {
	Worry int
}

type Monkey struct {
	Items     []Item
	Operation func(int) int
	Test      func(int) bool
	True      int
	False     int
	Inspected int
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	buf := bufio.NewScanner(f)

	bufferMonkey := new(Monkey)

	monkeyList := make([]*Monkey, 0)

	for buf.Scan() {
		lineSlice := strings.Split(strings.TrimSpace(buf.Text()), " ")
		if len(lineSlice) > 1 {
			switch lineSlice[0] {
			case "Monkey":
				monke := new(Monkey)
				bufferMonkey = monke
				monkeyList = append(monkeyList, monke)
			case "Starting":
				bufferMonkey.Items = make([]Item, 0, 1002)
				for _, i := range lineSlice[2:] {
					p := strings.Trim(i, ",")
					itemWorry, _ := strconv.Atoi(p)
					item := Item{
						Worry: itemWorry,
					}
					bufferMonkey.Items = append(bufferMonkey.Items, item)
				}
			case "Operation:":
				newOp := func(old int) int {
					switch lineSlice[4] {
					case "*":
						if lineSlice[5] == "old" {
							return (old * old) / 3
						} else {
							v, _ := strconv.Atoi(lineSlice[5])
							return (old * v) / 3
						}
					case "+":
						if lineSlice[5] == "old" {
							return (old + old) / 3
						} else {
							v, _ := strconv.Atoi(lineSlice[5])
							return (old + v) / 3
						}
					default:
						panic("invalid")
					}
				}

				bufferMonkey.Operation = newOp
			case "Test:":
				test := func(worry int) bool {
					v, _ := strconv.Atoi(lineSlice[3])
					return worry%v == 0
				}
				bufferMonkey.Test = test
			case "If":
				switch lineSlice[1] {
				case "true:":
					v, _ := strconv.Atoi(lineSlice[5])
					bufferMonkey.True = v
				case "false:":
					v, _ := strconv.Atoi(lineSlice[5])
					bufferMonkey.False = v
				}
			}
		}

	}

	for i := 0; i < 20; i++ {
		for _, monke := range monkeyList {
			for _, item := range monke.Items {
				monke.Inspected++
				wAfterOp := monke.Operation(item.Worry)
				item.Worry = wAfterOp

				if monke.Test(wAfterOp) {
					toMonkey := monkeyList[monke.True]
					toMonkey.Items = append(toMonkey.Items, item)

				} else {
					toMonkey := monkeyList[monke.False]
					toMonkey.Items = append(toMonkey.Items, item)
				}
				monke.Items = make([]Item, 0)
			}
		}
	}

	result := make([]int, 0)
	for _, monke := range monkeyList {
		result = append(result, monke.Inspected)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(result)))

	fmt.Println("Part 1:", result[0]*result[1])
}
