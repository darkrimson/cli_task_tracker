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
	const (
		StatusTodo       = "todo"
		StatusInProgress = "in-progress"
		StatusDone       = "done"
	)

	var repo repository.TaskRepositoryInterface = &repository.TaskRepository{}

	defer func() {
		err := repo.SaveToFile("tasks.json")
		util.LogError(err)
	}()

	err := repo.LoadFromFile("tasks.json")
	util.LogError(err)

	argumentsWithoutProg := os.Args[1:]
	if len(argumentsWithoutProg) == 0 {
		util.LogError(errors.New("it is necessary to write commands"))
	}

	switch argument := argumentsWithoutProg[0]; argument {
	case "add":
		if len(argumentsWithoutProg) == 1 {
			util.LogError(errors.New("the task description is not written"))
		}
		newTask, err := repo.AddTask(argumentsWithoutProg[1])
		util.LogError(err)

		fmt.Printf("Task added successfully (ID: %d)\n", newTask)

	case "update":
		if len(argumentsWithoutProg) == 1 {
			util.LogError(errors.New("the task description is not written"))
		}
		taskID, err := strconv.Atoi(argumentsWithoutProg[1])
		util.LogError(err)
		description := argumentsWithoutProg[2]
		if len(description) == 0 {
			util.LogError(errors.New("incorrect description"))
		}

		err = repo.UpdateTaskDescription(taskID, description)
		util.LogError(err)

		fmt.Printf("The task has been successfully updated (ID: %d)\n", taskID)

	case "delete":
		if len(argumentsWithoutProg) == 1 {
			util.LogError(errors.New("the task description is not written"))
		}
		taskID, err := strconv.Atoi(argumentsWithoutProg[1])
		util.LogError(err)

		err = repo.DeleteTask(taskID)
		util.LogError(err)

		fmt.Printf("The task was successfully deleted (ID: %d)\n", taskID)

	case "mark-in-progress":
		if len(argumentsWithoutProg) == 1 {
			util.LogError(errors.New("there is no such status"))
		}
		taskID, err := strconv.Atoi(argumentsWithoutProg[1])
		util.LogError(err)

		err = repo.UpdateTaskStatus(taskID, "in-progress")
		util.LogError(err)

	case "mark-done":
		if len(argumentsWithoutProg) == 1 {
			util.LogError(errors.New("there is no such status"))
		}
		taskID, err := strconv.Atoi(argumentsWithoutProg[1])
		util.LogError(err)

		err = repo.UpdateTaskStatus(taskID, "done")
		util.LogError(err)

	case "list":
		if len(argumentsWithoutProg) > 1 {
			switch argumentsWithoutProg[1] {
			case StatusDone:
				tasks, err := repo.GetTasksByStatus(StatusDone)
				util.LogError(err)
				util.PrintTasks(tasks)
			case StatusTodo:
				tasks, err := repo.GetTasksByStatus(StatusTodo)
				util.LogError(err)
				util.PrintTasks(tasks)
			case StatusInProgress:
				tasks, err := repo.GetTasksByStatus(StatusInProgress)
				util.LogError(err)
				util.PrintTasks(tasks)
			default:
				util.LogError(errors.New("unknown filter"))
			}
		} else {
			tasks, err := repo.GetAllTask()
			util.LogError(err)
			util.PrintTasks(tasks)
		}
	default:
		fmt.Println("unknown command:", argument)
		util.PrintHelp()
	}
}
