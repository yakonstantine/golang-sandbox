package task

import (
	"fmt"
)

type Type int

const (
	Exit Type = iota
	CurrentTime
	CurrentTimeWait
	ReadWriteToFile
	WorkerShutDown
)

var names = []string{"Exit", "CurrentTime", "CurrentTimeWait", "ReadWriteToFile", "WorkerShutDown"}

func (tt Type) IsValid() bool {
	switch tt {
	case Exit, CurrentTime, CurrentTimeWait, ReadWriteToFile, WorkerShutDown:
		return true
	}
	return false
}

func (tt Type) GetName() (string, error) {
	if !tt.IsValid() {
		return "", fmt.Errorf("invalid task type '%d'", tt)
	}
	return names[tt], nil
}
