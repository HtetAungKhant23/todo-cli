package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello from TODO CLI!")
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.load(&todos)

	deadline := time.Now().AddDate(0, 0, 2)
	todos.add("Learn how http work deeply.", &deadline)
	todos.add("Learn GoLang.", nil)
	storage.save(todos)

	todos.list()
}
