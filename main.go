package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.load(&todos)
	cmdFlg := NewCmdFlags()
	cmdFlg.Execute(&todos)

	// deadline := time.Now().AddDate(0, 0, 2)
	// todos.add("Learn how http work deeply.", &deadline)
	// todos.add("Learn GoLang.", nil)

	storage.save(todos)
}
