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
	if task.Done {
		t.Errorf("New task is done")
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

func TestSaveAndRetrieveTwoTasks(t *testing.T) {
	learnGo := newTaskOrFatal(t, "learn Go")
	learnTDD := newTaskOrFatal(t, "learn TDD")

	m := NewTaskManager()
	m.Save(learnGo)
	m.Save(learnTDD)

	all := m.All()
	if len(all) != 2 {
		t.Errorf("Expected two tasks, got %v", len(all))
	}
	if *all[0] != *learnGo && *all[1] != *learnGo {
		t.Errorf("Missing task : %v", learnGo)
	}
	if *all[0] != *learnTDD && *all[1] != *learnTDD {
		t.Errorf("Missing task : %v", learnTDD)
	}
}

func TestSaveModifyAndRetrieve(t *testing.T) {
	task := newTaskOrFatal(t, "learn Go")

	m := NewTaskManager()
	m.Save(task)
	task.Done = true
	if m.All()[0].Done {
		t.Errorf("Saved task wasn't done")
	}
}

func TestSaveTwiceAndRetrieve(t *testing.T) {
	task := newTaskOrFatal(t, "learn Go")

	m:= NewTaskManager()
	m.Save(task)
	m.Save(task)
	all := m.All()
	if len(all) != 1 {
		t.Errorf("Expected 1 task, instead got %v", len(all))
	}
	if *all[0] != *task {
		t.Errorf("Expected task %v, instead got task %v", task, all[0])
	}
}

func TestSaveAndFind(t *testing.T) {
	task := newTaskOrFatal(t, "learn Go")

	m := NewTaskManager()
	m.Save(task)

	foundTask, ok := m.Find(task.ID)
	if !ok {
		t.Errorf("Task not found")
	}
	if *task != *foundTask {
		t.Errorf("Expected to find task %v, instead got task %v", task, foundTask)
	}
}

func TestFindAndDeleteTask(t *testing.T) {
	task := newTaskOrFatal(t, "learn Go")

	m := NewTaskManager()
	m.Save(task)

	ok := m.Delete(task.ID)
	if !ok {
		t.Errorf("Task not deleted")
	}
	
}