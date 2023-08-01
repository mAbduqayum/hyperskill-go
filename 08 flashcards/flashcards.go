//package _8_flashcards

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var stringLog strings.Builder

func main() {
	load := flag.String("import-from", "export.txt", "Enter file name: ")
	export := flag.String("export-to", "export.txt", "Enter file name: ")
	flag.Parse()

	fmt.Printf("load file name: %s\n", *load)
	fmt.Printf("export file name: %s\n", *export)
	cards := NewCards()
	if *load != "" {
		cards.Import(*load)
	}
	if *export != "" {
		cards.Export(*export)
	}
	processCard(*cards)
}

func processCard(cards Cards) {
	for {
		fmt.Println("Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats):")
		_, _ = fmt.Fprintln(&stringLog, "Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats):")
		action := input()
		switch action {
		case "add":
			cards.Add()
		case "remove":
			cards.Remove()
		case "import":
			fmt.Println("File name:")
			_, _ = fmt.Fprintln(&stringLog, "File name:")
			fileName := input()
			cards.Import(fileName)
		case "export":
			fmt.Println("File name:")
			_, _ = fmt.Fprintln(&stringLog, "File name:")
			fileName := input()
			cards.Export(fileName)
		case "ask":
			cards.Ask()
		case "log":
			cards.Log()
		case "hardest card":
			cards.Hardest()
		case "reset stats":
			cards.Reset()
		case "exit":
			fmt.Println("Bye bye!")
			_, _ = fmt.Fprintln(&stringLog, "Bye bye!")
			return
		}
		fmt.Println()
		_, _ = fmt.Fprintln(&stringLog)
	}

}

type Cards struct {
	Cards    map[string]string `json:"cards"`
	Mistakes map[string]int    `json:"mistakes"`
}

func NewCards() *Cards {
	return &Cards{
		Cards:    make(map[string]string),
		Mistakes: make(map[string]int),
	}
}

func (c Cards) UpdateHardest(key string) {
	cnt := 0
	if val, ok := c.Mistakes[key]; ok {
		cnt = val
	}
	cnt++
	c.Mistakes[key] = cnt
}

func (c Cards) Add() {
	fmt.Println("The card:")
	_, _ = fmt.Fprintln(&stringLog, "The card:")
	term := input()
	_, ok := c.Cards[term]
	for ok {
		fmt.Printf("The card \"%s\" already exists. Try again:\n", term)
		_, _ = fmt.Fprintf(&stringLog, "The card \"%s\" already exists. Try again:\n", term)
		term = input()
		_, ok = c.Cards[term]
	}
	fmt.Println("The definition of the card:")
	_, _ = fmt.Fprintln(&stringLog, "The definition of the card:")
	def := input()
	_, ok = c.Key(def)
	for ok {
		fmt.Printf("The definition \"%s\" already exists. Try again:\n", def)
		_, _ = fmt.Fprintf(&stringLog, "The definition \"%s\" already exists. Try again:\n", def)
		def = input()
		_, ok = c.Key(def)
	}
	c.Cards[term] = def
	fmt.Printf("The pair (\"%s\":\"%s\") has been added.\n", term, def)
	_, _ = fmt.Fprintf(&stringLog, "The pair (\"%s\":\"%s\") has been added.\n", term, def)
}

func (c Cards) Remove() {
	fmt.Println("Which card?")
	_, _ = fmt.Fprintln(&stringLog, "Which card?")
	term := input()
	if _, ok := c.Cards[term]; !ok {
		fmt.Printf("Can't remove \"%s\": there is no such card.\n", term)
		_, _ = fmt.Fprintf(&stringLog, "Can't remove \"%s\": there is no such card.\n", term)
		return
	}
	delete(c.Cards, term)
	fmt.Println("The card has been removed.")
	_, _ = fmt.Fprintln(&stringLog, "The card has been removed.")
}

func (c Cards) Import(fileName string) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("File not found.")
		_, _ = fmt.Fprintln(&stringLog, "File not found.")
		return
	}
	var newCards Cards
	_ = json.Unmarshal(file, &newCards)
	for key, val := range newCards.Cards {
		c.Cards[key] = val
	}
	fmt.Printf("%d cards have been loaded.\n", len(newCards.Cards))
	_, _ = fmt.Fprintf(&stringLog, "%d cards have been loaded.\n", len(newCards.Cards))
}

