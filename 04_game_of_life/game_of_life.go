package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type World struct {
	cells [][]bool
	size  int
}

const (
	deadSymbol  = " "
	aliveSymbol = "O"
)

func main() {
	size := flag.Int("size", 10, "Size of the world")
	seed := flag.Int64("seed", 1, "Seed for random generation")
	generations := flag.Int("gens", 50, "Number of generations")
	flag.Parse()

	world := NewWorld(*size, *seed)
	RunGame(world, *generations)
}

func NewWorld(size int, seed int64) *World {
	randSource := rand.New(rand.NewSource(time.Now().UnixNano() + seed))
	cells := make([][]bool, size)
	for i := range cells {
		cells[i] = make([]bool, size)
		for j := range cells[i] {
			cells[i][j] = randSource.Intn(2) == 1
		}
	}
	return &World{cells, size}
}

func RunGame(world *World, generations int) {
	for i := 1; i <= generations; i++ {
		world.Display()
		world.Evolve()
		time.Sleep(300 * time.Millisecond)
		ClearScreen()
	}
	world.Display()
}

func (w *World) Display() {
	for _, row := range w.cells {
		for _, alive := range row {
			if alive {
				fmt.Print(aliveSymbol)
			} else {
				fmt.Print(deadSymbol)
			}
		}
		fmt.Println()
	}
}

func (w *World) Evolve() {
	newCells := make([][]bool, w.size)
	for i := range w.cells {
		newCells[i] = make([]bool, w.size)
		for j := range w.cells[i] {
			aliveNeighbors := w.countAliveNeighbors(i, j)
			newCells[i][j] = w.cells[i][j] && (aliveNeighbors == 2 || aliveNeighbors == 3) ||
				!w.cells[i][j] && aliveNeighbors == 3
		}
	}
	w.cells = newCells
}

func (w *World) countAliveNeighbors(x, y int) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < w.size && ny >= 0 && ny < w.size && w.cells[nx][ny] {
				count++
			}
		}
	}
	return count
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
