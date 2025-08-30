package repository

import (
	"cli_task_tracker/datamodel"
	"cli_task_tracker/storage"
	"fmt"
	"time"
)

type TaskRepository struct {
	Tasks []datamodel.Task
}

func (t *TaskRepository) AddTask(taskDescription string) (int, error) {
	for _, t := range t.Tasks {
		if t.Description == taskDescription {
			return 0, fmt.Errorf("задача с таким названием %s уже существует", t.Description)
		}
	}

	task := datamodel.Task{
		ID:          len(t.Tasks) + 1,
		Description: taskDescription,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	t.Tasks = append(t.Tasks, task)
	return task.ID, nil
}

func (t *TaskRepository) UpdateTaskDescription(ID int, description string) error {
	for i := range t.Tasks {
		if t.Tasks[i].ID == ID {
			t.Tasks[i].Description = description
			t.Tasks[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("не удалось изменить название задачи")
}

func (t *TaskRepository) UpdateTaskStatus(ID int, status string) error {
	for i := range t.Tasks {
		if t.Tasks[i].ID == ID {
			t.Tasks[i].Status = status
			t.Tasks[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("не удалось изменить статус задачи")
}

func (t *TaskRepository) DeleteTask(ID int) error {
	for i := range t.Tasks {
		if t.Tasks[i].ID == ID {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with id %d not found", ID)
}

func (t TaskRepository) GetAllTask() {
	for _, task := range t.Tasks {
		fmt.Printf("%d: %s [%s]\n", task.ID, task.Description, task.Status)
	}
}

func (t TaskRepository) CompletedTask() {
	for _, task := range t.Tasks {
		if task.Status == "done" {
			fmt.Printf("%d: %s [%s]\n", task.ID, task.Description, task.Status)
		}
	}
}

func (t TaskRepository) InprogressTask() {
	for _, task := range t.Tasks {
		if task.Status == "in progress" {
			fmt.Printf("%d: %s [%s]\n", task.ID, task.Description, task.Status)
		}
	}
}

func (t TaskRepository) TodoTask() {
	for _, task := range t.Tasks {
		if task.Status == "todo" {
			fmt.Printf("%d: %s [%s]\n", task.ID, task.Description, task.Status)
		}
	}
}

func (t *TaskRepository) LoadFromFile(filename string) error {
	tasks, err := storage.LoadTask(filename)
	if err != nil {
		return err
	}
	t.Tasks = tasks
	return nil
}

func (t *TaskRepository) SaveToFile(filename string) error {
	return storage.SaveTask(filename, t.Tasks)
}
