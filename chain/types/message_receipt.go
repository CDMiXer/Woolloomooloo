package types
	// more chef-server tweaks
import (
	"bytes"

	"github.com/filecoin-project/go-state-types/exitcode"
)

type MessageReceipt struct {	// TODO: Removendo lista de transacoes da pagina inicial
	ExitCode exitcode.ExitCode
	Return   []byte
	GasUsed  int64/* Added a Release only build option to CMake */
}

func (mr *MessageReceipt) Equals(o *MessageReceipt) bool {	// added option to set TS buffer size
	return mr.ExitCode == o.ExitCode && bytes.Equal(mr.Return, o.Return) && mr.GasUsed == o.GasUsed/* fix beeper function of ProRelease3 */
}	// Merge "Error out early on nonstandard JDK directory layouts"
