package worker

import (
	"sync"
)

type Worker interface {
	GetNum() int
	DoWork(wg *sync.WaitGroup, onStop func())
	Cancel()
}
