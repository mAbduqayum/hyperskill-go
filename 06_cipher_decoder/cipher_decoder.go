//package _6_Cipher_decoder

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	var g, p = extractGP(input)
	fmt.Println("OK")

	input, _ = reader.ReadString('\n')
	A := extractA(input)
	b := 7
	B := modularPower(g, b, p)
	fmt.Println("B is", B)
	sharedSecretKey := modularPower(A, b, p)

	question := encryptMessage("Will you marry me?", sharedSecretKey)
	fmt.Println(question)

	aliceAnswer, _ := reader.ReadString('\n')
	aliceAnswer = strings.TrimSpace(aliceAnswer)
	decryptedAnswer := decryptMessage(aliceAnswer, sharedSecretKey)
	switch decryptedAnswer {
	case "Yeah, okay!":
		fmt.Println(encryptMessage("Great!", sharedSecretKey))
	case "Let's be friends.":
		fmt.Println(encryptMessage("What a pity!", sharedSecretKey))
	}
}

func encryptMessage(s string, key int) string {
	key %= 26
	var rez strings.Builder
	for _, letter := range s {
		if unicode.IsLetter(letter) {
			offset := key
			if unicode.IsLower(letter) {
				letter = rune((int(letter-'a')+offset)%26 + 'a')
			} else {
				letter = rune((int(letter-'A')+offset)%26 + 'A')
			}
		}
		rez.WriteRune(letter)
	}
	return rez.String()
}

func decryptMessage(s string, key int) string {
	key %= 26
	var rez strings.Builder
	for _, letter := range s {
		if unicode.IsLetter(letter) {
			offset := 26 - key
			if unicode.IsLower(letter) {
				letter = rune((int(letter-'a')+offset)%26 + 'a')
			} else {
				letter = rune((int(letter-'A')+offset)%26 + 'A')
			}
		}
		rez.WriteRune(letter)
	}
	return rez.String()
}

func modularPower(base, exponent, modulus int) int {
	if modulus == 1 {
		return 0
	}
	c := 1
	for ePrime := 0; ePrime < exponent; ePrime++ {
		c = (c * base) % modulus
	}
	return c
}

func extractGP(input string) (int, int) {
	split := strings.Fields(input)
	var g, p int

	g, _ = strconv.Atoi(split[2])
	p, _ = strconv.Atoi(split[6])
	return g, p
}

func extractA(input string) int {
	split := strings.Fields(input)
	a, _ := strconv.Atoi(split[2])
	return a
}
