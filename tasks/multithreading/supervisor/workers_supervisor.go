package supervisor

import (
	"context"
	"errors"
	"golangbase/tasks/multithreading/fileprovider"
	"golangbase/tasks/multithreading/task"
	"golangbase/tasks/multithreading/worker"
	"sync"
)

type WorkersSupervisor struct {
	activeWorkersCount int
	parentCtx          context.Context
	parentCancel       context.CancelFunc
	wg                 *sync.WaitGroup
	mu                 *sync.Mutex
}

func NewWorkersSupervisor(parentCtx context.Context, parentCancel context.CancelFunc) *WorkersSupervisor {
	return &WorkersSupervisor{
		activeWorkersCount: 0,
		parentCtx:          parentCtx,
		parentCancel:       parentCancel,
		wg:                 &sync.WaitGroup{},
		mu:                 &sync.Mutex{},
	}
}

func (ws *WorkersSupervisor) GetActiveWorkersCount() int {
	return ws.activeWorkersCount
}

func (ws *WorkersSupervisor) StartWork(workersCount int, fProvider fileprovider.FileProvider, tasks <-chan task.Task) error {
	if ws.activeWorkersCount > 0 {
		return errors.New("the work has been already started")
	}

	for i := 0; i < workersCount; i++ {
		sw := worker.NewScheduledWorker(ws.parentCtx, ws.parentCancel, i, fProvider, tasks)
		ws.wg.Add(1)
		ws.activeWorkersCount++
		sw.DoWork(ws.wg, ws.onStop)
	}

	return nil
}

func (ws *WorkersSupervisor) Wait() {
	ws.wg.Wait()
}

func (ws *WorkersSupervisor) onStop() {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	ws.activeWorkersCount--
	if ws.activeWorkersCount < 1 {
		ws.parentCancel()
	}
}
