package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()
	processTest()
}

type sticker struct {
	s     string
	start int
	end   int
}

func processTest() {
	var s string
	fmt.Fscan(in, &s)

	var n int
	fmt.Fscan(in, &n)

	var stickers = make([]sticker, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &stickers[i].start, &stickers[i].end, &stickers[i].s)
	}

	var bytes = []byte(s)
	for _, s := range stickers {
		for i := s.start; i <= s.end; i++ {
			bytes[i-1] = s.s[i-s.start]
		}
	}
	fmt.Fprintln(out, string(bytes))
}
