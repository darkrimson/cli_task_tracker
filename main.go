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

		err = repo.SaveToFile("tasks.json")
		util.LogError(err)
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

		err = repo.SaveToFile("tasks.json")
		util.LogError(err)
		fmt.Printf("Задача успешно обновлена (ID: %d)\n", taskID)

	case "delete":
		if len(argumentsWithoutProg) == 1 {
			util.LogError(errors.New("описание задачи не написано"))
		}
		taskID, err := strconv.Atoi(argumentsWithoutProg[1])
		util.LogError(err)

		err = repo.DeleteTask(taskID)
		util.LogError(err)

		err = repo.SaveToFile("tasks.json")
		util.LogError(err)
		fmt.Printf("Задача успешно удалена (ID: %d)\n", taskID)

	case "mark-in-progress":
		if len(argumentsWithoutProg) == 1 {
			util.LogError(errors.New("task id not provided"))
		}
		taskID, err := strconv.Atoi(argumentsWithoutProg[1])
		util.LogError(err)

		err = repo.UpdateTaskStatus(taskID, "in-progress")
		util.LogError(err)

		err = repo.SaveToFile("tasks.json")
		util.LogError(err)

	case "mark-done":
		if len(argumentsWithoutProg) == 1 {
			util.LogError(errors.New("task id not provided"))
		}
		taskID, err := strconv.Atoi(argumentsWithoutProg[1])
		util.LogError(err)

		err = repo.UpdateTaskStatus(taskID, "done")
		util.LogError(err)

		err = repo.SaveToFile("tasks.json")
		util.LogError(err)

	case "list":
		if len(argumentsWithoutProg) > 1 {
			switch argumentsWithoutProg[1] {
			case "done":
				err := repo.CompletedTask()
				util.LogError(err)
			case "todo":
				err := repo.TodoTask()
				util.LogError(err)
			case "in-progress":
				err := repo.InprogressTask()
				util.LogError(err)
			default:
				util.LogError(errors.New("неизвестный фильтр"))
			}
		} else {
			err := repo.GetAllTask()
			util.LogError(err)
		}
	default:
		util.LogError(errors.New("option provided not valid"))
	}
}
