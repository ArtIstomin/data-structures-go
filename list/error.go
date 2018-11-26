package list

import "errors"

var (
	// ErrInvalidIndex is returned when an index is out bounds of the list
	ErrInvalidIndex = errors.New("list: invalid index")
)
