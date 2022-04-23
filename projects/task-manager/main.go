package main

import (
	"task-manager/cmd"
	"task-manager/database"
)

func main() {
	err := database.OpenDB("task.db")
	if err != nil {
		panic(err)
	}

	err = cmd.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
