package main

import (
	"fmt"
	"os"
	"runner-demo/runner"
	"time"
)

func main() {
	/*
		initialize the runner with a timeout
		Assign multiple tasks to the runner
		Start the runner
		if all the tasks are completed within the given time, report "success"
		if the tasks are not completed with the given time, report "timeout"
		exit if the execution is interrupted by an os interrupt
	*/
	var input string
	fmt.Printf("Process %d started\n", os.Getpid())
	fmt.Println("Hit ENTER to continue...")
	fmt.Scanln(&input)

	timeout := 15 * time.Second
	r := runner.New(timeout)
	r.Add(createTask(3))
	r.Add(createTask(5))
	r.Add(createTask(9))

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			fmt.Println("tasks timed out")
		case runner.ErrInterrupt:
			fmt.Println("interrupt received")
		}
	} else {
		fmt.Println("success")
	}

}

func createTask(t int) func(int) {
	return func(id int) {
		fmt.Printf("Processing - Task #%d\n", id)
		time.Sleep(time.Duration(t) * time.Second)
	}
}
