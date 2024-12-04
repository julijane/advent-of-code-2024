package main

import (
	"github.com/julijane/advent-of-code-2024/aoc"
)

var directions = []aoc.Coordinate{
	{X: -1, Y: 0},
	{X: -1, Y: -1},
	{X: 0, Y: -1},
	{X: 1, Y: -1},
	{X: 1, Y: 0},
	{X: 1, Y: 1},
	{X: 0, Y: 1},
	{X: -1, Y: 1},
}

func checkForStartPosPart1(grid *aoc.Grid, pos aoc.Coordinate) int {
	sum := 0

	for _, direction := range directions {
		if grid.StringFrom(pos, direction, 4, '.') == "XMAS" {
			sum++
		}
	}

	return sum
}

func checkForStartPosPart2(grid *aoc.Grid, pos aoc.Coordinate) int {
	stringAL2BR := grid.StringFrom(pos.UpLeft(), aoc.DirDR, 3, '.')
	stringAR2BL := grid.StringFrom(pos.UpRight(), aoc.DirDL, 3, '.')

	if (stringAL2BR == "MAS" || stringAL2BR == "SAM") &&
		(stringAR2BL == "MAS" || stringAR2BL == "SAM") {
		return 1
	}

	return 0
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (int, int) {
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
