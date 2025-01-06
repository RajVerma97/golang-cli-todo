package main

import (
	"fmt"
)

func main() {
	var todoManager TodoManager = &Todos{}

	storage := NewStorage[Todos]("todo.json")

	if err := storage.Load(todoManager.(*Todos)); err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}

	cmdFlags := NewCmdFlags()
	cmdFlags.ExecuteCommand(todoManager.(*Todos))

	if err := storage.Save(todoManager.(*Todos)); err != nil {
		fmt.Println("Error saving todos:", err)
	}
}
