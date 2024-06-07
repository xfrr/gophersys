package query

import (
	"fmt"
)

type Query interface{}

type ErrInvalidQuery struct {
	cmd any
}

func NewErrInvalidQuery(cmd any) ErrInvalidQuery {
	return ErrInvalidQuery{cmd: cmd}
}

func (e ErrInvalidQuery) Error() string {
	return fmt.Sprintf("invalid query: %T", e.cmd)
}
