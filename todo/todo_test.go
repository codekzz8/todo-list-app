package todo_test

import (
	"testing"
	"todo/todo"
)

func TestAddShouldAddItemToList(t *testing.T) {
	actualList := todo.TodoList{}
	expectedList := mockTodoList()

	actualList.Add("mockTitle")

	if len(actualList) != len(expectedList) {
		t.Errorf("Expected list length to be \"%d\" but was \"%d\"", len(expectedList), len(actualList))
	}

	actualItem := actualList[0]
	expectedItem := expectedList[0]
	assertField(t, "Title", expectedItem.Title, actualItem.Title)
}

func TestEditShouldEditItemTitle(t *testing.T) {
	todos := mockTodoList()
	expectedTitle := "new title"
	todos.Add("second item")

	err := todos.Edit(0, "new title")

	if err != nil {
		t.Errorf("Error was thrown but was not expected\nMessage: %s", err.Error())
	}
	assertField(t, "Title", expectedTitle, todos[0].Title)
}

func TestEditWithNegativeIndexShouldReturnError(t *testing.T) {
	todos := mockTodoList()
	expectedErrorMessage := "index is invalid"
	todos.Add("second item")

	err := todos.Edit(-1, "new title")
	assertField(t, "ErrorMessage", expectedErrorMessage, err.Error())
}

func TestEditWithIndexOutOfBoundsShouldReturnError(t *testing.T) {
	todos := mockTodoList()
	expectedErrorMessage := "index is invalid"
	todos.Add("second item")

	err := todos.Edit(3, "new title")
	assertField(t, "ErrorMessage", expectedErrorMessage, err.Error())
}

func TestRemoveShouldRemoveItem(t *testing.T) {
	todos := mockTodoList()

	err := todos.Remove(0)
	if err != nil {
		t.Errorf("Error was cast but was not expected\nMessage: %s", err.Error())
	}
	if len(todos) > 0 {
		t.Errorf("Expected length of the list to be 0 but it was %d", len(todos))
	}
}

func TestRemoveWithNegativeIndexShouldReturnError(t *testing.T) {
	todos := mockTodoList()
	expectedErrorMessage := "index is invalid"
	todos.Add("second item")

	err := todos.Remove(-1)
	assertField(t, "ErrorMessage", expectedErrorMessage, err.Error())
}

func TestRemoveWithIndexOutOfBoundsShouldReturnError(t *testing.T) {
	todos := mockTodoList()
	expectedErrorMessage := "index is invalid"
	todos.Add("second item")

	err := todos.Remove(3)
	assertField(t, "ErrorMessage", expectedErrorMessage, err.Error())
}

func assertField(t *testing.T, fieldName string, expected string, actual string) {
	if expected != actual {
		t.Errorf("Expected %s to be \"%s\" but was \"%s\"", fieldName, expected, actual)
	}
}

func mockTodoList() todo.TodoList {
	return todo.TodoList{mockTodoItem("mockTitle")}
}

func mockTodoItem(title string) todo.Todo {
	return todo.NewTodo(title)
}
