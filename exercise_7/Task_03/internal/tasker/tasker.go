package tasker

import (
	"fmt"
	"io"
	"log"
	"os"
	"todo/pkg/parser"
)

type Task struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var (
	Tasks []Task
)

func init() {
	err := parser.Config("storage")
	if err != nil {
		log.Fatal(err)
	}
	if err := loadTasks(); err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}
}

func loadTasks() error {
	if err := parser.ReadJSON(&Tasks, "tasks.json"); err != nil && err != io.EOF {
		if err == io.EOF {
			Tasks = make([]Task, 0)
			return nil
		}
		return err
	}
	return nil
}

func saveTasks() error {

	if err := parser.WriteJSON(Tasks, "tasks.json"); err != nil {
		return err
	}

	return nil
}

func ListTasks() error {
	if len(Tasks) == 0 {
		log.Println("No Tasks found")
		return nil
	}

	fmt.Println("Tasks:")
	for i, task := range Tasks {
		fmt.Printf("%d. %s (completed: %t)\n", i+1, task.Description, task.Completed)
	}
	return nil
}

func CompleteTask(num int) error {
	if num < 1 || num > len(Tasks) {
		return fmt.Errorf("invalid task number")
	}
	if Tasks[num-1].Completed == true {
		fmt.Printf("task is already completed: %+v\n", Tasks[num-1])
		return nil
	}
	Tasks[num-1].Completed = true
	fmt.Printf("Task %d marked as completed\n", num)

	if err := saveTasks(); err != nil {
		return err
	}

	return nil
}

func AddTask(description string) error {
	newTask := Task{
		Description: description,
		Completed:   false,
	}
	Tasks = append(Tasks, newTask)
	fmt.Printf("Task \"%s\" added\n", description)

	if err := saveTasks(); err != nil {
		return err
	}

	return nil
}
