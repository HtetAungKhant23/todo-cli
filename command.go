package main

import "flag"

type CmdFlags struct {
	Add  string
	List bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "add title of new task")
	flag.BoolVar(&cf.List, "list", false, "List all tasks")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.list()
	case cf.Add != "":
		todos.add(cf.Add, nil)
	default:
		println("Invalid command!")
	}
}
