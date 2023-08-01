//package _7_Regex_engine

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	vars := strings.Split(input, "|")
	rej, text := vars[0], vars[1]

	matched := rejEngine(rej, text)
	fmt.Println(matched)
}

func rejEngine(rej, text string) bool {
	rej = transformRej(rej)
	iterationCount := len(text) - len(rej) + specialCasesCnt(rej)
	for i := 0; i <= iterationCount; i++ {
		if matched := matchStr(rej, text, i); matched {
			return true
		}
	}
	return false
}

func transformRej(rej string) string {
	for strings.Index(rej, "\\\\") >= 0 {
		r := strings.Index(rej, "\\\\")
		followedByRej := r+1 < len(rej)-1 && strings.ContainsAny(string(rej[r+1]), "?*+")
		if followedByRej {
			rej = strings.Replace(rej, "\\\\", "\\"+string(rej[r+1]), 1)
		} else {
			rej = strings.Replace(rej, "\\\\", "\\", 1)
		}
	}
	return rej
}

func specialCasesCnt(rej string) int {
	cycleCount := 0
	cycleCount += strings.Count(rej, "?") * 2
	cycleCount += strings.Count(rej, "*") * 2
	cycleCount += strings.Count(rej, "+")
	cycleCount -= strings.Count(rej, "\\?")
	cycleCount -= strings.Count(rej, "\\*")
	cycleCount -= strings.Count(rej, "\\+")
	cycleCount += strings.Count(rej, "\\.")
	return cycleCount
}

func matchStr(rej, text string, startI int) bool {
	if rej == "" {
		return true
	}
	if text == "" || startI < 0 {
		return false
	}

	if r := strings.Index(rej, "?"); r >= 0 && !hasEscape(rej, r) {
		return matchQuestion(rej, text, startI, r)
	}
	if r := strings.Index(rej, "*"); r >= 0 && !hasEscape(rej, r) {
		return matchAsterisk(rej, text, startI, r)
	}
	if r := strings.Index(rej, "+"); r >= 0 && !hasEscape(rej, r) {
		return matchPlus(rej, text, startI, r)
	}

	return matchAffixes(rej, text, startI)
}

func isRejChar(rej string, i int) bool {
	return strings.ContainsAny(string(rej[i]), ".?+*")
}

func hasEscape(rej string, i int) bool {
	return i > 0 && rej[i-1] == '\\'
}

func followedByRej(rej string, i int) bool {
	return i < len(rej)-1 && isRejChar(rej, i+1)
}

func proceededByRej(rej string, i int) bool {
	return i > 0 && isRejChar(string(rej[i-1]), i)
}

func matchQuestion(rej, text string, startI, r int) bool {
	return matchRepeated(rej, text, startI, r, 0, 1)
}

func matchAsterisk(rej, text string, startI, r int) bool {
	return matchRepeated(rej, text, startI, r, 0, len(text))
}

func matchPlus(rej, text string, startI, r int) bool {
	return matchRepeated(rej, text, startI, r, 1, len(text))
}

func matchRepeated(rej, text string, startI, r, repeatStart, repeatEnd int) bool {
	left := rej[:r-1]
	right := ""
	if r < len(rej)-1 {
		right += rej[r+1:]
	}
	char := string(rej[r-1])
	for i := repeatStart; i+startI < len(text) && i <= repeatEnd; i++ {
		newRej := left + strings.Repeat(char, i) + right
		if matchStr(newRej, text, startI) {
			return true
		}
	}
	return false
}

func matchAffixes(rej, text string, startI int) bool {
	hasPrefix, hasSuffix := rej[0] == '^', rej[len(rej)-1] == '$'
	if hasPrefix && hasSuffix {
		rej = rej[1 : len(rej)-1]
		if len(rej) != len(text) {
			return false
		}
	} else if hasPrefix {
		rej = rej[1:]
		return matchStr(rej, text, 0)
	} else if hasSuffix {
		rej = rej[:len(rej)-1]
		return matchStr(rej, text, len(text)-len(rej))
	}
	return matchCharByChar(rej, text, startI)
}

func matchCharByChar(rej, text string, startI int) bool {
	for i, char := range rej {
		isSlash := char == '\\'
		if isSlash && followedByRej(rej, i) {
			if strings.ContainsAny(rej, "?*+") {
				startI--
			}
			continue
		}
		isAny := char == '.' && !hasEscape(rej, i)
		if isAny || char == rune(text[i+startI]) {
			continue
		}
		return false
	}
	return true
}
