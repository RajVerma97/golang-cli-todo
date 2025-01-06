package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	Clear  bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new task")
	flag.IntVar(&cf.Del, "del", -1, "Delete a task")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a task")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a task")
	flag.BoolVar(&cf.Clear, "clear", false, "Clear all tasks")

	flag.Parse()
	return &cf
}

func (cf *CmdFlags) ExecuteCommand(todos *Todos) {
	switch {
	case cf.Add != "":
		todos.AddTodo(cf.Add)
	case cf.Del != -1:
		todos.DeleteTodo(cf.Del)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, " ", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid edit command")
			return
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid edit command")
			os.Exit(1)
		}
		todos.EditTodo(index, parts[1])
	case cf.Toggle != -1:
		todos.ToggleTodo(cf.Toggle)
	case cf.Clear:
		todos.ClearTodos()
	default:
		fmt.Println("Invalid command")
	}

	todos.PrintTodosTable()
}
