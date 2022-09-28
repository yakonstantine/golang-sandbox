package task

type Visitor interface {
	VisitForExit()
	VisitForCurrentTime()
	VisitForCurrentTimeWait()
	VisitForReadWriteToFile()
	VisitForWorkerShutDown()
}
