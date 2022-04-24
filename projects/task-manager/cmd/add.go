package cmd

import (
	"fmt"
	"strings"
	"task-manager/database"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := database.CreateTask(task)
		if err != nil {
			fmt.Println("Something went error:", err)
		}
		fmt.Printf(`Added "%d. %s" to the task list`, id, task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
