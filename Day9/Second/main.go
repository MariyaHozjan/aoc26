package main

import (
	"AOC26/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func main() {
	puzzle, err := utils.ReadPuzzle("Day9/puzzle.txt")
	if err != nil {
		panic(err)
	}

	var Coordinates []Coordinate
	minX, minY := 999999, 999999
	maxX, maxY := -1, -1

	for _, l := range puzzle {
		s := strings.Split(l, ",")
		x, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
		Coordinates = append(Coordinates, Coordinate{x, y})
	}

	w, h := maxX-minX+1, maxY-minY+1
	grid := make([][]bool, h)
	for i := range grid {
		grid[i] = make([]bool, w)
	}

	for r := 0; r < h; r++ {
		var nodes []int
		for i := 0; i < len(Coordinates); i++ {
			a := Coordinates[i]
			b := Coordinates[(i+1)%len(Coordinates)]

			if a.x == b.x {
				y1, y2 := a.y-minY, b.y-minY
				if y1 > y2 {
					y1, y2 = y2, y1
				}
				if r >= y1 && r < y2 {
					nodes = append(nodes, a.x-minX)
				}
			}
		}
		sort.Ints(nodes)
		for i := 0; i < len(nodes); i += 2 {
			for x := nodes[i]; x < nodes[i+1]; x++ {
				grid[r][x] = true
			}
		}
	}

	for i := 0; i < len(Coordinates); i++ {
		a, b := Coordinates[i], Coordinates[(i+1)%len(Coordinates)]
		x1, x2 := a.x-minX, b.x-minX
		y1, y2 := a.y-minY, b.y-minY
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			for x := x1; x <= x2; x++ {
				grid[y][x] = true
			}
		}
	}

	ps := make([][]int, h+1)
	for i := range ps {
		ps[i] = make([]int, w+1)
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			v := 0
			if grid[y][x] {
				v = 1
			}
			ps[y+1][x+1] = v + ps[y][x+1] + ps[y+1][x] - ps[y][x]
		}
	}

	best := 0
	for i := 0; i < len(Coordinates); i++ {
		for j := i + 1; j < len(Coordinates); j++ {
			x1, x2 := Coordinates[i].x-minX, Coordinates[j].x-minX
			y1, y2 := Coordinates[i].y-minY, Coordinates[j].y-minY
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			if y1 > y2 {
				y1, y2 = y2, y1
			}

			area := (x2 - x1 + 1) * (y2 - y1 + 1)
			if area <= best {
				continue
			}

			count := ps[y2+1][x2+1] - ps[y1][x2+1] - ps[y2+1][x1] + ps[y1][x1]
			if count == area {
				best = area
			}
		}
	}

	fmt.Printf("Biggest rec possible: %d\n", best)
}
