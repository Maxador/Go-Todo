package task

import "testing"

func newTaskOrFatal(t *testing.T, title string) *Task {
	task, err := NewTask(title)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	return task
}

func TestNewTask(t *testing.T) {
	task := newTaskOrFatal(t, "learn Go")
	if task.Title !=  "learn Go" {
		t.Errorf("Expected learn Go, got %v", task.Title)
	}
}

func TestNewTaskWithEmptyTitle(t *testing.T) {
	_, err := NewTask("")
	if err == nil {
		t.Errorf("Expected 'Empty title' error, got %v", err)
	}
}