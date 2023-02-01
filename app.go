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

func readLine() {
	scanner := bufio.NewScanner(os.Stdin)
	prompt := "Send a line"

	fmt.Println(prompt)
	scanner.Scan()
	in := scanner.Text()
	res := "Thanks for the input -> "

	fmt.Printf("%s%s", res, in)
	newLine()
}

func promptSelect(label string, items []string) *promptui.Select {
	prompt := &promptui.Select{
		Label: label,
		Items: items,
	}

	return prompt
}

func main() {
	fmt.Println("uoffin・ウオフィン")
	newLine()

	label := "Choose a Language"
	lang := []string{
		"JP",
		"CN",
	}

	promptLang := promptSelect(label, lang)
	_, result, err := promptLang.Run()

	if err != nil {
		fmt.Printf("prompt failed -- %v\n", err)
		return
	}

	fmt.Printf("Language Set -- %q\n", result)

	promptSel := "Enter the correct character : "
	promptSelChar := "あ" // placeholder
	prompt := promptui.Prompt{
		Label:     promptSel + promptSelChar,
	}

	result, err = prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You answered %s\n", result)
}
