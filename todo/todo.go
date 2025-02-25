package todo

import (
	"errors"
	"fmt"
	"slices"
	"time"
)

type Todo struct {
	Title     string
	Completed bool
	CreatedAt time.Time
}

func NewTodo(title string) Todo {
	return Todo{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now()}
}

type TodoList []Todo

func (list *TodoList) Add(title string) {
	item := NewTodo(title)
	*list = append(*list, item)
}

func (list *TodoList) Remove(index int) error {
	listPointer := *list
	err := listPointer.validateIndex(index)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	*list = slices.Delete(listPointer, index, index+1)
	return nil
}

func (list *TodoList) Edit(index int, title string) error {
	listValue := *list
	err := listValue.validateIndex(index)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	listValue[index].Title = title
	return nil
}

func (list *TodoList) List() {
	for index, item := range *list {
		fmt.Printf("- %d: %s, %t, %s\n", index, item.Title, item.Completed, item.CreatedAt.Format(time.DateTime))
	}
}

func (list *TodoList) validateIndex(index int) error {
	if index < 0 || index >= len(*list) {
		err := errors.New("index is invalid")
		return err
	}
	return nil
}
