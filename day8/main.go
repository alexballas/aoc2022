package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	var (
		treemap     = make([][]rune, 0)
		treeVisible = make([][]bool, 0)
		treeScore   = make([][]int, 0)
		buf         = bufio.NewScanner(f)
		counter     int
	)

	for buf.Scan() {
		treemap = append(treemap, make([]rune, len(buf.Text())))
		treeVisible = append(treeVisible, make([]bool, len(buf.Text())))
		treeScore = append(treeScore, make([]int, len(buf.Text())))

		for n, c := range buf.Text() {
			treemap[counter][n] = c
		}

		counter++
	}

	for n, y := range treemap {
		for nn, x := range y {
			leftvisible := true
			rightvisible := true

			if n == 0 || n == len(treemap)-1 {
				treeVisible[n][nn] = true
				continue
			}

			if nn == 0 || nn == len(y)-1 {
				treeVisible[n][nn] = true
				continue
			}

			for nnn, xx := range y {
				if nn > nnn && x <= xx {
					leftvisible = false
				}

				if nn < nnn && x <= xx {
					rightvisible = false
				}
			}

			if leftvisible || rightvisible {
				treeVisible[n][nn] = true
			}
		}
	}

	for n, y := range treemap {
		for nn, x := range y {
			topvisible := true
			bottomvisible := true
			for n1, y1 := range treemap {
				for nnn, xx := range y1 {
					if nn == nnn {
						if n1 < n && x <= xx {
							topvisible = false
						}
						if n1 > n && x <= xx {
							bottomvisible = false
						}
					}
				}
			}
			if topvisible || bottomvisible {
				treeVisible[n][nn] = true
			}
		}
	}

	var final int
	for _, treeRowVisible := range treeVisible {
		for _, treeVisible := range treeRowVisible {
			if treeVisible {
				final++
			}
		}
	}

	fmt.Println("Part 1:", final)

	for _, scoreRow := range treeScore {
		for n := range scoreRow {
			scoreRow[n] = 1
		}
	}

	for n, y := range treemap {
		for nn, x := range y {
			var (
				minDistanceLeft  = nn
				minDistanceRight = len(y) - nn - 1
				blockRight       bool
			)

			for nnn, xx := range y {
				if nn > nnn && x <= xx {
					minDistanceLeft = nn - nnn
				}

				if nn < nnn && x <= xx && !blockRight {
					minDistanceRight = nnn - nn
					blockRight = true
				}
			}

			if minDistanceLeft == 0 {
				minDistanceLeft = 1
			}
			if minDistanceRight == 0 {
				minDistanceRight = 1
			}

			treeScore[n][nn] *= minDistanceLeft * minDistanceRight
		}
	}

	for n, y := range treemap {
		for nn, x := range y {
			var (
				topMinDistance    = n
				bottomMinDistance = len(treemap) - n - 1
				blockBottom       bool
			)

			for n1, y1 := range treemap {
				for nnn, xx := range y1 {
					if nn == nnn {
						if n1 < n && x <= xx {
							topMinDistance = n - n1
						}

						if n1 > n && x <= xx && !blockBottom {
							bottomMinDistance = n1 - n
							blockBottom = true
						}

						if topMinDistance == 0 {
							topMinDistance = 1
						}
						if bottomMinDistance == 0 {
							bottomMinDistance = 1
						}
					}
				}
			}

			treeScore[n][nn] *= topMinDistance * bottomMinDistance
		}
	}

	var max int
	for _, scoreRow := range treeScore {
		for _, n := range scoreRow {
			if n > max {
				max = n
			}
		}
	}

	fmt.Println("Part 2:", max)
}
