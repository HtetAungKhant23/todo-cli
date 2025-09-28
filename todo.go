package main

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID
	Title       string
	Completed   bool
	CompletedAt *time.Time
	Deadline    *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Todos []Todo

func (todos *Todos) add(title string, deadline *time.Time) {
	todo := Todo{
		ID:          uuid.New(),
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		Deadline:    deadline,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}
