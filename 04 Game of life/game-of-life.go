//package _4_game_of_life

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type World [][]bool

const (
	Dead  = " "
	Alive = "O"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var seed int64
	var size, generation int
	seed = 1
	generation = 10

	// Scan for the first line that contains n and seed
	if scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		size, _ = strconv.Atoi(parts[0])
		//seed, _ = strconv.ParseInt(parts[1], 10, 64)
		//generation, _ = strconv.Atoi(parts[2])
	}

	world := generateWorld(size, seed)
	for i := 1; i <= generation; i++ {
		world = nextGeneration(world)
		fmt.Printf("Generation #%d\n", i)
		beautify(world)
		//time.Sleep(500 * time.Millisecond)
		//fmt.Print("\033[H\033[2J")
	}
	beautify(world)
}

func beautify(world World) {
	for _, row := range world {
		for _, r := range row {
			if r {
				fmt.Print(Alive)
			} else {
				fmt.Print(Dead)
			}
		}
		fmt.Println()
	}
}

func generateWorld(size int, seed int64) World {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano() + seed))
	world := make(World, size)
	for i := 0; i < size; i++ {
		world[i] = make([]bool, size)
		for j := 0; j < size; j++ {
			world[i][j] = rnd.Intn(2) == 1
		}
	}
	return world
}

func nextGeneration(world World) World {
	nextWorld := make(World, len(world))
	for i, row := range world {
		nextWorld[i] = make([]bool, len(world))
		for j, r := range row {
			neighbors := countNeighbors(world, i, j)
			nextWorld[i][j] = false
			if r && (neighbors == 2 || neighbors == 3) {
				nextWorld[i][j] = true
			}
			if !r && neighbors == 3 {
				nextWorld[i][j] = true
			}
		}
	}
	return nextWorld
}

func countNeighbors(world World, x, y int) int {
	boundary := len(world) - 1
	cnt := 0
	if x > 0 && y > 0 && world[x-1][y-1] {
		cnt++
	}
	if x > 0 && world[x-1][y] {
		cnt++
	}
	if x > 0 && y < boundary && world[x-1][y+1] {
		cnt++
	}
	if y > 0 && world[x][y-1] {
		cnt++
	}
	if y < boundary && world[x][y+1] {
		cnt++
	}
	if x < boundary && y > 0 && world[x+1][y-1] {
		cnt++
	}
	if x < boundary && world[x+1][y] {
		cnt++
	}
	if x < boundary && y < boundary && world[x+1][y+1] {
		cnt++
	}
	return cnt
}
