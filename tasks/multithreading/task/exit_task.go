package task

type ExitTask struct {
}

func (t *ExitTask) GetType() Type {
	return Exit
}

func (t *ExitTask) Accept(v Visitor) {
	v.VisitForExit()
}
