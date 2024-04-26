package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, count, bound int
	var sign string
	fmt.Fscanln(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscanln(in, &count)
		lBound := 15
		uBound := 30
		valid := true
		for j := 0; j < count; j++ {
			fmt.Fscanln(in, &sign, &bound)
			if sign == ">=" && bound > lBound {
				lBound = bound
			} else if sign == "<=" && bound < uBound {
				uBound = bound
			}
			if lBound > uBound {
				valid = false
			}
			if !valid {
				fmt.Fprintln(out, -1)
			} else {
				fmt.Fprintln(out, lBound)
			}
		}
		fmt.Fprintln(out)
	}
}
