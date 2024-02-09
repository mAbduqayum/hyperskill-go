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

			// Redirect stdin and stdout
			_, stdinW, stdoutR, stdoutW, cleanup := redirectStdinStdout()
			defer cleanup()

			// Write input data to stdin
			stdinW.WriteString(string(input))
			stdinW.Close()

			// Call the main function in a goroutine to allow cleanup afterward
			done := make(chan bool)
			go func() {
				main()
				done <- true
			}()

			// Wait for main to finish
			<-done

			// Close the write end of stdout to signal to ReadAll that we're done writing
			stdoutW.Close()

			// Capture the output from stdout
			outBytes, err := ioutil.ReadAll(stdoutR)
			if err != nil {
				t.Fatalf("Failed to read from stdout: %v", err)
			}

			// Compare the output with the expected output
			if strings.TrimSpace(string(outBytes)) != strings.TrimSpace(string(expectedOutput)) {
				t.Errorf(
					"Test %s: Output did not match expected.\n"+
						"Expected:\t"+
						"########################################\n"+
						"%s\n"+
						"Got:\t"+
						"########################################\n"+
						"%s", file.Name(), string(expectedOutput), string(outBytes),
				)
			}
		})
	}
}

// redirectStdinStdout helps to replace os.Stdin and os.Stdout and returns the pipes and a cleanup function.
func redirectStdinStdout() (stdinR *os.File, stdinW *os.File, stdoutR *os.File, stdoutW *os.File, cleanup func()) {
	oldStdin := os.Stdin
	oldStdout := os.Stdout

	stdinR, stdinW, _ = os.Pipe()
	stdoutR, stdoutW, _ = os.Pipe()

	os.Stdin = stdinR
	os.Stdout = stdoutW

	return stdinR, stdinW, stdoutR, stdoutW, func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
		stdinR.Close()
		stdinW.Close()
		stdoutR.Close()
		stdoutW.Close()
	}
}
