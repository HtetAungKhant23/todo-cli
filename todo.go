package main

import (
	"os"
	"time"

	"github.com/aquasecurity/table"
	"github.com/google/uuid"
	"github.com/mergestat/timediff"
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

func (todos *Todos) list() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("ID", "Title", "Deadline", "Completed", "Completed At", "Created At")

	for _, todo := range *todos {
		completed := "❎"
		completedAt := ""
		deadline := ""

		if todo.Completed {
			completed = "✅"
			completedAt = timediff.TimeDiff(*todo.CompletedAt)
		}

		if todo.Deadline != nil {
			deadline = timediff.TimeDiff(*todo.Deadline)
		}

		table.AddRow(todo.ID.String(), todo.Title, deadline, completed, completedAt, timediff.TimeDiff(todo.UpdatedAt))
	}

	table.Render()
}
