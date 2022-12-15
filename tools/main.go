package main

import "fmt"

func main() {
	PrintHello()
	for i := 0; i < 5; i++ {
		PrintNumber(i)
	}
}

func PrintHello() {
	fmt.Println("Hello, Go")
}

// PrintNumber writes a number using the fmt.PrintLn function
func PrintNumber(number int) {
	fmt.Println(number)
}
