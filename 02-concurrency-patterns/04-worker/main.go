package main

import (
	"fmt"
	"time"
	"worker-demo/worker"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
	"Magesh",
	"Ganesh",
	"Ramesh",
	"Rajesh",
	"Suresh",
}

type NamePrinter struct {
	name  string
	delay time.Duration
}

func (np *NamePrinter) Task() {
	fmt.Println("Task execution commenced")
	time.Sleep(np.delay)
	fmt.Println("Name Printer - Name : ", np.name)
}

func main() {
	/* 5 = no of concurrent work jobs to be executed */
	timerCounter := 1
	p := worker.New(5)
	lrt := &LongRunningTask{}
	p.Run(lrt)
	for idx := 0; idx < 2; idx++ {
		for _, name := range names {
			np := NamePrinter{
				name:  name,
				delay: time.Duration(timerCounter*100) * time.Millisecond,
			}
			timerCounter++
			p.Run(&np)
		}
	}
	fmt.Println("All tasks are assigned")
	p.Shutdown()
}

type LongRunningTask struct {
}

func (p *LongRunningTask) Task() {
	fmt.Println("LRT commenced")
	time.Sleep(3 * time.Minute)
	fmt.Println("LRT completed")
}
