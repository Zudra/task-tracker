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

func NewUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a task state in the task list",
		Long: `Update a task state providing the Task ID and the new description or state
		
		Example:

		task-tracker update 1 "buy groceries for dinner" (description)
		or
		task-tracker update 1 "in-progress" (state)
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskCmd(args)
		},
	}

	return cmd
}

func RunUpdateTaskCmd(args []string) error {
	if len(args) < 2 {
		return errors.New("update needs the Task ID and the new description or state")
	}

	TaskID := args[0]
	UpdateObject := args[1]

	TaskIDInt, err := strconv.ParseInt(TaskID, 10, 32)
	if err != nil {
		return err
	}

	switch UpdateObject {
	case "todo":
		return RunUpdateStatusCmd(TaskIDInt, "todo")
	case "in-progress":
		return RunUpdateStatusCmd(TaskIDInt, "in-progress")
	case "done":
		return RunUpdateStatusCmd(TaskIDInt, "done")
	default:
		return RunUpdateDescriptionCmd(TaskIDInt, UpdateObject)
	}
}

func RunUpdateStatusCmd(TaskID int64, status task.TaskStatus) error {
	return task.UpdateTaskStatus(TaskID, status)
}

func RunUpdateDescriptionCmd(TaskID int64, description string) error {
	return task.UpdateTaskDescription(TaskID, description)
}
