package todo

import (
	"errors"
	"fmt"
	"slices"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTodo(title string) Todo {
	return Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}

type TodoList []Todo

func (list *TodoList) Add(title string) {
	item := NewTodo(title)
	*list = append(*list, item)
}

func (list *TodoList) Remove(index int) {
	listPointer := *list
	err := listPointer.validateIndex(index)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	*list = slices.Delete(listPointer, index, index+1)
}

func (list *TodoList) Edit(index int, title string) {
	listValue := *list
	err := listValue.validateIndex(index)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	listValue[index].Title = title
}

func (list *TodoList) List() {
	for index, item := range *list {
		completedAt := ""
		if item.CompletedAt != nil {
			completedAt = item.CompletedAt.Format(time.DateTime)
		}
		fmt.Printf("- %d: %s, %t, %s, %s\n", index, item.Title, item.Completed, item.CreatedAt.Format(time.DateTime), completedAt)
	}
}

func (list *TodoList) validateIndex(index int) error {
	if index < 0 || index >= len(*list) {
		err := errors.New("index is invalid")
		return err
	}
	return nil
}
