package task

import (
	"encoding/json"
	"fmt"
	"os"
)

func CreateTaskfile(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			fmt.Println("Error creating the Task file", err)
			return
		}
		fmt.Println("Task file was created")
		return
	} else {
		return
	}
}

func taskFilePath() string {
	filePath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory", err)
		return ""
	}

	taskPath := filePath + "/tasks.json"

	CreateTaskfile(taskPath)

	return taskPath
}

func ReadTasksFromFile() ([]Task, error) {
	filePath := taskFilePath()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error Opening the file", err)
		return nil, err
	}

	defer file.Close()

	tasks := []Task{}

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		fmt.Println("Error reading json from task file", err)
		return nil, err
	}

	return tasks, nil
}

func WriteTasksToFile(tasks []Task) error {
	filePath := taskFilePath()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the task file", err)
		return err
	}

	defer file.Close()

	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		fmt.Println("Error writing tasks to file", err)
		return err
	}

	return nil
}