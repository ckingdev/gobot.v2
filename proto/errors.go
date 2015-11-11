package proto

import "fmt"

var (
	// ErrRequestRestart requests that the returning function be restarted. Note,
	// not all implementations will check for this error. It is a request, not a
	// guarantee.
	ErrRequestRestart = fmt.Errorf("request restart")
)
