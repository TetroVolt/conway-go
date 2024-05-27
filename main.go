package main

import (
	"fmt"
	"time"
)

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

func initWorld(rows, cols int) [][]bool {
	world := make([][]bool, rows)
	for i := range world {
		world[i] = make([]bool, cols)
	}
	return world
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

func nextState(newWorld, world [][]bool) error {
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
			newWorld[r][c] = neighbors == 3 || (alive && neighbors == 2)
		}
	}

	return nil
}

func printWorld(world [][]bool) {
	for _, row := range world {
		for _, cell := range row {
			if cell {
				fmt.Print("O")
			} else {
				// fmt.Print(".")
				fmt.Print("·")
			}
		}
		fmt.Println()
	}
}

func renderAndComputeNextWorld(nextWorld, world [][]bool) error {
	printWorld(world)
	if err := nextState(nextWorld, world); err != nil {
		return err
	}
	return nil
}

func mainLoop() {
	const (
		nCols = 10
		nRows = 10
	)

	worldA := initWorld(nRows, nCols)
	worldB := initWorld(nRows, nCols)

	// glider
	worldA[0][1] = true
	worldA[1][2] = true
	worldA[2][0] = true
	worldA[2][1] = true
	worldA[2][2] = true

	for i := 0; i < 100; i++ {
		fmt.Printf("World Iteration [%d]:\n", i)

		var (
			nextWorld [][]bool = nil
			world     [][]bool = nil
		)

		if i%2 == 0 {
			nextWorld = worldB
			world = worldA
		} else {
			nextWorld = worldA
			world = worldB
		}

		if err := renderAndComputeNextWorld(nextWorld, world); err != nil {
			fmt.Printf("An error occurred computing next world: %s\n", err)
			return
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Conway's game of life")

	mainLoop()
}
