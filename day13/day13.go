package main

import (
	"github.com/julijane/advent-of-code-2024/aoc"
)

func twoNumbers(slice []int) (int, int) {
	return slice[0], slice[1]
}

func solveFor(aX, aY, bX, bY, pX, pY int) int {
	m := aX*bY - aY*bX
	n := aX*pY - aY*pX

	if n%m != 0 {
		return 0
	}

	b := n / m

	r := pX - bX*b
	if r%aX != 0 {
		return 0
	}

	a := r / aX

	return a*3 + b
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (any, any) {
	sumPart1 := 0
	sumPart2 := 0

	offsetPart2 := 10000000000000

	lines := input.PlainLines()
	for line := 0; line < len(lines); line += 4 {
		aX, aY := twoNumbers(aoc.ExtractNumbers(lines[line]))
		bX, bY := twoNumbers(aoc.ExtractNumbers(lines[line+1]))
		pX, pY := twoNumbers(aoc.ExtractNumbers(lines[line+2]))

		sumPart1 += solveFor(aX, aY, bX, bY, pX, pY)
		sumPart2 += solveFor(aX, aY, bX, bY, pX+offsetPart2, pY+offsetPart2)
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
