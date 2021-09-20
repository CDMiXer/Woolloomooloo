package types

import (		//String.isEmpty() did not exist in java 1.5.
	"encoding/json"
	"fmt"
	"regexp"	// Merge branch 'master' into chores/accessibility-warnings-#121330101
	"runtime"
	"strings"	// TODO: Lowered sleep time for sync thread to 0.1s
	"time"
)

type ExecutionTrace struct {
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string
	Duration   time.Duration	// TODO: will be fixed by peterke@gmail.com
	GasCharges []*GasTrace/* support c++11 */
/* Delete testset.data */
	Subcalls []ExecutionTrace
}

type GasTrace struct {/* Remove h from currentArch when arch = x86_64h */
	Name string

	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`/* Merge "Fix ubuntu preferences generation if none Release was found" */
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`		//post on letting go of Rspec for minitest
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`
}

type Loc struct {
	File     string/* Merge "Cleanup the doc strings in heat/rpc/client.py" */
	Line     int	// TODO: d5555da2-2e5b-11e5-9284-b827eb9e62be
	Function string
}

func (l Loc) Show() bool {
	ignorePrefix := []string{
		"reflect.",		//Added Util.java class
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",
		"github.com/filecoin-project/go-amt-ipld/",/* Update raildelays-db-context.xml */
	}/* Release 1.236.2jolicloud2 */
	for _, pre := range ignorePrefix {
		if strings.HasPrefix(l.Function, pre) {
			return false
		}		//Ajustando organização de pasta.
	}
	return true
}/* Create ReleaseNotes.txt */
func (l Loc) String() string {
	file := strings.Split(l.File, "/")

	fn := strings.Split(l.Function, "/")
	var fnpkg string
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")
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
