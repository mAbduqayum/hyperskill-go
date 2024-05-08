package main_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testDir      = "tests"
	answerSuffix = ".a"
)

var reInputFile = regexp.MustCompile(`^\d+$`)

func TestIt(t *testing.T) {
	mainFile, err := findMainFile()
	if !assert.NoError(t, err, "Failed to find main file") {
		return
	}

	testCases, err := getTestCases(testDir)
	if !assert.NoError(t, err, "Failed to get test cases") {
		return
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualOutput, err := runMainPackage(mainFile, tc.input)
			if !assert.NoErrorf(t, err, "Failed to run main package with file: '%s'", tc.name) {
				return
			}
			actualOutputString := strings.TrimSpace(string(actualOutput))
			expectedOutputString := strings.TrimSpace(string(tc.expectedOutput))
			if !assert.Equal(t, expectedOutputString, actualOutputString, "Test case %s failed", tc.name) {
				return
			}
		})
	}
}

func findMainFile() (string, error) {
	files, err := os.ReadDir(".")
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".go" {
			content, err := os.ReadFile(file.Name())
			if err == nil && bytes.Contains(content, []byte("package main")) {
				return file.Name(), nil
			}
		}
	}

	return "", fmt.Errorf("no Go file with 'package main' found in the current directory")
}

type testCase struct {
	name           string
	input          []byte
	expectedOutput []byte
}

func getTestCases(dir string) ([]testCase, error) {
	testFiles, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var testCases []testCase
	for _, file := range testFiles {
		if !reInputFile.MatchString(file.Name()) {
			continue
		}

		inputFile := filepath.Join(dir, file.Name())
		answerFile := filepath.Join(dir, file.Name()+answerSuffix)

		input, err := os.ReadFile(inputFile)
		if err != nil {
			return nil, err
		}

		expectedOutput, err := os.ReadFile(answerFile)
		if err != nil {
			return nil, err
		}

		testCases = append(testCases, testCase{
			name:           file.Name(),
			input:          input,
			expectedOutput: expectedOutput,
		})
	}

	return testCases, nil
}

func runMainPackage(fileName string, input []byte) ([]byte, error) {
	cmd := exec.Command("go", "run", fileName)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, string(input))
	}()

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return output, nil
}
