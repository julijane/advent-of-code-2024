package main

import (
	"regexp"
	"strings"

	"github.com/julijane/advent-of-code-2024/aoc"
)

func calcFor(line string) int {
	sum := 0

	mulInstructs := aoc.ExtractRegexps(line, `mul\(\d+\,\d+\)`)
	for _, mulInstruct := range mulInstructs {
		numbers := aoc.ExtractNumbers(mulInstruct)
		sum += numbers[0] * numbers[1]
	}

	return sum
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (int, int) {
	singleLine := ""

	sumPart1 := 0
	sumPart2 := 0

	for _, line := range input.Lines {
		singleLine += line.Data
	}

	if doPart1 {
		sumPart1 = calcFor(singleLine)
	}

	if doPart2 {
		lastDont := strings.LastIndex(singleLine, "don't()")
		lastDo := strings.LastIndex(singleLine, "do()")

		if lastDont > lastDo {
			singleLine = singleLine[:lastDont]
		}

		re := regexp.MustCompile(`don't\(\).*?do\(\)`)
		strippedLine := re.ReplaceAllString(singleLine, "")

		sumPart2 = calcFor(strippedLine)
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, false)
	aoc.Run("sample2.txt", calc, false, true)
	aoc.Run("input.txt", calc, true, true)
}
