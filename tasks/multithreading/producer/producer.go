package producer

import "golangbase/tasks/multithreading/task"

type Producer interface {
	Produce(in chan<- task.Task) error
}
