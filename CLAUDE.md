# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Structure

This is a Go programming exercise repository containing solutions to various programming challenges:

## Testing Framework

The repository uses a consistent testing pattern across problems:

### Running Tests

```bash
# Run tests for a specific problem
cd path/to/problem_directory
go test
```

### Test Structure

- Each problem directory contains:
    - `problem_name.go` - Main solution file
    - `problem_name_test.go` - Test file
    - `tests/` directory with numbered input/answer pairs
- Test files automatically discover the main Go file and run against all test cases in the `tests/` directory
- Input files are numbered (1, 2, 3, etc.) with corresponding answer files (1.a, 2.a, 3.a, etc.)

## Development Commands

## Code Patterns

### Never add comments

### Main Package Structure

Most solutions follow this pattern:

```go
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

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		solveTestCase()
	}
}

func solveTestCase() {
	var w, h int
	fmt.Fscan(in, &w, &h)
}
```
