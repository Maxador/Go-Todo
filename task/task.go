package task

import "fmt"

type Task struct {
	ID int64
	Title string
	Done bool
}

type TaskManager struct {
	tasks []*Task
	lastID int64
}

func NewTask(title string) (*Task, error) {
	if title == "" {
		return nil, fmt.Errorf("Empty title")
	}
	return &Task { 0, title, false }, nil
}

func NewTaskManager() *TaskManager {
	return &TaskManager {}
}

func (m *TaskManager) Save(task *Task) error {
	if task.ID == 0 {
		m.lastID++
		task.ID = m.lastID
		m.tasks = append(m.tasks, cloneTask(task))
		return nil
	}
	for i, t := range m.tasks {
		if t.ID == task.ID {
			m.tasks[i] = cloneTask(task)
			return nil
		}
	}
	return fmt.Errorf("Unknown task")
}

func cloneTask(t *Task) *Task {
	c := *t
	return &c
}

func (m *TaskManager) All() []*Task {
	return m.tasks
}

func (m *TaskManager) Find(taskID int64) (*Task, bool) {
	for _, t := range m.tasks {
		if t.ID == taskID {
			return t, true
		}
	}
	return nil, false
}

func (m *TaskManager) GetTaskIndex(taskID int64) (int, bool) {
	for i, t := range m.tasks {
		if t.ID == taskID {
			return i, true
		}
	}
	return -1, false
}

func (m *TaskManager) Delete(taskID int64) bool {
	i, ok := m.GetTaskIndex(taskID)
	if ok {
		copy(m.tasks[i:], m.tasks[i+1:])
		m.tasks[len(m.tasks)-1] = nil
		m.tasks = m.tasks[:len(m.tasks)-1]
	}
	return ok
}