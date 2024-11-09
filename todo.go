package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	// If index is invalid
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	// If index is valid
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	// Take out the item and give back a new list without it
	*todos = append(t[:index], t[index+1:]...)

	return nil
}
