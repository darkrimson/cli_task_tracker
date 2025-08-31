package util

import (
	"cli_task_tracker/datamodel"
	"fmt"
)

func PrintTasks(tasks []datamodel.Task) {
	for _, task := range tasks {
		fmt.Printf("%d: %s [%s]\n", task.ID, task.Description, task.Status)
	}
}
