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
	var line string
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &line)
		matches, ok := checkLicensesLine(line)
		if ok {
			fmt.Fprintln(out, strings.Join(matches, " "))
		} else {
			fmt.Fprintln(out, "-")
		}
	}
}

func checkLicensesLine(licenses string) ([]string, bool) {
	regex := regexp.MustCompile(`^([A-Z]\d\d[A-Z][A-Z]|[A-Z]\d[A-Z][A-Z])`)
	var matches []string
	for len(licenses) > 0 {
		match := regex.FindString(licenses)
		if match == "" {
			return nil, false
		}
		matches = append(matches, match)
		licenses = licenses[len(match):]
	}
	return matches, true
}
