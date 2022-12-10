package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	visited     map[struct{ x, y int }]struct{}
	x           int
	y           int
	uniqueMoves int
}

func main() {
	f, _ := os.Open("input.txt")
	//	f, _ := os.Open("input_test.txt")

	defer f.Close()

	buf := bufio.NewScanner(f)

	head := &Pos{
		visited: map[struct {
			x int
			y int
		}]struct{}{},
	}
	tail := &Pos{
		visited: map[struct {
			x int
			y int
		}]struct{}{},
	}

	for buf.Scan() {

		lineSlice := strings.Split(buf.Text(), " ")
		direction := lineSlice[0]
		moves, err := strconv.Atoi(lineSlice[1])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "U":
			for i := 0; i < moves; i++ {
				oldHead := *head
				head.y++

				// Moved to a new position on the same axis X as Tail
				if head.x == tail.x && !isNear(head, tail) {
					tail.y++
					tail.updateMoves()
					continue
				}

				// Head was already diagonally to Tail. So it now moved
				// even further. Taul should occupy Heads old position.
				if !isNear(head, tail) {
					tail.x = oldHead.x
					tail.y = oldHead.y
					tail.updateMoves()
					continue
				}
			}
		case "D":
			for i := 0; i < moves; i++ {
				oldHead := *head
				head.y--
				if head.x == tail.x && !isNear(head, tail) {
					tail.y--
					tail.updateMoves()
					continue
				}

				if !isNear(head, tail) {
					tail.x = oldHead.x
					tail.y = oldHead.y
					tail.updateMoves()
					continue
				}
			}
		case "L":
			for i := 0; i < moves; i++ {
				oldHead := *head
				head.x--
				if head.y == tail.y && !isNear(head, tail) {
					tail.x--
					tail.updateMoves()
					continue
				}

				if !isNear(head, tail) {
					tail.x = oldHead.x
					tail.y = oldHead.y
					tail.updateMoves()
					continue
				}

			}
		case "R":
			for i := 0; i < moves; i++ {
				oldHead := *head
				head.x++
				if head.y == tail.y && !isNear(head, tail) {
					tail.x++
					tail.updateMoves()
					continue
				}

				if !isNear(head, tail) {
					tail.x = oldHead.x
					tail.y = oldHead.y
					tail.updateMoves()
					continue
				}
			}
		}
	}

	fmt.Println("Part 1:", tail.uniqueMoves)
}

func isNear(h *Pos, t *Pos) bool {
	if (t.x == h.x && t.y == h.y) ||
		((t.x+1 == h.x || t.x-1 == h.x) && (t.y+1 == h.y || t.y-1 == h.y)) ||
		(t.x == h.x && (t.y+1 == h.y || t.y-1 == h.y)) ||
		(t.y == h.y && (t.x+1 == h.x || t.x-1 == h.x)) {
		return true
	}
	return false
}

func (p *Pos) updateMoves() {
	_, exists := p.visited[struct {
		x int
		y int
	}{
		x: p.x,
		y: p.y,
	}]

	if !exists {
		p.visited[struct {
			x int
			y int
		}{
			x: p.x,
			y: p.y,
		}] = struct{}{}

		p.uniqueMoves++
	}
}
