package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving note failed")
		return
	}

	fmt.Println("note saved")

	todoText := getTodoData()
	todo, err := todo.New(todoText)

}

func getNoteData() (string, string) {
	title := getUserInput("note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getTodoData() string {
	text := getUserInput("Todo text: ")
	return text
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	// remove line break
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
