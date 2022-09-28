package statistics

type EmptySourceError struct {
	msg string
}

func (e *EmptySourceError) Error() string {
	return e.msg
}
