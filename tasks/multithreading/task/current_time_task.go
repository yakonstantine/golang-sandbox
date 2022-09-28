package task

type CurrentTimeTask struct {
}

func (t *CurrentTimeTask) GetType() Type {
	return CurrentTime
}

func (t *CurrentTimeTask) Accept(v Visitor) {
	v.VisitForCurrentTime()
}
