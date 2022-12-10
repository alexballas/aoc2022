package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}
type Node struct {
	visited     map[Point]struct{}
	name        int
	x           int
	y           int
	uniqueMoves int
	isTail      bool
}

func main() {
	f, _ := os.Open("input.txt")

	buf := bufio.NewScanner(f)

	nodes := make([]*Node, 0, 10)

	for i := 0; i < 10; i++ {
		node := &Node{
			name:    i,
			visited: map[Point]struct{}{},
		}
		nodes = append(nodes, node)
	}

	head := nodes[0]
	tail := nodes[len(nodes)-1]

	tail.isTail = true
	tail.updateMoves()

	for buf.Scan() {
		lineSlice := strings.Split(buf.Text(), " ")
		direction := lineSlice[0]
		moves, err := strconv.Atoi(lineSlice[1])
		if err != nil {
			panic(err)
		}

		doMoves(moves, direction, head, nodes)

	}

	fmt.Println("Part 2:", tail.uniqueMoves)
	f.Close()
}

func isNear(h *Node, t *Node) bool {
	if (t.x == h.x && t.y == h.y) ||
		((t.x+1 == h.x || t.x-1 == h.x) && (t.y+1 == h.y || t.y-1 == h.y)) ||
		(t.x == h.x && (t.y+1 == h.y || t.y-1 == h.y)) ||
		(t.y == h.y && (t.x+1 == h.x || t.x-1 == h.x)) {
		return true
	}
	return false
}

func (p *Node) updateMoves() {
	_, exists := p.visited[Point{
		x: p.x,
		y: p.y,
	}]

	if !exists && p.isTail {

		p.visited[Point{
			x: p.x,
			y: p.y,
		}] = struct{}{}

		p.uniqueMoves++
	}
}

func doMoves(moves int, direction string, head *Node, nodes []*Node) {
	for j := 0; j < moves; j++ {
		oldPos := struct {
			node         Node
			newDiagonial Node
			diagonial    bool
		}{
			node:         *head,
			newDiagonial: *head,
			diagonial:    false,
		}

		switch direction {
		case "U":
			head.y++
		case "D":
			head.y--
		case "L":
			head.x--
		case "R":
			head.x++
		}

		for i := 0; i < 9; i++ {
			keep := struct {
				node         Node
				newDiagonial Node
				diagonial    bool
			}{
				node:         *nodes[i+1],
				newDiagonial: *nodes[i+1],
				diagonial:    false,
			}

			if !oldPos.diagonial {
				switch direction {
				case "U":
					if nodes[i].x == nodes[i+1].x && !isNear(nodes[i], nodes[i+1]) {
						nodes[i+1].y++
						nodes[i+1].updateMoves()
						oldPos = keep
						continue
					}
				case "D":
					if nodes[i].x == nodes[i+1].x && !isNear(nodes[i], nodes[i+1]) {
						nodes[i+1].y--
						nodes[i+1].updateMoves()
						oldPos = keep
						continue
					}

				case "L":
					if nodes[i].y == nodes[i+1].y && !isNear(nodes[i], nodes[i+1]) {
						nodes[i+1].x--
						nodes[i+1].updateMoves()
						oldPos = keep
						continue
					}

				case "R":
					if nodes[i].y == nodes[i+1].y && !isNear(nodes[i], nodes[i+1]) {
						nodes[i+1].x++
						nodes[i+1].updateMoves()
						oldPos = keep
						continue
					}
				}
			}

			if oldPos.diagonial && !isNear(nodes[i], nodes[i+1]) {
				var diagonial bool
				if nodes[i+1].x == nodes[i].x && nodes[i+1].y < nodes[i].y {
					nodes[i+1].y++
					nodes[i+1].updateMoves()
				} else if nodes[i+1].x == nodes[i].x && nodes[i+1].y > nodes[i].y {
					nodes[i+1].y--
					nodes[i+1].updateMoves()
				} else if nodes[i+1].y == nodes[i].y && nodes[i+1].x > nodes[i].x {
					nodes[i+1].x--
					nodes[i+1].updateMoves()
				} else if nodes[i+1].y == nodes[i].y && nodes[i+1].x < nodes[i].x {
					nodes[i+1].x++
					nodes[i+1].updateMoves()

				} else {
					if oldPos.newDiagonial.x > nodes[i+1].x {
						nodes[i+1].x++
					}

					if oldPos.newDiagonial.x < nodes[i+1].x {
						nodes[i+1].x--
					}

					if oldPos.newDiagonial.y > nodes[i+1].y {
						nodes[i+1].y++
					}

					if oldPos.newDiagonial.y < nodes[i+1].y {
						nodes[i+1].y--
					}

					diagonial = true
					nodes[i+1].updateMoves()

				}

				oldPos = keep
				if diagonial {
					oldPos.diagonial = true
					oldPos.newDiagonial = *nodes[i+1]
				}
				continue
			}

			if !isNear(nodes[i], nodes[i+1]) {
				nodes[i+1].x = oldPos.node.x
				nodes[i+1].y = oldPos.node.y
				nodes[i+1].updateMoves()
				oldPos = keep
				oldPos.diagonial = true
				oldPos.newDiagonial = *nodes[i+1]
			}
		}
	}

}
