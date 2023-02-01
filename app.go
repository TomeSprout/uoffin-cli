package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func newLine() {
	fmt.Print("\n")
}

func promptSelect(label string, items []string) promptui.Select {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	return prompt
}

func main() {
	fmt.Println("uoffin シ")
	newLine()

	scanner := bufio.NewScanner(os.Stdin)

	var customPrompt string = "Send a line"
	fmt.Println(customPrompt)

	scanner.Scan()
	text := scanner.Text()
	res := "Thanks for the input -> "
	fmt.Printf("%s%s", res, text)
	newLine()

	label := "Choose something"
	items := []string{
		"JP",
		"CN",
	}

	prompt := promptSelect(label, items)
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("prompt failed -- %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)

}
