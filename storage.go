package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadTask(filename string) ([]Task, error) {
	var tasks []Task

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
		fmt.Println("Ошибка при преобразовании JSON: ", err)
	}
	return tasks, err
}

func SaveTask(filename string, tasks []Task) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка при сохрании файла: ", err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		fmt.Println("Ошибка при преобразовании в JSON ", err)
	}
}
