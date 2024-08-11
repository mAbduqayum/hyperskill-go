package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := 0
	fmt.Sscanf(scanner.Text(), "%d", &t)

	results := make([]interface{}, 0, t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		n := 0
		fmt.Sscanf(scanner.Text(), "%d", &n)

		jsonStr := ""
		for j := 0; j < n; j++ {
			scanner.Scan()
			jsonStr += scanner.Text()
		}

		var data interface{}
		json.Unmarshal([]byte(jsonStr), &data)

		prettified := prettify(data)
		results = append(results, prettified)
	}

	output, _ := json.Marshal(results)
	fmt.Println(string(output))
}

func prettify(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, value := range v {
			prettified := prettify(value)
			if prettified != nil {
				result[key] = prettified
			}
		}
		if len(result) == 0 {
			return nil
		}
		return result
	case []interface{}:
		result := make([]interface{}, 0, len(v))
		for _, item := range v {
			prettified := prettify(item)
			if prettified != nil {
				result = append(result, prettified)
			}
		}
		if len(result) == 0 {
			return nil
		}
		return result
	default:
		return data
	}
}
