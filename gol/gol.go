package gol

import (
	"fmt"
)

func NewWorld(rows, cols uint) [][]bool {
	world := make([][]bool, rows)
	for i := range world {
		world[i] = make([]bool, cols)
	}
	return world
}

func NextState(newWorld, world [][]bool) error {
	if err := checkDims(newWorld, world); err != nil {
		return err
	}

	/*
	   Truth table
	   alive, n == 2, n == 3      nextAlive
	     0       0      0             0
	     0       0      1             1
	     0       1      0             0

	     1       0      0             0
	     1       0      1             1
	     1       1      0             1
	*/

	for r, row := range world {
		for c, alive := range row {
			neighbors := countNeighbors(world, r, c)
			newWorld[r][c] =
				neighbors == 3 || (alive && neighbors == 2)
		}
	}

	return nil
}

func PrintWorld(world [][]bool) {
	// const (
	//     dead = "·"
	//     alive = "Θ"
	// )

	const (
		alive = "▣"
		dead  = "▢"
	)

	for _, row := range world {
		for _, cell := range row {
			if cell {
				fmt.Print(alive)
			} else {
				fmt.Print(dead)
			}
		}
		fmt.Println()
	}
}

func SpawnGlider(world [][]bool, r, c int) {
	// Glider
	//    0 1 2
	//  0 . O .
	//  1 . . O
	//  2 O O O

	glider := [][]bool{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}

	for i, row := range glider {
		for j, cell := range row {
			world[r+i][c+j] = cell
		}
	}
}

func positiveMod(a, b int) int {
	return ((a % b) + b) % b
}

func countNeighbors(world [][]bool, r, c int) int {
	neighbors := 0

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			y := positiveMod(r+dy, len(world))
			x := positiveMod(c+dx, len(world[y]))

			if y == r && x == c {
				continue
			}

			if world[y][x] {
				neighbors++
			}
		}
	}

	return neighbors
}

func checkDims(a, b [][]bool) error {
	if len(a) == 0 || len(b) == 0 || len(a[0]) == 0 || len(b[0]) == 0 {
		return fmt.Errorf("Zero rows or columns detected")
	}

	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return fmt.Errorf("Mismatched dimensions")
	}

	return nil
}
