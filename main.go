package main

import "fmt"

func main() {

	var todos TodoManager = &Todos{}

	todos.AddTodo("Fuck this bro")
	todos.AddTodo("Spongeboob")
	todos.AddTodo("Cool")
	fmt.Println(todos)
	todos.EditTodo(1, "updated todo")
	// // err := todos.DeleteTodo(3)
	// // if err != nil {
	// // fmt.Println(err.Error())
	// // }
	// err := todos.CompleteTodo(5)
	// if err != nil {
	// fmt.Println(err.Error())
	// }
	// todos.ClearTodos()
	// todos.ClearTodos()
	fmt.Println(todos)

}
