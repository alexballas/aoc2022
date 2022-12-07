package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	parent     *node
	name       string
	childNodes []*node
	fileSize   int
}

const (
	Total    = 70000000
	Required = 30000000
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	buf := bufio.NewScanner(f)

	root := new(node)
	currentNode := new(node)

	for buf.Scan() {
		line := buf.Text()
		input := strings.Split(line, " ")
		switch input[0] {
		case "$":
			switch input[1] {
			case "cd":
				switch input[2] {
				case "/":
					root.name = "/"
					root.parent = nil
					currentNode = root
				case "..":
					currentNode = currentNode.parent
				default:
					newnode := &node{
						name:   input[2],
						parent: currentNode,
					}
					currentNode.childNodes = append(currentNode.childNodes, newnode)
					currentNode = newnode
				}
			}
		case "dir":
		default:
			curSize, err := strconv.Atoi(input[0])
			if err != nil {
				panic(err)
			}

			currentNode.fileSize += curSize
		}
	}

	nodeListSizes := make(map[*node]int)
	root.GetSetSize(nodeListSizes, 100000)

	var out int
	for _, size := range nodeListSizes {
		out += size
	}
	fmt.Println("Part 1:", out)

	totalUsed := root.GetSetSize(nodeListSizes, 0)
	nodesSpaceSlice := make([]int, 0, len(nodeListSizes))

	for _, size := range nodeListSizes {
		nodesSpaceSlice = append(nodesSpaceSlice, size)
	}

	sort.Ints(nodesSpaceSlice)

	currentFree := Total - totalUsed
	for _, i := range nodesSpaceSlice {
		if currentFree+i >= Required {
			fmt.Println("Part 2:", i)
			break
		}
	}
}

func (n *node) GetSetSize(list map[*node]int, limit int) int {
	var size int
	for _, i := range n.childNodes {
		nodeSize := i.GetSetSize(list, limit)
		size += nodeSize
	}

	size += n.fileSize

	if size <= limit || limit == 0 {
		list[n] = size
	}

	return size
}
