package main

import (
	"errors"
	"fmt"
)

func main() {
	title ,err := getUserinput("Note Title:")
	if err != nil {
		fmt.Println(err)
		return
	}
	content, err := getUserInput("Note Content:")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("invalid input")
	}
	return value, nil
}