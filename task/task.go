package task

import "errors"

type Task struct {
	Title string
}

type TaskManager struct {
	task *Task
}

func NewTask(title string) (*Task, error) {
	if title == "" {
		return nil, errors.New("Empty title")
	}
	return &Task{title}, nil
}

func NewTaskManager() *TaskManager {
	return &TaskManager {}
}

func (m *TaskManager) Save(task *Task) error {
	m.task = task
	return nil
}

func (m *TaskManager) All() []*Task {
	return []*Task {m.task}
}