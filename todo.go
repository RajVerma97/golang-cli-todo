package main

import (
	"errors"
	"fmt"
	"os"
	"time"
	"github.com/olekukonko/tablewriter"
)

type Todo struct {
	ID          int
	Title       string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type TodoManager interface {
	AddTodo(string)
	DeleteTodo(int) error
	ToggleTodo(int) error
	ClearTodos()
	EditTodo(int, string) error
	PrintTodosTable()
}

type Todos []Todo

func (todos *Todos) AddTodo(title string) {

	todo := Todo{
		ID:          len(*todos) + 1,
		Title:       title,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)

}

func (todos *Todos) DeleteTodo(id int) error {
	t := *todos

	for index, todo := range t {
		if id == todo.ID {
			*todos = append(t[:index], t[index+1:]...)
			return nil
		}
	}

	return errors.New("invalid id")

}

func (todos *Todos) ToggleTodo(id int) error {

	completionTime := time.Now()

	for index, todo := range *todos {
		if id == todo.ID {
			(*todos)[index].CompletedAt = &completionTime
			(*todos)[index].IsCompleted = !(*todos)[index].IsCompleted
			return nil
		}
	}

	return errors.New("invalid id")

}

func (todos *Todos) ClearTodos() {
	(*todos) = make(Todos, 0)
}

func (todos *Todos) EditTodo(id int, title string) error {

	for index, todo := range *todos {
		if todo.ID == id {
			(*todos)[index].Title = title
			return nil
		}
	}

	return errors.New("invalid id")
}
func (todos *Todos) PrintTodosTable() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Is Completed", "Created At", "Completed At"})

	for _, todo := range *todos {
		var completedAtStr string = ""
		if todo.CompletedAt != nil {
			completedAtStr = todo.CompletedAt.Format("Mon, 02 May 2001 15:04:05")
		}

		createdAtStr := todo.CreatedAt.Format("Mon, 02 May 2001 15:04:05")
		var isCompletedStr string = "❌"

		if todo.IsCompleted {
			isCompletedStr = "✅"
		}

		table.Append([]string{
			fmt.Sprintf("%d", todo.ID),
			todo.Title,
			isCompletedStr,
			createdAtStr,
			completedAtStr,
		})
	}

	table.Render()
}
