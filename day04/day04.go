package main

import (
	"github.com/julijane/advent-of-code-2024/aoc"
)

func checkForStartPosPart1(grid *aoc.Grid, pos aoc.Coordinate) int {
	sum := 0

	for _, direction := range aoc.DirsAll {
		if grid.StringFrom(pos, direction, 4, '.') == "XMAS" {
			sum++
		}
	}

	return sum
}

func checkForStartPosPart2(grid *aoc.Grid, pos aoc.Coordinate) int {
	stringUL2DR := grid.StringFrom(pos.UpLeft(), aoc.DirDR, 3, '.')
	stringUR2DL := grid.StringFrom(pos.UpRight(), aoc.DirDL, 3, '.')

	if (stringUL2DR == "MAS" || stringUL2DR == "SAM") &&
		(stringUR2DL == "MAS" || stringUR2DL == "SAM") {
		return 1
	}

	return 0
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (any, any) {
	sumPart1 := 0
	sumPart2 := 0

	grid := input.Grid()
	for sX := 0; sX < grid.Width; sX++ {
		for sY := 0; sY < grid.Height; sY++ {
			pos := aoc.Coordinate{X: sX, Y: sY}
			sumPart1 += checkForStartPosPart1(grid, pos)
			sumPart2 += checkForStartPosPart2(grid, pos)
		}
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
