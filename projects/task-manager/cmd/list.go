package cmd

import (
	"fmt"
	"os"
	"task-manager/database"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks in your task list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := database.ReadAllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete!")
			return
		}
		fmt.Println("You have following tasks: ")
		for i, v := range tasks {
			fmt.Printf("%d. %s\n", i+1, v.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
