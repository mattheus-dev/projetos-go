package services

import (
	"taskmanager/entities"
	"taskmanager/utils"
)

type TaskService struct {
	Tasks []entities.Task
}

// Cria uma nova tarefa
func (s *TaskService) CreateTask(title string) {
	id := utils.GenerateID()
	task := entities.Task{ID: id, Title: title, Completed: false}
	s.Tasks = append(s.Tasks, task)
}

// Lista todas as tarefas
func (s *TaskService) ListTasks() []entities.Task {
	return s.Tasks
}

// Marca uma tarefa como concluída (por ID)
func (s *TaskService) CompleteTask(id string) bool {
	for i, t := range s.Tasks {
		if t.ID == id {
			s.Tasks[i].Completed = true
			return true
		}
	}
	return false
}

// Remove uma tarefa (por ID)
func (s *TaskService) RemoveTask(id string) bool {
	for i, t := range s.Tasks {
		if t.ID == id {
			s.Tasks = append(s.Tasks[:i], s.Tasks[i+1:]...)
			return true
		}
	}
	return false
}