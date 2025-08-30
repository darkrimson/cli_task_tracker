package main

import (
	"cli_task_tracker/repository"
	"cli_task_tracker/util"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	repo := &repository.TaskRepository{}
	repo.LoadFromFile("tasks.json")

	argumentsWithoutProg := os.Args[1:]
	if len(argumentsWithoutProg) == 0 {
		util.LogError(errors.New("необходимо написать команды"))
	}

	switch argument := argumentsWithoutProg[0]; argument {
	case "add":
		if len(argumentsWithoutProg) == 1 {
			util.LogError(errors.New("описание задачи не написано"))
		}
		newTask, err := repo.AddTask(argumentsWithoutProg[1])
		util.LogError(err)

		repo.SaveToFile("tasks.json")
		fmt.Printf("Задача успешно добавлена (ID: %d)\n", newTask)

	case "update":
		if len(argumentsWithoutProg) == 1 {
			util.LogError(errors.New("описание задачи не написано"))
		}
		taskID, err := strconv.Atoi(argumentsWithoutProg[1])
		util.LogError(err)
		description := argumentsWithoutProg[2]
		if len(description) == 0 {
			util.LogError(errors.New("description not provided"))
		}

		err = repo.UpdateTaskDescription(taskID, description)
		util.LogError(err)

		repo.SaveToFile("task.json")
		fmt.Printf("Задача успешно обновлена (ID: %d)\n", taskID)
	}
}
