package error_lite

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// New ... Creates a new golang error
func New(text string) error {
	return &errorString{text}
}
