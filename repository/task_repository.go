package repository

import (
	"cli_task_tracker/datamodel"
	"cli_task_tracker/storage"
	"fmt"
	"time"
)

type TaskRepositoryInterface interface {
	AddTask(description string) (int, error)
	UpdateTaskDescription(ID int, description string) error
	UpdateTaskStatus(ID int, status string) error
	DeleteTask(ID int) error
	GetAllTask() ([]datamodel.Task, error)
	GetTasksByStatus(status string) ([]datamodel.Task, error)
	LoadFromFile(filename string) error
	SaveToFile(filename string) error
}

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

func (t TaskRepository) GetAllTask() ([]datamodel.Task, error) {
	if len(t.Tasks) == 0 {
		return nil, fmt.Errorf("список пуст")
	}
	return t.Tasks, nil
}

func (t TaskRepository) GetTasksByStatus(status string) ([]datamodel.Task, error) {
	var result []datamodel.Task
	for _, task := range t.Tasks {
		if task.Status == status {
			result = append(result, task)
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("нет задач со статусом %s", status)
	}
	return result, nil
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
