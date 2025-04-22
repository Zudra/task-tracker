/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the tasks registered in your tracker",
		Long: `List all the tasks based on their status: 

		task-tracker list (list all)
		task-tracker list todo (list all todo tasks)
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListCmd(args)
		},
	}

	return cmd
}

func RunListCmd(args []string) error {
	if len(args) > 0 {
		status := task.TaskStatus(args[0])
		return task.ListTasks(status)
	}

	return task.ListTasks("all")
}

func init() {
	rootCmd.AddCommand(NewListCmd())
}