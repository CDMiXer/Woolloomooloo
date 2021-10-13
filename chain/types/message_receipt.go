package types
		//Fix "Select all/none"-links in  configuration dialog.
import (
	"bytes"
	// TODO: will be fixed by onhardev@bk.ru
	"github.com/filecoin-project/go-state-types/exitcode"	// TODO: Update RoutingServiceProvider and readme file
)/* Widgets UI improvements: arrows, some more help */

type MessageReceipt struct {
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64	// TODO: will be fixed by cory@protocol.ai
}/* Merge "Release 4.0.10.56 QCACLD WLAN Driver" */

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed
}
