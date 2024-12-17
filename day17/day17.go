package main

import (
	"math"
	"slices"
	"strconv"

	"github.com/julijane/advent-of-code-2024/aoc"
)

const (
	regA = 0
	regB = 1
	regC = 2

	instADV = 0
	instBXL = 1
	instBST = 2
	instJNZ = 3
	instBXC = 4
	instOUT = 5
	instBDV = 6
	instCDV = 7
)

type CPU struct {
	registers [3]int
	ip        int
	program   []int
	output    []int
}

func NewCPU(input *aoc.Input) *CPU {
	lines := input.PlainLines()

	a := aoc.ExtractNumbers(lines[0])[0]
	b := aoc.ExtractNumbers(lines[1])[0]
	c := aoc.ExtractNumbers(lines[2])[0]

	program := aoc.ExtractNumbers(lines[4])

	ip := 0

	return &CPU{
		registers: [3]int{a, b, c},
		ip:        ip,
		program:   program,
	}
}

func (cpu *CPU) getComboParam(param int) int {
	if param <= 3 {
		return param
	}

	return cpu.registers[param-4]
}

func (cpu *CPU) Run() []int {
instloop:
	for cpu.ip < len(cpu.program) {
		opcode := cpu.program[cpu.ip]
		param := cpu.program[cpu.ip+1]

		switch opcode {
		case instADV:
			n := cpu.getComboParam(param)
			cpu.registers[regA] = cpu.registers[regA] >> n
		case instBXL:
			cpu.registers[regB] = cpu.registers[regB] ^ param
		case instBST:
			cpu.registers[regB] = cpu.getComboParam(param) & 0x7
		case instJNZ:
			if cpu.registers[regA] != 0 {
				cpu.ip = param
				continue instloop
			}
		case instBXC:
			cpu.registers[regB] = cpu.registers[regB] ^ cpu.registers[regC]
		case instOUT:
			cpu.output = append(cpu.output, cpu.getComboParam(param)&0x7)
		case instBDV:
			n := cpu.getComboParam(param)
			cpu.registers[regB] = cpu.registers[regA] >> n
		case instCDV:
			n := cpu.getComboParam(param)
			cpu.registers[regC] = cpu.registers[regA] >> n
		}

		cpu.ip += 2
	}

	return cpu.output
}

type partialA struct {
	partialA    int
	numSections int
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (any, any) {
	resultPart2 := 0

	cpu := NewCPU(input)

	// save for part2
	startB := cpu.registers[regB]
	startC := cpu.registers[regC]

	resultPart1Ints := cpu.Run()
	resultPart1 := ""
	for i, val := range resultPart1Ints {
		if i > 0 {
			resultPart1 += ","
		}
		resultPart1 += strconv.Itoa(val)
	}

	// for part2 we make the following assumptions which hold true for all examples and the regular input of today:
	//
	// 1. There is only one jnz instruction and it is the last instruction of the program.
	// 2. The ADV instruction(s) only use a literal value as operand.

	if doPart2 {
		resultPart2 = math.MaxInt64

		expectedOutputFull := slices.Clone(cpu.program)

		// we need to determine the start value for register A. We can do that in iterations where we work our way back from the
		// end of the expected output and determine a number of bits of A. How many bits of A we need per iteration/output depends
		// on how many bits of A are removed by ADV instruction(s) per loop iteration. This could have been hardcoded for the task,
		// but we calculate it here to make the solution more generic.

		numABitsPerOutput := 0
		for ip := 0; ip < len(cpu.program); ip += 2 {
			if cpu.program[ip] == instADV {
				operand := cpu.program[ip+1]
				if operand > 3 {
					panic("part2: ADV instruction must not use a register as operand")
				}
				numABitsPerOutput += operand
			}
		}

		maxValueForABits := 1 << numABitsPerOutput

		// Now we go back from the last output byte and find which values of A would have produced this output
		// and save this in our check queue. We then continue for the prior output bytes by bit shifting the
		// value of A found so far accordingly and try all possible values for the next section of A. Repeat
		// until done.

		// A queue of partial values of A. We need to
		checkQueue := []partialA{{0, 0}}

		for len(checkQueue) > 0 {
			candidate := checkQueue[0]
			checkQueue = checkQueue[1:]

			startA := candidate.partialA << numABitsPerOutput

			// we check all possible values for the current section of A
			for j := 0; j < maxValueForABits; j++ {
				// reset CPU state
				cpu.registers[regA] = startA | j
				cpu.registers[regB] = startB
				cpu.registers[regC] = startC
				cpu.output = []int{}
				cpu.ip = 0

				output := cpu.Run()

				if len(output) != candidate.numSections+1 {
					// too short
					continue
				}

				expectedOutput := expectedOutputFull[len(expectedOutputFull)-len(cpu.output):]

				if slices.Equal(output, expectedOutput) {
					if len(expectedOutputFull) == len(cpu.output) {
						if startA|j < resultPart2 {
							resultPart2 = startA | j
						}
						continue
					}

					checkQueue = append(checkQueue, partialA{startA | j, candidate.numSections + 1})
				}
			}
		}

	}

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, false)
	aoc.Run("sample2.txt", calc, false, true)
	aoc.Run("input.txt", calc, true, true)
}
