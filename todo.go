package main

import (
	"errors"
	"time"
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
	CompleteTodo(int) error
	ClearTodos()
	EditTodo(int, string) error
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

func (todos *Todos) CompleteTodo(id int) error {

	completionTime := time.Now()

	for index, todo := range *todos {
		if id == todo.ID {
			(*todos)[index].CompletedAt = &completionTime
			(*todos)[index].IsCompleted = true
			return nil
		}
	}

	return errors.New("invalid id")

}

func (todos *Todos) ClearTodos() {
	(*todos) = nil
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
