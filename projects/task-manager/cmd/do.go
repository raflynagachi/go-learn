package cmd

import (
	"fmt"
	"strconv"
	"task-manager/database"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark specific task in your task list as completed",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Failed to parse the argument: %s", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := database.ReadAllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}

		for _, id := range ids {
			if id < 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", err)
				continue
			}
			task := tasks[id-1]
			err = database.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
