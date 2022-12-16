package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"os/exec"
	"time"
)

type Node struct {
	Children []*Node
	Point    image.Point
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	matrix := make([][]rune, 0)

	buf := bufio.NewScanner(f)

	var part1, end image.Point
	starts := make([]image.Point, 0)
	var counter int

	for buf.Scan() {
		lineSlice := make([]rune, 0)
		var counter2 int
		for _, r := range buf.Text() {
			if r == 'S' {
				r = 'a'
				part1 = image.Point{counter2, counter}
			}

			if r == 'a' {
				starts = append(starts, image.Point{counter2, counter})
			}

			if r == 'E' {
				r = 'z'
				end = image.Point{counter2, counter}
			}

			lineSlice = append(lineSlice, r)
			counter2++
		}
		counter++

		matrix = append(matrix, lineSlice)
	}

	var drawpath []image.Point
	var minPath []image.Point

	var part1res, min int

	for _, start := range starts {
		res, path := BreadthFirstSearch(matrix, start, end)
		if min == 0 && path != nil {
			min = res
			minPath = path
		}

		if res <= min && path != nil {
			min = res
			minPath = path
		}

		if start == part1 {
			part1res = res
			drawpath = path
		}
	}

	// Draw it just for fun
	reverse(minPath)
	reverse(drawpath)

	for _, point := range drawpath {
		for y, l := range matrix {
			for x := range l {
				curr := image.Point{X: x, Y: y}
				if point == curr {
					matrix[y][x] = ' '
				}

			}
		}

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()

		for _, line := range matrix {
			fmt.Println(string(line))
		}

		time.Sleep(time.Millisecond * 10)
	}

	for _, point := range minPath {

		for y, l := range matrix {
			for x := range l {
				curr := image.Point{X: x, Y: y}
				if point == curr {
					matrix[y][x] = '.'
				}

			}
		}

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()

		for _, line := range matrix {
			fmt.Println(string(line))
		}

		time.Sleep(time.Millisecond * 10)
	}

	fmt.Println("Part 1:", part1res)
	fmt.Println("Part 2:", min)

}

func BreadthFirstSearch(matrix [][]rune, start, end image.Point) (int, []image.Point) {
	queue := []image.Point{start}
	parents := make(map[image.Point]image.Point)
	visited := make(map[image.Point]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		visited[current] = true

		if current == end {
			path := []image.Point{end}
			for curr := parents[end]; curr != start; curr = parents[curr] {
				path = append(path, curr)
			}

			return len(path), path
		}

		for _, n := range getNeighbors(matrix, current) {
			_, exists := parents[n]
			if !exists {
				parents[n] = current
			}

			queue = append(queue, n)
		}
	}
	return -1, nil
}

func getNeighbors(matrix [][]rune, p image.Point) []image.Point {
	neighbors := make([]image.Point, 0)

	pointsNear := []image.Point{
		{
			X: 1,
			Y: 0,
		},
		{
			X: -1,
			Y: 0,
		},
		{
			X: 0,
			Y: 1,
		},
		{
			X: 0,
			Y: -1,
		},
	}
	for _, pn := range pointsNear {
		newPoint := p.Add(pn)
		rect := image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{len(matrix[0]), len(matrix)},
		}

		if newPoint.In(rect) &&
			(matrix[p.Y][p.X]+1 == matrix[newPoint.Y][newPoint.X] || matrix[p.Y][p.X] >= matrix[newPoint.Y][newPoint.X]) {
			neighbors = append(neighbors, newPoint)
		}
	}

	return neighbors
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
