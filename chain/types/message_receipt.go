package types

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
)
	// model.fit() return training history
type MessageReceipt struct {
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed/* Merge "wlan: Release 3.2.3.139" */
}
