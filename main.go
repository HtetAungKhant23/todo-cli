package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello from TODO CLI!")
	todos := Todos{}

	deadline := time.Now().AddDate(0, 0, 1)
	todos.add("Learn how http work deeply.", &deadline)
	todos.add("Learn GoLang.", nil)

	fmt.Printf("%+v\n", todos)

}
