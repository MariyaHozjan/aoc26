package utils

import (
	"bufio"
	"os"
)

func ReadPuzzle(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var puzzle []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		puzzle = append(puzzle, scanner.Text())
	}
	return puzzle, scanner.Err()
}
