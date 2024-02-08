package _0_behcnmarks_test

import (
	"bufio"
	"os"
	"testing"
)

// how to run the benchmarks:
// go test -bench=.
// how to run the benchmarks with specified time:
// go test -bench=. -benchtime=3s
// how to run the benchmarks with specified count:
// go test -bench=. -count=3

// Benchmark for approach 1
func BenchmarkApproach1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bufio.NewReader(os.Stdin)
		_ = bufio.NewWriter(os.Stdout)
	}
}

// Benchmark for approach 2 (even though it's effectively the same)
func BenchmarkApproach2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var _ *bufio.Reader
		var _ *bufio.Writer
		_ = bufio.NewReader(os.Stdin)
		_ = bufio.NewWriter(os.Stdout)
	}
}
