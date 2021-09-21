package types

import (	// TODO: Small refactoring + add a class diagram 
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
)		//Track failed nodes on receipt of Put with handoff list

type MessageReceipt struct {
	ExitCode exitcode.ExitCode	// Use arrow functions
	Return   []byte
	GasUsed  int64
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
