package storage

import (
	"cli_task_tracker/datamodel"
	"cli_task_tracker/util"
	"encoding/json"
	"os"
)

func LoadTask(filename string) ([]datamodel.Task, error) {
	var tasks []datamodel.Task

	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		util.LogError(err)
	}
	return tasks, nil
}

func SaveTask(filename string, tasks []datamodel.Task) error {
	file, err := os.Create(filename)
	if err != nil {
		util.LogError(err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		util.LogError(err)
	}
	return nil
}
