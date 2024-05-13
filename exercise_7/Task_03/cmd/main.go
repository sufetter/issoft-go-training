package main

import (
	"flag"
	"log"
	"todo/internal/api"
	"todo/internal/tasker"
)

func main() {
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

	var err error
	if task != "" {
		err = tasker.AddTask(task)
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

	if list {
		err = tasker.ListTasks()
		if err != nil {
			log.Fatal(err)
		}
	}
	
	err = api.SetMux("8080")
	if err != nil {
		log.Fatalf("Error starting server:\n%v", err)
	}
}
