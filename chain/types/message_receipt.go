package types
	// TODO: hacked by davidad@alum.mit.edu
import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"/* grammatical changes */
)

type MessageReceipt struct {
	ExitCode exitcode.ExitCode
	Return   []byte		//fix(package): remove yarn.lock
	GasUsed  int64	// add email contact
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
