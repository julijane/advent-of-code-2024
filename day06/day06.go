package main

import (
	"github.com/julijane/advent-of-code-2024/aoc"
)

// doWalk walks the grid, returns if it ended up in a loop.
// If not, it also returns the visited positions (without the startpos)
// An additional obstacle position can be passed to the function
func doWalk(grid *aoc.Grid, additionalObstacle aoc.Coordinate, wantVisited bool) (bool, []aoc.Coordinate) {
	startPos := grid.Find('^')

	curPointer := aoc.NewPointer(startPos.X, startPos.Y, 0)
	currentContent := byte('^')
	visitedWithDirection := make(map[aoc.Pointer]struct{})

	for currentContent != '~' {
		if _, ok := visitedWithDirection[curPointer]; ok {
			return true, nil
		}

		visitedWithDirection[curPointer] = struct{}{}

		nextPos := curPointer.PeekMove()
		nextContent := grid.Get(nextPos, '~')
		if nextContent == '#' || nextPos == additionalObstacle {
			curPointer.TurnRight()
		} else {
			curPointer.Move()
			currentContent = nextContent
		}
	}

	// For part2 we don't need the visited list, so we allow to signal to skip creating it
	if !wantVisited {
		return false, nil
	}

	// First create a map of visited positions without direction
	visited := make(map[aoc.Coordinate]struct{})
	for pointer := range visitedWithDirection {
		visited[pointer.C] = struct{}{}
	}

	// Then convert the map to a slice
	visitedSlice := make([]aoc.Coordinate, 0, len(visited)-1)
	for coordinate := range visited {
		// Don't include the startpos, as for part2 we want to skip that
		// For part1 we can just add one in the caller function
		if coordinate != startPos {
			visitedSlice = append(visitedSlice, coordinate)
		}
	}

	return false, visitedSlice
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (int, int) {
	sumPart2 := 0

	grid := input.Grid()

	// Part1
	// We just need to walk until we exit the grid and count the visited positions
	_, visited := doWalk(grid, aoc.Coordinate{X: -2, Y: 0}, true)
	sumPart1 := len(visited) + 1

	// Part2
	// We can only set one obstacle at a time, so we only need to test the positions
	// visited in Part1 as potential obstacle positions
	for _, additionalObstaclePos := range visited {
		if loop, _ := doWalk(grid, additionalObstaclePos, false); loop {
			sumPart2++
		}
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
