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
	initializeFile(storage.FileName)

	data, err := os.ReadFile(storage.FileName)
	if err != nil {
		fmt.Println("Error reading todo items file")
		return err
	}

	return json.Unmarshal(data, &fileData)
}

func initializeFile(fileName string) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// todoItems.json does not exist, create it
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error creating todo list file")
		}
		defer file.Close()
	}
}
