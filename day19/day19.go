package main

import (
	"strings"

	"github.com/julijane/advent-of-code-2024/aoc"
)

var possibleCache map[string]int

func howManyPossible(target string, availablePattern []string) int {
	if numPossible, ok := possibleCache[target]; ok {
		return numPossible
	}

	numPossible := 0
	for _, pattern := range availablePattern {
		if strings.HasPrefix(target, pattern) {
			shorterTarget := target[len(pattern):]
			if shorterTarget == "" {
				numPossible++
				continue
			}

			numPossible += howManyPossible(shorterTarget, availablePattern)
		}
	}

	possibleCache[target] = numPossible
	return numPossible
}

func calc(input *aoc.Input, _, _ bool, param ...any) (any, any) {
	resultPart1 := 0
	resultPart2 := 0

	availablePattern := strings.Split(input.Lines[0].Data, ", ")
	possibleCache = make(map[string]int)

	for _, line := range input.Lines[2:] {
		numPossible := howManyPossible(line.Data, availablePattern)
		if numPossible > 0 {
			resultPart1++
		}

		resultPart2 += numPossible
	}

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
