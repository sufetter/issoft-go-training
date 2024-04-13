package tasker

import "time"

type Task struct {
	createdAt   time.Time
	Description string
}

type Tasks struct {
	tasks []Task
}

func (t *Task) String() string {
	return t.Description
}

func (t *Tasks) Add(description string) {
	t.tasks = append(t.tasks, Task{
		createdAt:   time.Now(),
		Description: description,
	})
}

func (t *Tasks) Remove(position int) {
	t.tasks = append(t.tasks[:position], t.tasks[position+1:]...)
}

func (t *Tasks) List() *[]Task {
	return &t.tasks
}
