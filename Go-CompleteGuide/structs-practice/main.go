package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(title)
	fmt.Println(content)
}

func getNoteData() (string, string, error) {
	title ,err := getUserInput("Note Title:")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	content, err := getUserInput("Note Content:")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return content, title, nil
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