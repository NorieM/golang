package main

import (
	"fmt"
	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err:=note.New(title, content)

	if err !=nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title:")	
	content:= getUserInput("Note Content:")

	return content, title
}


func getUserInput(prompt string) string {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)

	return value
}