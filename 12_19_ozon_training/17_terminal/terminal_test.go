package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestTerminalFromFile(t *testing.T) {
	testsDir := "./tests"
	testFilePattern := regexp.MustCompile(`^\d+$`) // Pattern to match test files with digits only

	inputFiles, err := ioutil.ReadDir(testsDir)
	if err != nil {
		t.Fatalf("Failed to find test files: %v", err)
	}

	for _, file := range inputFiles {
		if file.IsDir() || !testFilePattern.MatchString(file.Name()) {
			continue // Skip directories and files that do not match the pattern
		}

		t.Run(file.Name(), func(t *testing.T) {
			inputFilePath := filepath.Join(testsDir, file.Name())
			expectedOutputFilePath := filepath.Join(testsDir, file.Name()+".a")

			input, err := ioutil.ReadFile(inputFilePath)
			if err != nil {
				t.Fatalf("Failed to read input file %s: %v", inputFilePath, err)
			}

			expectedOutput, err := ioutil.ReadFile(expectedOutputFilePath)
			if err != nil {
				t.Fatalf("Failed to read expected output file %s: %v", expectedOutputFilePath, err)
			}

			stdinW, stdoutR, stdoutW, cleanup := redirectStdinStdout()
			defer cleanup()

			stdinW.Write(input)
			stdinW.Close()

			// Call the main function in a goroutine to allow stdout to be read concurrently
			done := make(chan bool)
			go func() {
				main()
				stdoutW.Close() // Close stdoutW after main returns
				done <- true
			}()

			// Read the captured output from stdout
			outBytes, err := ioutil.ReadAll(stdoutR)
			if err != nil {
				t.Fatalf("Failed to read from stdout: %v", err)
			}

			// Wait for main to finish
			<-done

			actualOutput := normalizeOutput(outBytes)
			expectedOutputString := normalizeOutput(expectedOutput)

			if actualOutput != expectedOutputString {
				t.Errorf("Test %s: Output did not match expected.\nExpected:\n%s\nGot:\n%s", file.Name(), expectedOutputString, actualOutput)
			}
		})
	}
}

func redirectStdinStdout() (*os.File, *os.File, *os.File, func()) {
	oldStdin := os.Stdin
	oldStdout := os.Stdout

	stdinW, _, _ := os.Pipe()
	stdoutR, stdoutW, _ := os.Pipe()

	os.Stdin = stdinW // We use the write end of stdin pipe as the new stdin.
	os.Stdout = stdoutW

	return stdinW, stdoutR, stdoutW, func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
		stdinW.Close()
		stdoutR.Close()
		stdoutW.Close()
	}
}

func normalizeOutput(output []byte) string {
	normalized := strings.ReplaceAll(string(output), "\r\n", "\n")
	normalized = strings.ReplaceAll(normalized, "\r", "\n")
	normalized = strings.TrimSpace(normalized)
	lines := strings.Split(normalized, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return strings.Join(lines, "\n")
}
