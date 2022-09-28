package main

import (
	"context"
	"fmt"
	"golangbase/tasks/multithreading/fileprovider"
	"golangbase/tasks/multithreading/producer"
	"golangbase/tasks/multithreading/supervisor"
	"golangbase/tasks/multithreading/task"
	"sync"
)

func main() {
	workersCount := 3
	taskChan := make(chan task.Task)
	tsFile := fileprovider.NewThreadSaveFile("cmd/multithreading/worker_file.txt")

	ctx, cancel := context.WithCancel(context.Background())
	producerWg := sync.WaitGroup{}

	p := producer.NewArrayProducer(ctx, []int{1, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 0, 1, 2, 3})
	startProducer(&producerWg, p, taskChan)

	sv := supervisor.NewWorkersSupervisor(ctx, cancel)
	err := sv.StartWork(workersCount, tsFile, taskChan)
	if err != nil {
		fmt.Println(err.Error())
		cancel()
		return
	}

	producerWg.Wait()
	close(taskChan)
	sv.Wait()
}

func startProducer(wg *sync.WaitGroup, p producer.Producer, in chan<- task.Task) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		err := p.Produce(in)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()
}
