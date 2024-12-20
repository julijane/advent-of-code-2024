package main

import (
	"github.com/julijane/advent-of-code-2024/aoc"
)

func isOffending(numbers []int) bool {
	isIncreasing := numbers[1] > numbers[0]

	for i := 1; i < len(numbers); i++ {
		if isIncreasing != (numbers[i] > numbers[i-1]) {
			return true
		}

		diff := aoc.AbsInt(numbers[i] - numbers[i-1])
		if diff < 1 || diff > 3 {
			return true
		}
	}

	return false
}

func check(input *aoc.Input, isPart2 bool) int {
	sum := 0

lineLoop:
	for _, line := range input.Lines {
		numbers := aoc.ExtractNumbers(line.Data)
		if !isOffending(numbers) {
			sum++
			continue lineLoop
		}

		if !isPart2 {
			continue lineLoop
		}

		for skip := 0; skip < len(numbers); skip++ {
			newSlice := make([]int, skip)
			copy(newSlice, numbers[:skip])
			newSlice = append(newSlice, numbers[skip+1:]...)

			if !isOffending(newSlice) {
				sum++
				continue lineLoop
			}
		}
	}

	return sum
}

func calc(input *aoc.Input, _, _ bool, _ ...any) (any, any) {
	return check(input, false), check(input, true)
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
