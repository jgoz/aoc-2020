package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Operator int

const (
	None Operator = iota
	Add
	Mult
)

func evaluateSimple(expr string) (value int, consumed int, err error) {
	op := None

Eval:
	for i := 0; i < len(expr); i++ {
		consumed = i + 1
		token := expr[i]
		var next int

		switch token {
		case ' ':
			continue
		case '(':
			var advance int
			next, advance, err = evaluateSimple(expr[i+1:])
			if err != nil {
				break Eval
			}
			i += advance
		case ')':
			break Eval
		case '+':
			op = Add
			continue
		case '*':
			op = Mult
			continue
		default:
			// we are lucky because the input file only uses single-character digits
			next, err = strconv.Atoi(string(token))
			if err != nil {
				break Eval
			}
		}

		switch op {
		case None:
			value = next
		case Add:
			value += next
		case Mult:
			value *= next
		}
	}
	return
}

func part1(exprs []string) (sum int, err error) {
	for _, expr := range exprs {
		value, _, err := evaluateSimple(expr)
		if err != nil {
			break
		}
		sum += value
	}
	return
}

func checkOp(token interface{}, against interface{}) (isOperator bool, higherPrecedence bool) {
	switch token {
	case "*", "+":
		isOperator = true
		if against == "+" && token == "*" {
			higherPrecedence = true
		}
	}
	return
}

func evaluateAdvanced(expr string) (value int) {
	// shunting-yard algorithm with evaluation at the end

	var stack []interface{}
	var result []interface{}

	for i := 0; i < len(expr); i++ {
		token := string(expr[i])

		switch token {
		case " ":
			continue
		case "(":
			stack = append(stack, token)
		case ")":
			var op interface{}
			for len(stack) > 0 {
				// pop item ("(" or operator) from stack
				op, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if op == "(" {
					break // discard "("
				}
				result = append(result, op) // add operator to result
			}
		case "+", "*":
			for len(stack) > 0 {
				// consider top item on stack
				top := stack[len(stack)-1]
				isOp, higherPrecedence := checkOp(top, token)
				if !isOp || higherPrecedence {
					break
				}
				// top item is an operator that needs to come off
				stack = stack[:len(stack)-1] // pop it
				result = append(result, top) // add it to result
			}
			// push operator (the new one) to stack
			stack = append(stack, token)
		default:
			num := int(expr[i] - 48)
			result = append(result, num)
		}
	}
	// drain stack to result
	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	var operands []int
	for _, token := range result {
		switch op := token.(type) {
		case int:
			operands = append(operands, op)
		case string:
			if len(operands) < 2 {
				break
			}
			switch op {
			case "+":
				operands[len(operands)-2] += operands[len(operands)-1]
				operands = operands[:len(operands)-1]
			case "*":
				operands[len(operands)-2] *= operands[len(operands)-1]
				operands = operands[:len(operands)-1]
			}
		}
	}
	return operands[0]
}

func part2(exprs []string) (sum int, err error) {
	for _, expr := range exprs {
		value := evaluateAdvanced(expr)
		sum += value
	}
	return
}

var part func([]string) (int, error)
var usePart2 bool

func init() {
	flag.BoolVar(&usePart2, "2", false, "Run part 2")
}

func main() {
	flag.Parse()
	if usePart2 {
		part = part2
	} else {
		part = part1
	}

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum, err := part(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
