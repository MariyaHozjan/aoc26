package main

import (
	"AOC26/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var puzzle, err = utils.ReadPuzzleComma("Day2/puzzle.txt")
	if err != nil {
		fmt.Println(err)
	}

	invalid := 0

	for _, val := range puzzle {
		start, end, err := SplitAndConvert(string(val))
		if err != nil {
			fmt.Println(err)
			continue
		}

		for num := start; num <= end; num++ {
			numStr := strconv.Itoa(num)

			if isInvalidRepeated(numStr) {
				invalid += num
			}
		}
	}

	fmt.Println(invalid)
}

func isInvalidRepeated(s string) bool {
	L := len(s)

	for patternLen := 1; patternLen <= L/2; patternLen++ {
		if L%patternLen != 0 {
			continue
		}

		numRepeats := L / patternLen
		if numRepeats < 2 {
			continue
		}

		repeat := s[:patternLen]
		candidate := strings.Repeat(repeat, numRepeats)

		if candidate == s {
			return true
		}
	}
	return false
}

func SplitAndConvert(input string) (int, int, error) {
	parts := strings.Split(strings.TrimSpace(input), "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid format: expected 2 parts, got %d", len(parts))
	}

	num1, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	num2, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return num1, num2, nil
}
