package producer

import (
	"context"
	"fmt"
	"golangbase/tasks/multithreading/task"
)

type ArrayProducer struct {
	tasks  []int
	ctx    context.Context
	cancel context.CancelFunc
}

func NewArrayProducer(parentCtx context.Context, tasks []int) *ArrayProducer {
	ctx, cancel := context.WithCancel(parentCtx)
	return &ArrayProducer{
		tasks:  tasks,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (p *ArrayProducer) Produce(in chan<- task.Task) error {
	defer p.cancel()
	for _, t := range p.tasks {
		ok, err := p.addToChan(in, t)
		if err != nil {
			return err
		}
		if !ok {
			break
		}
	}

	fmt.Printf("All tasks have been produced\n")
	return nil
}

func (p *ArrayProducer) addToChan(in chan<- task.Task, tskNum int) (bool, error) {
	tt := task.Type(tskNum)
	tsk, err := task.NewTask(tt)
	if err != nil {
		return false, err
	}

	for {
		select {
		case <-p.ctx.Done():
			fmt.Printf("Producer stop\n")
			return false, nil
		case in <- tsk:
			return true, nil
		default:
		}
	}
}
