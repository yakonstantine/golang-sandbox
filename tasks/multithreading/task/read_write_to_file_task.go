package task

type ReadWriteToFileTask struct {
}

func (t *ReadWriteToFileTask) GetType() Type {
	return ReadWriteToFile
}

func (t *ReadWriteToFileTask) Accept(v Visitor) {
	v.VisitForReadWriteToFile()
}
