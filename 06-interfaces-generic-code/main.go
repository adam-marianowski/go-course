package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving note failed")
		return err
	}

	fmt.Println("note saved")
	return nil
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		return
	}

	todoText := getTodoData()
	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

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
