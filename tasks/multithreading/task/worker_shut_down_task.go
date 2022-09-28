package task

type WorkerShutDownTask struct {
}

func (t *WorkerShutDownTask) GetType() Type {
	return WorkerShutDown
}

func (t *WorkerShutDownTask) Accept(v Visitor) {
	v.VisitForWorkerShutDown()
}
