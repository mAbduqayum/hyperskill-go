package main_test

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

const (
	fileName     = "in_out.go"
	testDir      = "tests"
	outputSuffix = ".out"
	answerSuffix = ".a"
)

var (
	reInputFile = regexp.MustCompile(`^\d+$`)
)

func TestProgram(t *testing.T) {
	testFiles, err := os.ReadDir(testDir)
	if err != nil {
		t.Fatalf("Failed to read test directory: %v", err)
	}
	for _, testFile := range testFiles {
		if !reInputFile.MatchString(testFile.Name()) {
			continue // Skip files that don't match the regex
		}
		t.Run(testFile.Name(), func(t *testing.T) {
			testInputPath := filepath.Join(testDir, testFile.Name())
			testOutputPath := testInputPath + answerSuffix
			testInput, err := os.Open(testInputPath)
			if err != nil {
				t.Fatalf("Failed to open test input file %s: %v", testInputPath, err)
			}
			defer testInput.Close()
			testOutput, err := os.Open(testOutputPath)
			if err != nil {
				t.Fatalf("Failed to open test output file %s: %v", testOutputPath, err)
			}
			defer testOutput.Close()
			cmd := exec.Command("go", "run", fileName)
			cmd.Stdin = testInput
			cmd.Stderr = os.Stderr
			output, err := cmd.Output()
			if err != nil {
				t.Fatalf("Failed to run program for test %s: %v", testFile.Name(), err)
			}
			scannerOutput := bufio.NewScanner(strings.NewReader(string(output)))
			scannerExpected := bufio.NewScanner(testOutput)
			var lineNumber int
			for scannerOutput.Scan() && scannerExpected.Scan() {
				lineNumber++
				expected := scannerExpected.Text()
				actual := scannerOutput.Text()
				if actual != expected {
					saveProblematicFiles(testFile.Name(), output)
					t.Errorf("Test %s: Expected %s, found %s at line %d of %s",
						testFile.Name(), expected, actual, lineNumber, testOutputPath)
				}
			}
			if err := scannerOutput.Err(); err != nil {
				t.Fatalf("Failed to scan program output for test %s: %v", testFile.Name(), err)
			}
			if err := scannerExpected.Err(); err != nil {
				t.Fatalf("Failed to scan expected output for test %s: %v", testFile.Name(), err)
			}
			t.Logf("Test Passed: %s", testFile.Name())
		})
	}
}
func saveProblematicFiles(fileName string, output []byte) {
	outputPath := filepath.Join(testDir, fileName+outputSuffix)
	err := os.WriteFile(outputPath, output, 0644)
	if err != nil {
		fmt.Printf("Failed to save program output to %s: %v\n", outputPath, err)
		return
	}
	fmt.Printf("Program output saved as %s\n", outputPath)
}
