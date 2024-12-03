package main

import (
	"regexp"

	"github.com/julijane/advent-of-code-2024/aoc"
)

func calcFor(code string, isPart2 bool) int {
	sum := 0

	reText := `mul\((\d+),(\d+)\)`
	if isPart2 {
		reText += `|don't\(\)|do\(\)`
	}

	re := regexp.MustCompile(reText)

	mulEnabled := true
	matches := re.FindAllStringSubmatch(code, -1)

	for _, match := range matches {
		if match[0] == "don't()" {
			mulEnabled = false
		} else if match[0] == "do()" {
			mulEnabled = true
		} else if mulEnabled {
			sum += aoc.Atoi(match[1]) * aoc.Atoi(match[2])
		}
	}

	return sum
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (int, int) {
	sumPart1 := 0
	sumPart2 := 0

	code := input.SingleString()

	if doPart1 {
		sumPart1 = calcFor(code, false)
	}

	if doPart2 {
		sumPart2 = calcFor(code, true)
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, false)
	aoc.Run("sample2.txt", calc, false, true)
	aoc.Run("input.txt", calc, true, true)
}
