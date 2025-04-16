package task

import (
	"fmt"
	"time"
)

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "todo"
	TASK_STATUS_IN_PROGRESS TaskStatus = "in-progress"
	TASK_STATUS_DONE        TaskStatus = "done"
)

type Task struct {
	ID          int64      `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func NewTask(id int64, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      TASK_STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ListTasks(status TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	var filteredTasks []Task
	switch status {
	case "all":
		filteredTasks = tasks
	case TASK_STATUS_TODO:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_TODO {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case TASK_STATUS_IN_PROGRESS:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_IN_PROGRESS {
				filteredTasks = append(filteredTasks, task)
			}
		}
	case TASK_STATUS_DONE:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_DONE {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}

	for _, task := range filteredTasks {
		fmt.Println()
		fmt.Printf("ID: %v,\n Status: %s,\n Description: %s,\n CreatedAt: %s, \nUpdatedAt: %s\n", task.ID, task.Status, task.Description, task.CreatedAt, task.UpdatedAt)
		fmt.Println()
	}

	return nil
}

func FormatStatusFromString(status string) TaskStatus {
	switch status {
	case "all":
		return "all"
	case "to-do":
		return TASK_STATUS_TODO
	case "in-progress":
		return TASK_STATUS_IN_PROGRESS
	case "done":
		return TASK_STATUS_DONE
	default:
		return "all"
	}
}

func AddTask(description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var maxID int64 = 0
	var nextTaskID int64
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	nextTaskID = maxID + 1

	task := NewTask(nextTaskID, description)

	tasks = append(tasks, *task)
	fmt.Printf("New task added ID: %v\n", task.ID)

	return WriteTasksToFile(tasks)
}

func DeleteTask(id int64) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID != id {
			updatedTasks = append(updatedTasks, task)
		}
	}

	if len(updatedTasks) == len(tasks) {
		return fmt.Errorf("the task was not found (ID: %d)", id)
	}

	fmt.Printf("The task was successfully deleted ID: %d\n\n", id)
	return WriteTasksToFile(updatedTasks)
}

func UpdateTaskStatus(id int64, status TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var taskExists bool = false
	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID == id {
			taskExists = true
			switch status {
			case TASK_STATUS_TODO:
				task.Status = TASK_STATUS_TODO
			case TASK_STATUS_IN_PROGRESS:
				task.Status = TASK_STATUS_IN_PROGRESS
			case TASK_STATUS_DONE:
				task.Status = TASK_STATUS_DONE
			}
			task.UpdatedAt = time.Now()

			updatedTasks = append(updatedTasks, task)
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !taskExists {
		return fmt.Errorf("the task given was not found (ID: %d)", id)
	}

	fmt.Printf("The task was succsessfully updated! (ID: %d)\n\n", id)

	return WriteTasksToFile(updatedTasks)
}
