
package main

import (
    "fmt"
    "time"
)

func pMod(a int, b int) int {
    return ((a % b) + b) % b;
}

func countNeighbors(world [][]bool, r int, c int) int {
    nNeighbors := 0

    for dy := -1; dy <= 1; dy++ {
        for dx := -1; dx <= 1; dx++ {
            y := pMod(r + dy, len(world))
            x := pMod(c + dx, len(world[y]))

            if y == r && x == c {
                continue
            }

            if world[y][x] {
                nNeighbors += 1;
            }
        }
    }

    return nNeighbors
}

func initWorld(rows uint, cols uint) ([][]bool, error) {
    if rows == 0 || cols == 0 {
        return nil, fmt.Errorf("Cannot construct new world with zero row or column")
    }

    world := make([][]bool, rows)

    for i := range rows {
        world[i] = make([]bool, cols);
    }

    return world, nil
}

func checkDims(a [][]bool, b [][]bool) error {
    aRows := len(a)
    bRows := len(a)

    if aRows == 0 || bRows == 0 {
        return fmt.Errorf("Zero rows detected")
    }

    aCols := len(a[0])
    bCols := len(a[0])

    if aCols == 0 || bCols == 0 {
        return fmt.Errorf("Zero cols detected")
    }

    if aRows != bRows {
        return fmt.Errorf("Mismatched rows")
    }

    if aCols != bCols {
        return fmt.Errorf("Mismatched columns")
    }

    return nil
}

func nextState(newWorld [][]bool, world [][]bool) error {
    err := checkDims(newWorld, world)
    if err != nil {
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

            if (neighbors == 3 || (alive && neighbors == 2)) {
                newWorld[r][c] = true
            } else {
                newWorld[r][c] = false
            }
        }
    }

    return nil
}

func printWorld(world [][]bool) {
    for _, row := range world {
        for _, cell := range row {
            if cell {
                fmt.Print("#")
            } else {
                fmt.Print(".")
            }
        }
        fmt.Print("\n")
    }
}

func mainLoop() {
    const nCols = 10
    const nRows = 10

    worldA, err := initWorld(nRows, nCols)
    if err != nil {
        fmt.Printf("Error creating worldA: %s\n", err)
        return
    }
    worldB, err := initWorld(nRows, nCols)
    if err != nil {
        fmt.Printf("Error creating worldB: %s\n", err)
        return
    }

    // glider
    //   012
    // 0  #
    // 1   #
    // 2 ###
    worldA[0][1] = true
    worldA[1][2] = true
    worldA[2][0] = true
    worldA[2][1] = true
    worldA[2][2] = true

    err = nil
    for i := range(100) {
        fmt.Printf("World [%s]:\n", i)
        if i % 2 == 0 {
            printWorld(worldA)
            err = nextState(worldB, worldA)
        } else {
            printWorld(worldB)
            err = nextState(worldA, worldB)
        }

        if err != nil {
            fmt.Printf("Error creating world: %s\n", err)
            return
        }
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    fmt.Println("Conway's game of life")
    mainLoop()
}

