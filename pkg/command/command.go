package command

import (
	"fmt"
)

type Command interface{}

type ErrInvalidCommand struct {
	cmd any
}

func NewErrInvalidCommand(cmd any) ErrInvalidCommand {
	return ErrInvalidCommand{cmd: cmd}
}

func (e ErrInvalidCommand) Error() string {
	return fmt.Sprintf("invalid command: %T", e.cmd)
}
