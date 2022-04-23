package main

import "task-manager/cmd"

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
