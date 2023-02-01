package main

import (
	"bufio"
	"fmt"
	"os"
)

func newLine() {
	fmt.Print("\n")
}

func main() {
	fmt.Println("uoffin シ")
	newLine()

	scanner := bufio.NewScanner(os.Stdin)

	var prompt string = "Send a line"
	fmt.Println(prompt)

	scanner.Scan()
	text := scanner.Text()
	res := "Thanks for the input -> "
	fmt.Printf("%s%s", res, text)
	newLine()
}

// Don't forget to run command --> `go mod tidy`