func (c Cards) Export(fileName string) {
	rez, _ := json.Marshal(&c)
	_ = os.WriteFile(fileName, rez, 0666)
	fmt.Printf("%d cards have been saved.\n", len(c.Cards))
	_, _ = fmt.Fprintf(&stringLog, "%d cards have been saved.\n", len(c.Cards))
}

func (c Cards) Ask() {
	if len(c.Cards) == 0 {
		fmt.Println("No Cards!")
		_, _ = fmt.Fprintln(&stringLog, "No Cards!")
		return
	}
	fmt.Println("How many times to ask?")
	_, _ = fmt.Fprintln(&stringLog, "How many times to ask?")
	times, _ := strconv.Atoi(input())
	for i := 0; i < times; i++ {
		question, answer := c.RandomTerm()
		fmt.Printf("Print the definition of \"%s\":\n", question)
		_, _ = fmt.Fprintf(&stringLog, "Print the definition of \"%s\":\n", question)
		userAnswer := input()
		if userAnswer == answer {
			fmt.Println("Correct!")
			_, _ = fmt.Fprintln(&stringLog, "Correct!")
		} else {
			c.UpdateHardest(question)
			userQuestion, ok := c.Key(userAnswer)
			if ok {
				fmt.Printf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\".\n", answer, userQuestion)
				_, _ = fmt.Fprintf(&stringLog, "Wrong. The right answer is \"%s\", but your definition is correct for \"%s\".\n", answer, userQuestion)
			} else {
				fmt.Printf("Wrong. The right answer is \"%s\".\n", answer)
				_, _ = fmt.Fprintf(&stringLog, "Wrong. The right answer is \"%s\".\n", answer)
			}
		}
	}
}

func (c Cards) RandomTerm() (string, string) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	keys := make([]string, 0, len(c.Cards))
	for k := range c.Cards {
		keys = append(keys, k)
	}
	randIndex := rnd.Intn(len(keys))
	randomKey := keys[randIndex]
	randomValue := c.Cards[randomKey]
	return randomKey, randomValue
}

func (c Cards) Key(def string) (string, bool) {
	for key, val := range c.Cards {
		if val == def {
			return key, true
		}
	}
	return "", false
}

func (c Cards) Log() {
	fmt.Println("File name:")
	_, _ = fmt.Fprintln(&stringLog, "File name:")
	fileName := input()
	_ = os.WriteFile(fileName, []byte(stringLog.String()), 0666)
	fmt.Println("The log has been saved.")
	_, _ = fmt.Fprintln(&stringLog, "The log has been saved.")
}

func (c Cards) Hardest() {
	hards := make([]string, 0)
	for key, val := range c.Mistakes {
		if len(hards) == 0 {
			hards = append(hards, key)
			continue
		}
		hardestCount := c.Mistakes[hards[0]]
		if val > hardestCount {
			hards = make([]string, 0)
			hards = append(hards, key)
		}
	}

	if len(c.Mistakes) == 0 || c.Mistakes[hards[0]] == 0 {
		fmt.Println("There are no cards with errors.")
		_, _ = fmt.Fprintln(&stringLog, "There are no cards with errors.")
		return
	}
	hardestCount := c.Mistakes[hards[0]]
	rez := strings.Join(hards, ", ")
	if len(hards) == 1 {
		fmt.Printf("The hardest card is \"%s\". You have %d errors answering it.", rez, hardestCount)
		_, _ = fmt.Fprintf(&stringLog, "The hardest card is \"%s\". You have %d errors answering it.", rez, hardestCount)
		return
	}
	fmt.Printf("The hardest cards are \"%s\". You have %d errors answering them.\n", rez, hardestCount)
	_, _ = fmt.Fprintf(&stringLog, "The hardest cards are \"%s\". You have %d errors answering them.\n", rez, hardestCount)
}

func (c Cards) Reset() {
	c.Mistakes = make(map[string]int)
	fmt.Println("ard statistics have been reset.")
	_, _ = fmt.Fprintln(&stringLog, "ard statistics have been reset.")
}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	return func() string {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		return line
	}()
}
