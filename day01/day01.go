package main

import (
	"slices"

	"github.com/julijane/advent-of-code-2024/aoc"
)

func part1(col1, col2 []int) int {
	sum := 0

	slices.Sort(col1)
	slices.Sort(col2)

	for i := 0; i < len(col1); i++ {
		sum += aoc.AbsInt(col1[i] - col2[i])
	}
	return sum
}

func part2(col1, col2 []int) int {
	sum := 0

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

func calc(input *aoc.Input, _, _ bool) (any, any) {
	col1 := []int{}
	col2 := []int{}

	for _, line := range input.Lines {
		numbers := aoc.ExtractNumbers(line.Data)

		col1 = append(col1, numbers[0])
		col2 = append(col2, numbers[1])
	}

	return part1(col1, col2), part2(col1, col2)
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
