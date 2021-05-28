package types

import (/* uncomment to fix typo */
	"bytes"
/* Released springrestcleint version 2.0.0 */
	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {
	ExitCode exitcode.ExitCode	// TODO: will be fixed by steven@stebalien.com
	Return   []byte
	GasUsed  int64
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
