package main

import (
	"slices"

	"github.com/julijane/advent-of-code-2024/aoc"
)

func getCols(input *aoc.Input) ([]int, []int) {
	col1 := []int{}
	col2 := []int{}

	for _, line := range input.Lines {
		numbers := aoc.ExtractNumbers(line.Data)

		col1 = append(col1, numbers[0])
		col2 = append(col2, numbers[1])
	}

	return col1, col2
}

func part1(input *aoc.Input) int {
	sum := 0
	col1, col2 := getCols(input)

	slices.Sort(col1)
	slices.Sort(col2)

	for i := 0; i < len(col1); i++ {
		if col1[i] > col2[i] {
			sum += col1[i] - col2[i]
		} else {
			sum += col2[i] - col1[i]
		}
	}
	return sum
}

func part2(input *aoc.Input) int {
	sum := 0
	col1, col2 := getCols(input)

	for _, col1num := range col1 {
		numfound := 0
		for _, col2num := range col2 {
			if col2num == col1num {
				numfound++
			}
		}
		sum += col1num * numfound
	}

	return sum
}

func calc(input *aoc.Input, runPart1, runPart2 bool) (int, int) {
	sumPart1 := 0
	sumPart2 := 0

	if runPart1 {
		sumPart1 = part1(input)
	}

	if runPart2 {
		sumPart2 = part2(input)
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
