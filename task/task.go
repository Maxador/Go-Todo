package task

import "errors"

type Task struct {
	Title string
}

func NewTask(title string) (*Task, error) {
	if title == "" {
		return nil, errors.New("Empty title")
	}
	return &Task{title}, nil
}