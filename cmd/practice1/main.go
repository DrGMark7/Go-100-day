package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

func main() {
	fmt.Println("This is Calculator")

	var (
		num1   float64
		num2   float64
	)
	
	fmt.Print("Enter notation: ")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	
	word_list := strings.Split(strings.TrimSpace(line), " ")
	fmt.Println("You entered:", word_list)

	num1, err1 := strconv.ParseFloat(word_list[0], 64)
	num2, err2 := strconv.ParseFloat(word_list[2], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid numbers. Please enter valid floating-point numbers.")
		return
	}
	operator := word_list[1]

	switch operator {
	case "+":
		fmt.Printf("Result: %.2f\n", num1+num2)
	case "-":		
		fmt.Printf("Result: %.2f\n", num1-num2)
	case "*":		
		fmt.Printf("Result: %.2f\n", num1*num2)
	case "/":		
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
		return
		} else {
			fmt.Printf("Result: %.2f\n", num1/num2)
		}
	case "%":
		fmt.Printf("Result: %.2f\n", float64(int(num1)%int(num2)))
		
	default:
		fmt.Println("Invalid operator. Please use one of the following: +, -, *, /, %")
		return
	}

	fmt.Println("Calculation completed successfully.")
}