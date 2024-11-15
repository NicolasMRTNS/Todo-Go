package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

func (storage *Storage[T]) Save(data T) error {
	// 4 empty spaces for formatting the json file
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	// Permission that everyone can read the file and owner can write
	return os.WriteFile(storage.FileName, fileData, 0644)
}

func (storage *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(storage.FileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}
