package commands

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo/todo"
)

type Commands struct {
	Add    string
	Remove int
	Edit   string
	List   bool
}

func NewCommands() *Commands {
	cmds := Commands{}

	flag.StringVar(&cmds.Add, "add", "", "Add a new todo item with specified title")
	flag.IntVar(&cmds.Remove, "remove", 0, "Remove todo item from list at specified index (default = 0)")
	flag.StringVar(&cmds.Edit, "edit", "", "Edit a todo item at specified index follow by the new title (index:title)")
	flag.BoolVar(&cmds.List, "list", false, "List all the todo items")

	flag.Parse()
	return &cmds
}

func (cmd *Commands) Execute(todoList *todo.TodoList) {
	switch {
	case cmd.Add != "":
		todoList.Add(cmd.Add)
	case cmd.Edit != "":
		args := strings.SplitN(cmd.Edit, ":", 2)
		if len(args) != 2 {
			fmt.Println("Invalid format. Please use index:title")
		}

		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid index. Please provide an integer for the item index")
			os.Exit(1)
		}

		todoList.Edit(index, args[1])
	case cmd.List:
		todoList.List()
	case cmd.Remove >= 0:
		todoList.Remove(cmd.Remove)
	default:
		fmt.Println("Invalid command")
	}
}
