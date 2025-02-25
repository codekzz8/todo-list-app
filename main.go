package main

import (
	"todo/commands"
	"todo/storage"
	"todo/todo"
)

func main() {
	todoList := todo.TodoList{}
	storage := storage.NewStorage()
	storage.Load(&todoList)
	cmds := commands.NewCommands()
	cmds.Execute(&todoList)
	storage.Save(todoList)
}
