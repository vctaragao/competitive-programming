package main

import (
	"bufio"
	"fmt"
	"os"
)

type Instruction struct {
	hand  rune
	piece int
}

type Ring struct {
	R, L int
}

func (r *Ring) Print() {
	fmt.Printf("R: %d, L: %d\n", r.R, r.L)
}

func main() {
	var N, Q int
	fmt.Scanf("%d %d\n", &N, &Q)

	insts := make([]Instruction, Q)
	input := bufio.NewReader(os.Stdin)

	for i := 0; i < Q; i++ {
		inst := Instruction{}
		_, err := fmt.Fscanf(input, "%c %d\n", &inst.hand, &inst.piece)
		if err != nil {
			panic(err)
		}

		insts[i] = inst
	}

	sumSteps := 0
	ring := Ring{
		L: 1,
		R: 2,
	}

	for _, inst := range insts {
		steps := 0
		if inst.hand == 'R' {
			if inst.piece > ring.R {
				steps = move(N, ring.R, inst.piece, ring.L)
			} else {
				if ring.L < ring.R && ring.L > inst.piece {
					steps = N - (ring.R - inst.piece)
				} else {
					steps = ring.R - inst.piece
				}
			}

			ring.R = inst.piece
		}

		if inst.hand == 'L' {
			if inst.piece > ring.L {
				if ring.R > ring.L && ring.R < inst.piece {
					steps = ring.L - inst.piece + N
				} else {
					steps = inst.piece - ring.L
				}
			} else {
				if ring.R < ring.L && ring.R > inst.piece {
					steps = N - (ring.L - inst.piece)
				} else {
					steps = ring.L - inst.piece
				}
			}

			ring.L = inst.piece
		}

		sumSteps += steps
	}

	fmt.Println(sumSteps)
}

func move(N, from, to, ng int) int {
	if from > to {
		to, from = from, to
	}

	if from < ng && ng < to {
		return N + from - to
	}

	return to - from
}
