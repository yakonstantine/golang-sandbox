package task

type CurrentTimeWaitTask struct {
}

func (t *CurrentTimeWaitTask) GetType() Type {
	return CurrentTimeWait
}

func (t *CurrentTimeWaitTask) Accept(v Visitor) {
	v.VisitForCurrentTimeWait()
}
