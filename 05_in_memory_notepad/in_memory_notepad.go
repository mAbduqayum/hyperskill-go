//package _5_In_memory_notepad

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	processCommands()
}

func processCommands() {
	maxMemos := readMaxMemos()
	memoPad := NewMemoPad(maxMemos)
	for {
		command, data := readAndParseInput()
		switch command {
		case "create":
			memoPad.add(data)
		case "update":
			index, data, err := parseData(data)
			if err != nil {
				continue
			}
			memoPad.update(index, data)
		case "delete":
			index, _, err := parseData(data)
			if err != nil {
				continue
			}
			memoPad.remove(index)
		case "list":
			memoPad.list()
		case "clear":
			memoPad.clear()
		case "exit":
			fmt.Println("[Info] Bye!")
			return
		default:
			fmt.Println("[Error] Unknown command")
		}
	}
}

type MemoPad struct {
	memos    []string
	maxMemos int
}

func NewMemoPad(maxMemos int) *MemoPad {
	return &MemoPad{
		memos:    make([]string, 0, maxMemos),
		maxMemos: maxMemos,
	}
}

func (mp *MemoPad) add(data string) {
	if len(mp.memos) >= mp.maxMemos {
		fmt.Println("[Error] Notepad is full")
		return
	}
	data = strings.TrimSpace(data)
	if data == "" {
		fmt.Println("[Error] Missing note argument")
		return
	}
	mp.memos = append(mp.memos, data)
	fmt.Println("[OK] The note was successfully created")
}

func (mp *MemoPad) update(index int, data string) {
	data = strings.TrimSpace(data)
	if data == "" {
		fmt.Println("[Error] Missing note argument")
		return
	}
	if index > mp.maxMemos || index < 0 {
		fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", index, mp.maxMemos)
		return
	}
	if index > len(mp.memos) {
		fmt.Println("[Error] There is nothing to update")
		return
	}
	mp.memos[index-1] = data
	fmt.Printf("[OK] The note at position %d was successfully updated\n", index)
}

func (mp *MemoPad) remove(index int) {
	if index > mp.maxMemos || index < 0 {
		fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", index, mp.maxMemos)
		return
	}
	if index > len(mp.memos) {
		fmt.Println("[Error] There is nothing to delete")
		return
	}
	mp.memos = append(mp.memos[:index-1], mp.memos[index:]...)
	fmt.Printf("[OK] The note at position %d was successfully deleted\n", index)
}

func (mp *MemoPad) list() {
	if len(mp.memos) == 0 {
		fmt.Println("[Info] Notepad is empty")
		return
	}
	for i, val := range mp.memos {
		fmt.Printf("[Info] %d: %s\n", i+1, val)
	}
}

func (mp *MemoPad) clear() {
	mp.memos = make([]string, 0, mp.maxMemos)
	fmt.Println("[OK] All notes were successfully deleted")
}

func parseData(data string) (int, string, error) {
	indexStr, data := parseInput(data)
	if strings.TrimSpace(indexStr) == "" {
		fmt.Println("[Error] Missing position argument")
		return 0, "", errors.New("missing position argument")
	}
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		fmt.Printf("[Error] Invalid position: %s\n", indexStr)
		return 0, "", err
	}
	return index, data, err
}

func readMaxMemos() int {
	fmt.Print("Enter the maximum number of notes: ")
	var maxMemos int
	fmt.Scanln(&maxMemos)
	return maxMemos
}

func readAndParseInput() (string, string) {
	input := readInput()
	command, data := parseInput(input)
	return command, data
}

func readInput() string {
	fmt.Print("Enter a command and data: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func parseInput(input string) (string, string) {
	parsed := strings.Split(input, " ")
	return parsed[0], strings.Join(parsed[1:], " ")
}
