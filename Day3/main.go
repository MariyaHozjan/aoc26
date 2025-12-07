package main

import (
	"AOC26/utils"
	"fmt"
	"log"
	"strconv"
)

func bestJoltage(bank string, k int) string {
	totalDigits := len(bank)
	result := ""
	searchStart := 0

	for chosen := 0; chosen < k; chosen++ {
		digitsLeftToPick := k - chosen
		maxSearchIndex := totalDigits - digitsLeftToPick
		bestDigit := bank[searchStart] // best number weve seen
		bestDigitIndex := searchStart  // idx of where it was
		for i := searchStart; i <= maxSearchIndex; i++ {
			if bank[i] > bestDigit {
				bestDigit = bank[i]
				bestDigitIndex = i
			}
		}
		result += string(bestDigit)
		searchStart = bestDigitIndex + 1
	}
	return result
}

func main() {
	banks, err := utils.ReadPuzzle("Day3/puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	totalJoltage := 0
	
	for _, bank := range banks {
		joltageStr := bestJoltage(bank, 12)
		val, err := strconv.Atoi(joltageStr)
		if err != nil {
			log.Fatal(err)
		}
		totalJoltage += val
	}
	fmt.Println(totalJoltage)
}
