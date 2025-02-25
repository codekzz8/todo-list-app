package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"todo/todo"
)

type Storage struct {
	FileName string
}

func NewStorage() Storage {
	return Storage{"todoItems.json"}
}

func (storage *Storage) Save(data todo.TodoList) error {
	fileData, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		fmt.Println("Error writting todo items to json")
		return err
	}

	return os.WriteFile(storage.FileName, fileData, 0644)
}

func (storage *Storage) Load(fileData *todo.TodoList) error {
	data, err := os.ReadFile(storage.FileName)
	if err != nil {
		// file was not yet created, but it will be after writting the list to it for the first time
		// or, worst case scenario, an actual IO exception occured, in which case, unlucky
		return err
	}

	return json.Unmarshal(data, &fileData)
}
