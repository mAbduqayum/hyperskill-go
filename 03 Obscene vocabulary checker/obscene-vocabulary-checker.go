//package _3_obscene_vocabulary_checker

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	var fileName string
	fmt.Scanln(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a set to store taboo words
	tabooWords := make(map[string]struct{})

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		tabooWords[scanner.Text()] = struct{}{}
	}

	for {
		var inputSentence string
		fmt.Scanln(&inputSentence)
		if inputSentence == "exit" {
			break
		}
		fmt.Println(filterText(tabooWords, inputSentence))
	}
	fmt.Println("Bye!")
}

func filterText(tabooWords map[string]struct{}, sentence string) string {
	re := regexp.MustCompile(`[a-zA-Z]+`)
	sentence = re.ReplaceAllStringFunc(sentence, func(word string) string {
		if _, ok := tabooWords[strings.ToLower(word)]; ok {
			return strings.Repeat("*", len(word))
		}
		return word
	})
	return sentence
}
