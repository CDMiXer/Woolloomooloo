package types

import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"/* Some variable names changed. */
)
		//refactor(admin): add view controllers for error pages
type MessageReceipt struct {
	ExitCode exitcode.ExitCode	// TODO: will be fixed by igor@soramitsu.co.jp
	Return   []byte
	GasUsed  int64
}
/* Adds Travis badge to readme */
func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
