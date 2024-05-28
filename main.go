package main

import (
	"conway/gol"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Conway's game of life")

	const (
		nCols      = uint(10)
		nRows      = uint(10)
		genertions = uint(10)
	)

	golLoop(nCols, nRows, genertions)
}

func golLoop(nCols, nRows, gens uint) {
	worldA := gol.NewWorld(nRows, nCols)
	worldB := gol.NewWorld(nRows, nCols)

	// glider
	gol.SpawnGlider(worldA, 0, 0)

	var (
		nextWorld [][]bool
		world     [][]bool
	)

	for i := range gens {
		// manage buffers
		if i%2 == 0 {
			world, nextWorld = worldA, worldB
		} else {
			world, nextWorld = worldB, worldA
		}

		// render world
		fmt.Printf("Generation [%d]:\n", i+1)
		gol.PrintWorld(world)

		// compute next generation
		err := gol.NextState(nextWorld, world)

		if err != nil {
			fmt.Printf("An error occurred computing next world: %s\n", err)
			return
		}

		time.Sleep(50 * time.Millisecond)
	}
}
