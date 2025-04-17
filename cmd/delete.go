/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"strconv"
	"task-tracker/internal/task"

	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete",
		Short: "Deletes a task from the task-tracker",
		Long: `Delete the task from the task-tracker when provided the following arguments:
				
				Task ID - Id from the task you want to delete`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDeleteTaskCmd(args)
		},
	}

	return cmd
}

func RunDeleteTaskCmd(args []string) error {
	if len(args) != 1 {
		return errors.New("A Task ID is required")
	}

	TaskID := args[0]
	
	TaskIdInt, err := strconv.ParseInt(TaskID, 10, 32)
	if err != nil {
		return err
	}

	return task.DeleteTask(TaskIdInt)
}
