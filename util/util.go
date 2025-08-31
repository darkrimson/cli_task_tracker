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
	fmt.Println("Использование:")
	fmt.Println("  task-cli add \"описание\"           - добавить новую задачу")
	fmt.Println("  task-cli update ID \"описание\"     - изменить описание задачи с ID")
	fmt.Println("  task-cli delete ID                 - удалить задачу с ID")
	fmt.Println("  task-cli mark-in-progress ID       - отметить задачу как выполняемую")
	fmt.Println("  task-cli mark-done ID              - отметить задачу как выполненную")
	fmt.Println("  task-cli list [todo|done|in-progress] - показать задачи, можно фильтровать по статусу")
	fmt.Println("  task-cli help                      - показать эту справку")
	fmt.Println()
}
