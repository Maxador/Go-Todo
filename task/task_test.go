package task

import "testing"

func TestNewTask( t *testing.T) {
	task := NewTask("learn Go")
	if task.Title !=  "learn Go" {
		t.Errorf("Expected learn Go, got %v", task.Title)
	}
}