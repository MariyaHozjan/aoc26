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

type IndexPairs struct {
	first, second int
}

type RecSizes struct {
	pair IndexPairs
	size int
}

func main() {
	puzzle, err := utils.ReadPuzzle("Day9/puzzle.txt")
	if err != nil {
		panic(err)
	}

	var Coordinates []Coordinate
	sizes := make([]RecSizes, 0)

	for _, line := range puzzle {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		Coordinates = append(Coordinates, Coordinate{x, y})
	}

	for i := 0; i < len(Coordinates); i++ {
		for j := i + 1; j < len(Coordinates); j++ {
			pair := IndexPairs{i, j}
			dx := (Coordinates[i].x - Coordinates[j].x) + 1
			dy := (Coordinates[i].y - Coordinates[j].y) + 1
			rec := utils.Abs(dx * dy)
			sizes = append(sizes, RecSizes{pair, rec})
		}
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i].size > sizes[j].size
	})

	fmt.Println("Biggest rectangle possible: ", sizes[0].size)
}
