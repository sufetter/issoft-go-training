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

func ListTasks() error {
	var tasks []Task
	if err := parser.ReadJSON(&tasks, "tasks.json"); err != nil {
		if err == io.EOF {
			log.Println("No tasks found")
			return nil
		}
		return err
	}

	fmt.Println("Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s (completed: %t)\n", i+1, task.Description, task.Completed)
	}
	return nil
}

func CompleteTask(num int) error {
	var tasks []Task
	if err := parser.ReadJSON(&tasks, "tasks.json"); err != nil {
		if err == io.EOF {
			log.Println("No tasks found")
			return nil
		}
		return err
	}

	if num < 1 || num > len(tasks) {
		return fmt.Errorf("invalid task number")
	}
	if tasks[num-1].Completed == true {
		fmt.Printf("task is already completed: %+v\n", tasks[num-1])
		return nil
	}
	tasks[num-1].Completed = true
	fmt.Printf("Task %d marked as completed\n", num)

	if err := parser.WriteJSON(tasks, "tasks.json"); err != nil {
		return err
	}

	return nil
}

func AddTask(description string) error {
	var tasks []Task
	if err := parser.ReadJSON(&tasks, "tasks.json"); err != nil && err != io.EOF {
		if os.IsNotExist(err) {
			log.Print("file doesn't exist, creating it...\n")
		} else {
			return err
		}
	}

	newTask := Task{
		Description: description,
		Completed:   false,
	}
	tasks = append(tasks, newTask)
	fmt.Printf("Task \"%s\" added\n", description)

	if err := parser.WriteJSON(tasks, "tasks.json"); err != nil {
		return err
	}

	return nil
}
