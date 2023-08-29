package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	space                     = ` `
	spacesPattern             = regexp.MustCompile(` +`)
	assignmentPattern         = regexp.MustCompile(` ?= ?`)
	numberAssignmentPattern   = regexp.MustCompile(`^[a-zA-Z]+ ?= ?\d+$`)
	variableAssignmentPattern = regexp.MustCompile(`^[a-zA-Z]+ ?= ?[a-zA-Z]+$`)
	variablePattern           = regexp.MustCompile(`^[a-zA-Z]+$`)
	operatorPattern           = regexp.MustCompile(`[*/^()+\-]`)
)

var precedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"^": 3,
	"(": 4,
	")": 4,
}

func main() {
	input := inputClosure()
	variables := make(map[string]int)
	for {
		userInput := input()
		userInput = spacesPattern.ReplaceAllString(userInput, space)
		if userInput == "" {
			continue
		}
		if strings.HasPrefix(userInput, "/") {
			switch userInput {
			case "/help":
				fmt.Println("The program evaluates basic algebraic expressions.")
				fmt.Println("The program first converts the expression to postfix notation,")
				fmt.Println("then it evaluates the expression using the stack")
			case "/exit":
				fmt.Println("Bye!")
				return
			default:
				fmt.Println("Unknown command")
			}
			continue
		}

		if variablePattern.MatchString(userInput) {
			if val, ok := variables[userInput]; ok {
				fmt.Println(val)
			} else {
				fmt.Println("Unknown variable")
			}
			continue
		}
		if numberAssignmentPattern.MatchString(userInput) {
			values := assignmentPattern.Split(userInput, -1)
			newVariable, _ := strconv.Atoi(values[1])
			variables[values[0]] = newVariable
			continue
		}
		if variableAssignmentPattern.MatchString(userInput) {
			values := assignmentPattern.Split(userInput, -1)
			if val, ok := variables[values[1]]; ok {
				variables[values[0]] = val
			} else {
				fmt.Println("Unknown variable")
			}
			continue
		}
		postfix, err := infixToPostfix(userInput)
		if err != nil {
			fmt.Println("Invalid expression")
			continue
		}
		rez, err := evalPostfix(postfix, variables)
		if err != nil {
			fmt.Println("Invalid expression")
		} else {
			fmt.Println(rez)
		}
	}
}

func inputClosure() func() string {
	reader := bufio.NewReader(os.Stdin)
	return func() string {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		return line
	}
}

func isOperator(token string) bool {
	return operatorPattern.MatchString(token)
}

func transformOperator(operator string) string {
	isPlus := strings.Contains(operator, "+")
	isMinus := strings.Contains(operator, "-")
	if !isPlus && !isMinus {
		return operator
	}
	if strings.Count(operator, "-")%2 == 0 {
		return "+"
	}
	return "-"
}

func infixToPostfix(infix string) (string, error) {
	var postfix []string
	var stack []string

	infix = strings.ReplaceAll(infix, "(", "( ")
	infix = strings.ReplaceAll(infix, ")", " )")
	tokens := strings.Fields(infix)

	for _, token := range tokens {
		if !isOperator(token) {
			postfix = append(postfix, token)
			continue
		}
		operator := transformOperator(token)
		if len(stack) == 0 || stack[len(stack)-1] == "(" {
			stack = append(stack, operator)
			continue
		}
		if operator == ")" {
			for stack[len(stack)-1] != "(" {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					return "", fmt.Errorf("error: Mismatched parentheses")
				}
			}
			stack = stack[:len(stack)-1]
			continue
		}
		if precedence[operator] > precedence[stack[len(stack)-1]] {
			stack = append(stack, operator)
			continue
		} else {
			for len(stack) > 0 && precedence[operator] <= precedence[stack[len(stack)-1]] && stack[len(stack)-1] != "(" {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, operator)
			continue
		}
	}

	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return strings.Join(postfix, " "), nil
}

func evalPostfix(postfix string, variables map[string]int) (int, error) {
	var stack []int
	tokens := strings.Fields(postfix)

	for _, token := range tokens {
		if isOperator(token) {
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			operand1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result, err := applyOperator(operand1, operand2, token)
			if err != nil {
				return 0, err
			}
			stack = append(stack, result)
		} else {
			// If it's an operand, check if it's a variable.
			if val, ok := variables[token]; ok {
				stack = append(stack, val)
			} else {
				// If it's not a variable, try to convert it to a float.
				value, err := strconv.Atoi(token)
				if err != nil {
					return 0, err
				}
				stack = append(stack, value)
			}
		}
	}

	if len(stack) == 1 {
		return stack[0], nil
	} else {
		return 0, fmt.Errorf("invalid postfix expression")
	}
}

func applyOperator(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	case "^":
		return int(math.Pow(float64(a), float64(b))), nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", operator)
	}
}
