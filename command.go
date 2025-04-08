package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Allows you to add a new todo item.")
	flag.StringVar(&cf.Edit, "edit", "", "Allows you to edit a todo item.")
	flag.IntVar(&cf.Del, "delete", -1, "Allows you to delete a todo item.")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Allows you to toggle a todo item.")
	flag.BoolVar(&cf.List, "list", false, "Allows you to list todo list.")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List: 
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)

		if len(parts) != 2 {
			fmt.Println("Error, Invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			
			fmt.Println("Error, Invalid index for edit. Please use id:new_title")
			os.Exit(1)
		}

		todos.edit(index, parts[1])
		
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Del != -1:
		todos.delete(cf.Del)
	default: 
		fmt.Println("Invalid command")
	}
	
}
