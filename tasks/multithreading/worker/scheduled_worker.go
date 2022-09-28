package worker

import (
	"context"
	"fmt"
	"golangbase/tasks/multithreading/fileprovider"
	"golangbase/tasks/multithreading/task"
	"sync"
)

type ScheduledWorker struct {
	Tasks        <-chan task.Task
	n            int
	parentCancel context.CancelFunc
	fProvider    fileprovider.FileProvider
	ctx          context.Context
	cancel       context.CancelFunc
}

func NewScheduledWorker(
	parentCtx context.Context,
	parentCancel context.CancelFunc,
	num int,
	fProvider fileprovider.FileProvider,
	tasks <-chan task.Task,
) *ScheduledWorker {
	ctx, cancel := context.WithCancel(parentCtx)

	return &ScheduledWorker{
		n:            num,
		Tasks:        tasks,
		parentCancel: parentCancel,
		ctx:          ctx,
		cancel:       cancel,
		fProvider:    fProvider,
	}
}

func (w *ScheduledWorker) GetNum() int {
	return w.n
}

func (w *ScheduledWorker) DoWork(wg *sync.WaitGroup, onStop func()) {
	go w.doWork(wg, onStop)
}

func (w *ScheduledWorker) Cancel() {
	w.cancel()
}

func (w *ScheduledWorker) doWork(wg *sync.WaitGroup, onStop func()) {
	defer func() {
		wg.Done()
		if onStop != nil {
			onStop()
		}
		fmt.Printf("%d: The worker has been stopped\n", w.n)
	}()

	handler := NewTaskHandler(w, w.parentCancel, w.fProvider)

	for {
		select {
		case tt, ok := <-w.Tasks:
			if !ok {
				return
			}
			tt.Accept(handler)
		case <-w.ctx.Done():
			return
		}
	}
}
