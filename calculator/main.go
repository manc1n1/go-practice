package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Go calculator")

	cmd := readLine("Enter command [a]dd, [s]ubtract, [m]ultiply, or [d]ivide: ")
	fmt.Println(cmd)

	switch cmd {
	case "a":
		num1, num2 := getNums()
		result := num1 + num2
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
	case "s":
		num1, num2 := getNums()
		result := num1 - num2
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
	case "m":
		num1, num2 := getNums()
		result := num1 * num2
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
	case "d":
		num1, num2 := getNums()
		result := float32(num1) / float32(num2)
		sResult := fmt.Sprintf("%f", result)
		fmt.Println(sResult)
	default:
		fmt.Println("Please only enter [a]dd, [s]ubtract, [m]ultiply, or [d]ivide")
	}
}

func readLine(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func getNums() (int, int) {
	num1String := readLine("First Number: ")
	num1, err := strconv.Atoi(num1String)
	if err != nil {
		fmt.Println(err)
	}

	num2String := readLine("Second Number: ")
	num2, err := strconv.Atoi(num2String)
	if err != nil {
		fmt.Println(err)
	}

	return num1, num2
}
