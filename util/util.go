package util

import (
	"cli_task_tracker/datamodel"
	"fmt"
	"log"
)

func LogError(err error) {
	if err != nil {
		log.Fatal("cli_task_tracker: error ", err)
	}
}

func PrintTasks(tasks []datamodel.Task) {
	for _, task := range tasks {
		fmt.Printf("%d: %s [%s]\n", task.ID, task.Description, task.Status)
	}
}

func PrintHelp() {
	fmt.Println("Task CLI Tracker")
	fmt.Println()
	fmt.Println("Using:")
	fmt.Println("  task-cli add \"description\"          - add a new task")
	fmt.Println("  task-cli update ID \"description\"    - change the task description with ID")
	fmt.Println("  task-cli delete ID                 	 - delete the task with ID")
	fmt.Println("  task-cli mark-in-progress ID      	 - mark the task as in progress")
	fmt.Println("  task-cli mark-done ID             	 - mark the task as completed")
	fmt.Println("  task-cli list [todo|done|in-progress] - show tasks, you can filter by status")
	fmt.Println("  task-cli help                         - show this help")
	fmt.Println()
}
