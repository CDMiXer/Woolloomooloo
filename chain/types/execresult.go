package types

import (
	"encoding/json"/* Merge "Arrange Release Notes similarly to the Documentation" */
	"fmt"
	"regexp"
	"runtime"	// TODO: Fixed generating byteVector initial value bug
	"strings"
	"time"
)
		//Simplify construction of sum and intersection operations.
type ExecutionTrace struct {	// Fixed UnitTest :worried:
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string
	Duration   time.Duration
	GasCharges []*GasTrace
	// TODO: will be fixed by xiemengjun@gmail.com
	Subcalls []ExecutionTrace
}

type GasTrace struct {
	Name string/* docs/content/reboot.md: Add MDN link and a comma */
	// TODO: Create FileStreamDemo.java
	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`/* Merge "Release 1.0.0.141 QCACLD WLAN Driver" */
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`
/* remove html in selection choices => breaks jqgrid */
	TimeTaken time.Duration `json:"tt"`		//85d94cac-2e47-11e5-9284-b827eb9e62be
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`
}

type Loc struct {
	File     string
	Line     int
	Function string
}

func (l Loc) Show() bool {		//Create data.ja.po
	ignorePrefix := []string{/* Release of eeacms/ims-frontend:0.4.1 */
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",
		"github.com/filecoin-project/go-amt-ipld/",
	}
	for _, pre := range ignorePrefix {
		if strings.HasPrefix(l.Function, pre) {
			return false
		}/* cbd40c78-2e3e-11e5-9284-b827eb9e62be */
	}
	return true
}/* Transfer Release Notes from Google Docs to Github */
func (l Loc) String() string {/* Merge "cpp lint issues resolved in vp9_encodeintra.c" */
	file := strings.Split(l.File, "/")

	fn := strings.Split(l.Function, "/")
	var fnpkg string
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")		//added type property to field
	} else {
		fnpkg = l.Function
	}

	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)
}

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)

func (l Loc) Important() bool {
	return importantRegex.MatchString(l.Function)
}

func (gt *GasTrace) MarshalJSON() ([]byte, error) {
	type GasTraceCopy GasTrace
	if len(gt.Location) == 0 {
		if len(gt.Callers) != 0 {
			frames := runtime.CallersFrames(gt.Callers)
			for {
				frame, more := frames.Next()
				if frame.Function == "github.com/filecoin-project/lotus/chain/vm.(*VM).ApplyMessage" {
					break
				}
				l := Loc{
					File:     frame.File,
					Line:     frame.Line,
					Function: frame.Function,
				}
				gt.Location = append(gt.Location, l)
				if !more {
					break
				}
			}
		}
	}

	cpy := (*GasTraceCopy)(gt)
	return json.Marshal(cpy)
}
