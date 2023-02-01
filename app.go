package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"

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

func setTemplate() *promptui.PromptTemplates {
	template := &promptui.PromptTemplates{
		Prompt:          "{{ . }} ",
		ValidationError: "{{ . | italic }}",
		Valid:           "{{ . | green }} ",
		Invalid:         "{{ . | red }} ",
		Success:         "{{ . | blue | underline }} ",
	}

	return template
}

func promptSelect(label string, items []string) *promptui.Select {
	prompt := &promptui.Select{
		Label: label,
		Items: items,
	}

	return prompt
}

func validateCJK(input string) error {
	// const cjk = "" +
	// 	"\u2e80-\u2eff" +
	// 	"\u2f00-\u2fdf" +
	// 	"\u3040-\u309f" +
	// 	"\u30a0-\u30ff" +
	// 	"\u3100-\u312f" +
	// 	"\u3200-\u32ff" +
	// 	"\u3400-\u4dbf" +
	// 	"\u4e00-\u9fff" +
	// 	"\uf900-\ufaff"

	// ! Not properly compiling CJK REGEXP
	// r, _ := regexp.Compile(cjk)
	m, _ := regexp.MatchString("あ", input) // placeholder

	if !m {
		return errors.New("invalid character")
	}
	return nil
}

func main() {
	fmt.Println("uoffin・ウオフィン")
	newLine()

	label := "Choose a Language"
	lang := []string{
		"JP",
		"CN",
	}

	template := setTemplate()
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
		Templates: template,
		Validate:  validateCJK,
	}

	result, err = prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You answered %s\n", result)
}
