package main

import (
	"bufio"
	"fmt"
	"log"
	"sort"

	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
	Z int
}

type Circuit struct {
	I        int
	J        int
	Distance int64
}

func EuclideanDistance(p1 Point, p2 Point) int64 {
	dx := int64(p1.X - p2.X)
	dy := int64(p1.Y - p2.Y)
	dz := int64(p1.Z - p2.Z)
	return dx*dx + dy*dy + dz*dz
}

func ReadCoordinate(path string) ([]Point, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var points []Point
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")

		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		z, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, err
		}

		points = append(points, Point{x, y, z})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return points, nil
}

func main() {
	var coordinates, err = ReadCoordinate("Day8/puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}

	n := len(coordinates)

	distances := make([]Circuit, 0, (n*n-1)/2)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := EuclideanDistance(coordinates[i], coordinates[j])
			distances = append(distances, Circuit{
				I:        i,
				J:        j,
				Distance: d,
			})
			//fmt.Printf("%d,%d -> %d\n", coordinates[i], coordinates[j], d)
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	circuitID := make([]int, n)
	members := make(map[int][]int, n)

	for i := 0; i < n; i++ {
		circuitID[i] = i
		members[i] = []int{i}
	}

	nc := len(distances)
	if nc > len(distances) {
		nc = len(distances)
	}

	lastA, lastB := -1, -1

	for w := 0; w < nc; w++ {
		a, b := distances[w].I, distances[w].J
		ca, cb := circuitID[a], circuitID[b]
		if ca == cb {
			continue
		}

		lastA, lastB = a, b

		if len(members[ca]) < len(members[cb]) {
			ca, cb = cb, ca
		}

		for _, v := range members[cb] {
			circuitID[v] = ca
			members[ca] = append(members[ca], v)
		}
		delete(members, cb)
	}

	/*
		PART ONE SOLUTION
			sizes := make([]int, 0, len(members))
			for _, ms := range members {
				sizes = append(sizes, len(ms))
			}
			sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })

			fmt.Println("Top:", sizes[0], sizes[1], sizes[2])
			fmt.Println("Product:", sizes[0]*sizes[1]*sizes[2])
	*/

	x1, x2 := coordinates[lastA].X, coordinates[lastB].X
	fmt.Println("Product of last connections", x1*x2)
}
