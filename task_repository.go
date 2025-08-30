package main

import (
	"fmt"
	"time"
)

type TaskRepository struct {
	tasks []Task
}

func (t *TaskRepository) AddTask(taskDescription string) error {
	for _, t := range t.tasks {
		if t.Description == taskDescription {
			return fmt.Errorf("Задача с таким названием %s уже существует", t.Description)
		}
	}

	task := Task{
		ID:          len(t.tasks) + 1,
		Description: taskDescription,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	t.tasks = append(t.tasks, task)
	return nil
}

func (t *TaskRepository) UpdateTaskDescription(ID int, description string) error {
	for i := range t.tasks {
		if t.tasks[i].ID == ID {
			t.tasks[i].Description = description
			t.tasks[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("Не удалось изменить название задачи")
}

func (t *TaskRepository) UpdateTaskStatus(ID int, status string) error {
	for i := range t.tasks {
		if t.tasks[i].ID == ID {
			t.tasks[i].Status = status
			t.tasks[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("Не удалось изменить статус задачи")
}

func (t *TaskRepository) DeleteTask(ID int) error {
	for i := range t.tasks {
		if t.tasks[i].ID == ID {
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with id %d not found", ID)
}

func (t TaskRepository) GetAllTask() {
	for _, task := range t.tasks {
		fmt.Printf("%d: %s [%s]\n", task.ID, task.Description, task.Status)
	}
}

func (t TaskRepository) CompletedTask() {
	for _, task := range t.tasks {
		if task.Status == "done" {
			fmt.Printf("%d: %s [%s]\n", task.ID, task.Description, task.Status)
		}
	}
}

func (t TaskRepository) InprogressTask() {
	for _, task := range t.tasks {
		if task.Status == "in progress" {
			fmt.Printf("%d: %s [%s]\n", task.ID, task.Description, task.Status)
		}
	}
}

func (t TaskRepository) TodoTask() {
	for _, task := range t.tasks {
		if task.Status == "todo" {
			fmt.Printf("%d: %s [%s]\n", task.ID, task.Description, task.Status)
		}
	}
}
