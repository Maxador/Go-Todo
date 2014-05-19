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

func TestSaveTask(t *testing.T) {
	task := newTaskOrFatal(t, "learn Go")
	m := NewTaskManager()
	err := m.Save(task)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestSaveTaskAndRetrieve(t *testing.T) {
	task := newTaskOrFatal(t, "learn Go")
	m := NewTaskManager()
	err := m.Save(task)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	all := m.All()
	if len(all) != 1 {
		t.Errorf("Expected one task, got %v", len(all))
	}
	if all[0].Title != task.Title {
		t.Errorf("Expected %v and got %v instead", task.Title, all[0].Title)
	}
}