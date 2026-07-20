package cuda

type oomError struct {
	res       int64
	allocated int64
}

func (e oomError) Reserved() int64  { _ = "STUB: not implemented"; return 0 }
func (e oomError) Allocated() int64 { _ = "STUB: not implemented"; return 0 }
func (e oomError) Error() string    { _ = "STUB: not implemented"; return "" }

const (
	typeMismatch  = "TypeMismatch: a %T and b %T"
	shapeMismatch = "Shape mismatch. Expected %v. Got %v"
)
