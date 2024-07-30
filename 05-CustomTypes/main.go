package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"customTypes/note"
	"customTypes/todo"
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

// better create an embedded interface instead of:
// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	title, content := getNodeData()

	todoText := getUserInput("Todo text:")
	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

// func saveData(data ) {}  a single function that could receive different types of data, we can do that using interfaces
func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("saving failed")
		return err
	}

	fmt.Println("saving succeeded")
	return nil
}

func getNodeData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
