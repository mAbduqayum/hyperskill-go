package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)

	result := make([]any, t)
	for i := 0; i < t; i++ {
		result[i] = processTest()
	}

	jsonResult, _ := json.Marshal(result)
	fmt.Fprintln(out, string(jsonResult))
}

func processTest() any {
	var n int
	fmt.Fscan(in, &n)

	in.ReadString('\n')
	var lines strings.Builder
	for i := 0; i < n; i++ {
		line, _ := in.ReadString('\n')
		lines.WriteString(line)
	}

	var jsonObj any
	json.Unmarshal([]byte(lines.String()), &jsonObj)
	jsonObj = prettify(jsonObj)
	return jsonObj
}

func prettify(obj any) any {
	switch v := obj.(type) {
	case map[string]any:
		return prettifyMap(v)
	case []any:
		return prettifySlice(v)
	default:
		return obj
	}
}

func prettifyMap[T map[string]any](obj T) T {
	for key, value := range obj {
		if isEmpty(value) {
			delete(obj, key)
		} else {
			obj[key] = prettify(value)
		}
	}
	return obj
}

func prettifySlice[T []any](obj T) T {
	result := make([]any, 0)
	for _, item := range obj {
		if !isEmpty(item) {
			result = append(result, prettify(item))
		}
	}
	return result
}

func isEmpty(obj any) bool {
	switch v := obj.(type) {
	case map[string]any:
		if len(v) == 0 {
			return true
		}
		for _, value := range v {
			if !isEmpty(value) {
				return false
			}
		}
		return true
	case []any:
		if len(v) == 0 {
			return true
		}
		for _, item := range v {
			if !isEmpty(item) {
				return false
			}
		}
		return true
	default:
		return false
	}
}
