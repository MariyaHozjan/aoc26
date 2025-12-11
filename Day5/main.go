package main

import (
	"AOC26/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type rng struct {
	start int64
	end   int64
}

func splitRange(rangeInput string) (start int64, end int64) {
	idx := strings.Index(rangeInput, "-")
	if idx == -1 {
		log.Fatalf("invalid range: %q", rangeInput)
	}

	s := rangeInput[:idx]
	e := rangeInput[idx+1:]

	var err error
	start, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	end, err = strconv.ParseInt(e, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return start, end
}

func mergeRanges(ranges []rng) int64 {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	currentStart := ranges[0].start
	currentEnd := ranges[0].end

	var totalMerged int64 = 0

	for i := 1; i < len(ranges); i++ {
		rng := ranges[i]

		if rng.start <= currentEnd {
			if rng.end > currentEnd {
				currentEnd = rng.end
			}
		} else {
			totalMerged += currentEnd - currentStart + 1

			currentStart = rng.start
			currentEnd = rng.end
		}
	}
	totalMerged += currentEnd - currentStart + 1

	return totalMerged
}

func main() {
	lines, err := utils.ReadPuzzle("Day5/puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}

	splitIndex := -1
	for i, line := range lines {
		if line == "" {
			splitIndex = i
			break
		}
	}

	ingRanges := lines[:splitIndex]
	ingredientIds := lines[splitIndex+1:]

	var ranges []rng
	for _, ingRange := range ingRanges {
		start, end := splitRange(ingRange)
		ranges = append(ranges, rng{start, end})
	}

	freshCount := 0
	allFreshCounts := mergeRanges(ranges)
	fmt.Println(allFreshCounts)

	for _, ingredientId := range ingredientIds {

		id, err := strconv.ParseInt(ingredientId, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		isFresh := false
		for _, r := range ranges {
			if id >= r.start && id <= r.end {
				isFresh = true
				break
			}
		}

		if isFresh {
			freshCount++
		}
	}

	fmt.Println(freshCount)
}
