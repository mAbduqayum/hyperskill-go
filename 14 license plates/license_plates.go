package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	licensesLines := make([]string, t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &licensesLines[i])
		matched, matches := checkLicensesLine(licensesLines[i], []string{})
		if matched {
			fmt.Fprintln(out, strings.Join(matches, " "))
		} else {
			fmt.Fprintln(out, "-")
		}
	}
}

func checkLicensesLine(licenses string, matches []string) (bool, []string) {
	if len(licenses) < 4 {
		return false, []string{}
	}
	re1 := regexp.MustCompile(`^[A-Z]\d\d[A-Z][A-Z]`) // len 5
	re2 := regexp.MustCompile(`^[A-Z]\d[A-Z][A-Z]`)   // len 4

	match1 := re1.FindString(licenses)
	if len(match1) == 5 {
		matches = append(matches, match1)
		if len(licenses) == 5 {
			return true, matches
		}
		return checkLicensesLine(licenses[5:], matches)
	}
	match2 := re2.FindString(licenses)
	if len(match2) == 4 {
		matches = append(matches, match2)
		if len(licenses) == 4 {
			return true, matches
		}
		return checkLicensesLine(licenses[4:], matches)
	}
	return false, []string{}
}
