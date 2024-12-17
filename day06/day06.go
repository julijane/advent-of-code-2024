package main

import (
	"github.com/julijane/advent-of-code-2024/aoc"
)

func doWalkStep(grid *aoc.Grid, pointer *aoc.Pointer) bool {
	for {
		nextContent := grid.Get(pointer.PeekMove(), '~')

		if nextContent == '~' {
			return false
		}

		if nextContent != '#' {
			pointer.Move()
			return true
		}

		pointer.TurnRight()
	}
}

func doWalkPart1(grid *aoc.Grid, pos aoc.Coordinate) []aoc.Coordinate {
	visited := make(map[aoc.Coordinate]struct{})
	walker := aoc.Pointer{C: pos, Dir: 0}

	for {
		visited[walker.C] = struct{}{}
		if !doWalkStep(grid, &walker) {
			break
		}
	}

	// Then convert the map to a slice
	visitedSlice := make([]aoc.Coordinate, len(visited))

	i := 0
	for coordinate := range visited {
		visitedSlice[i] = coordinate
		i++
	}

	return visitedSlice
}

func doWalkPart2(grid *aoc.Grid, additionalObstacle, startPos aoc.Coordinate) bool {
	oldContent := grid.Get(additionalObstacle, '~')
	grid.Set(additionalObstacle, '#')
	defer grid.Set(additionalObstacle, oldContent)

	slowWalker := aoc.Pointer{C: startPos, Dir: 0}
	fastWalker := aoc.Pointer{C: startPos, Dir: 0}

	for {
		doWalkStep(grid, &slowWalker)

		doWalkStep(grid, &fastWalker)
		if !doWalkStep(grid, &fastWalker) {
			return false
		}

		if slowWalker == fastWalker {
			return true
		}
	}
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (any, any) {
	sumPart2 := 0

	grid := input.Grid()

	startPos := grid.Find('^')

	// Part1
	visited := doWalkPart1(grid, startPos)
	sumPart1 := len(visited)

	// Part2
	// We can only set one obstacle at a time, so we only need to test the positions
	// visited in Part1 as potential obstacle positions
	for _, additionalObstaclePos := range visited {
		if doWalkPart2(grid, additionalObstaclePos, startPos) {
			sumPart2++
		}
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
