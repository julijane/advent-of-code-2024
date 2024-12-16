package main

import (
	"container/heap"
	"math"
	"slices"

	"github.com/julijane/advent-of-code-2024/aoc"
)

type Entry struct {
	posWithDir aoc.Pointer
	cost       int
	path       []aoc.Pointer
}

// ---

type PriorityQueue []*Entry

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	field := x.(*Entry)
	*pq = append(*pq, field)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	field := old[n-1]
	*pq = old[0 : n-1]
	return field
}

// ---

func calc(input *aoc.Input, doPart1, doPart2 bool) (int, int) {
	sumPart1 := math.MaxInt64
	sumPart2 := 0

	grid := input.Grid()

	startPos := grid.Find('S')
	endPos := grid.Find('E')

	// we start facing right
	startPosWithDir := aoc.Pointer{C: startPos, Dir: 1}

	walkerSeen := make(map[aoc.Pointer]struct{})
	walkerSeen[startPosWithDir] = struct{}{}

	fieldsForFinalScore := make(map[int]map[aoc.Coordinate]struct{})

	pq := &PriorityQueue{}
	heap.Init(pq)

	pq.Push(&Entry{
		posWithDir: startPosWithDir,
		cost:       0,
	})

	for pq.Len() > 0 {
		entry := heap.Pop(pq).(*Entry)

		if entry.posWithDir.C == endPos {
			if entry.cost < sumPart1 {
				sumPart1 = entry.cost
			}

			if _, ok := fieldsForFinalScore[entry.cost]; !ok {
				fieldsForFinalScore[entry.cost] = make(map[aoc.Coordinate]struct{})
			}

			for _, pos := range entry.path {
				fieldsForFinalScore[entry.cost][pos.C] = struct{}{}
			}
			continue
		}

		walkerSeen[entry.posWithDir] = struct{}{}

		// we can either go ahead, turn right or turn left
		newEntries := []Entry{
			{posWithDir: aoc.Pointer{C: entry.posWithDir.PeekMove(), Dir: entry.posWithDir.Dir}, cost: entry.cost + 1},
			{posWithDir: aoc.Pointer{C: entry.posWithDir.C, Dir: (entry.posWithDir.Dir + 1) % 4}, cost: entry.cost + 1000},
			{posWithDir: aoc.Pointer{C: entry.posWithDir.C, Dir: (entry.posWithDir.Dir + 3) % 4}, cost: entry.cost + 1000},
		}

		for _, newEntry := range newEntries {
			if grid.Get(newEntry.posWithDir.C, '#') == '#' {
				continue
			}

			if _, ok := walkerSeen[newEntry.posWithDir]; ok {
				continue
			}

			newEntry.path = slices.Clone(entry.path)
			newEntry.path = append(newEntry.path, entry.posWithDir)

			heap.Push(pq, &newEntry)
		}
	}

	sumPart2 = len(fieldsForFinalScore[sumPart1]) + 1

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("sample2.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
