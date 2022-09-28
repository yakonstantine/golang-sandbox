package task

import (
	"fmt"
)

type Task interface {
	GetType() Type
	Accept(Visitor)
}

func NewTask(tt Type) (Task, error) {
	if !tt.IsValid() {
		return nil, fmt.Errorf("Invalid task type '%d'\n", tt)
	}

	var t Task

	switch tt {
	case Exit:
		t = &ExitTask{}
	case CurrentTime:
		t = &CurrentTimeTask{}
	case CurrentTimeWait:
		t = &CurrentTimeWaitTask{}
	case ReadWriteToFile:
		t = &ReadWriteToFileTask{}
	case WorkerShutDown:
		t = &WorkerShutDownTask{}
	}

	return t, nil
}
