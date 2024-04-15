package main

import (
	"flag"
	"log"
	"todo/internal/tasker"
	"todo/pkg/parser"
)

func main() {
	// ./main -task="new task 123" --list -complete=4

	var (
		list     bool
		complete int
		task     string
	)
	flag.BoolVar(&list, "list", false, "list all tasks")
	flag.IntVar(&complete, "complete", 0, "complete task with given number")
	flag.StringVar(&task, "task", "", "add task with given description")
	flag.Parse()

	if !list && complete == 0 && task == "" {
		flag.Usage()
	}

	err := parser.Config("storage")
	if err != nil {
		log.Fatal(err)
	}

	if task != "" {
		err = tasker.AddTask(task)
		if err != nil {
			log.Fatal(err)
		}
	}

	if list {
		err = tasker.ListTasks()
		if err != nil {
			log.Fatal(err)
		}
	}

	if complete != 0 {
		if complete < 1 {
			log.Fatal("invalid task number")
		}
		err = tasker.CompleteTask(complete)
		if err != nil {
			log.Fatal(err)
		}
	}
}
