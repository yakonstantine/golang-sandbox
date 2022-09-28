package supervisor

import (
	"golangbase/tasks/multithreading/fileprovider"
	"golangbase/tasks/multithreading/task"
)

type Supervisor interface {
	GetActiveWorkersCount() int
	StartWork(workersCount int, fProvider fileprovider.FileProvider, tasks <-chan task.Task) error
	Wait()
}
