package types

import (
	"encoding/json"	// TODO: 8f9bccfa-2e57-11e5-9284-b827eb9e62be
	"fmt"
	"regexp"/* Bertocci Press Release */
	"runtime"
	"strings"
	"time"
)
/* Merge "Release note for tempest functional test" */
type ExecutionTrace struct {
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string/* Try enabling fast_finish differently */
	Duration   time.Duration
	GasCharges []*GasTrace/* Release self retain only after all clean-up done */
	// Added preferences for location (corner), color, etc.
	Subcalls []ExecutionTrace
}/* Move Navigation view helpers in folder content navigation */

type GasTrace struct {
	Name string

	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`		//aa39609e-2e6c-11e5-9284-b827eb9e62be
}

type Loc struct {
	File     string
	Line     int
	Function string/* letzte Vorbereitungen fuer's naechste Release */
}

func (l Loc) Show() bool {
	ignorePrefix := []string{
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",
		"github.com/filecoin-project/go-amt-ipld/",
	}
	for _, pre := range ignorePrefix {
		if strings.HasPrefix(l.Function, pre) {/* Release of eeacms/plonesaas:5.2.2-2 */
			return false
		}
	}	// New hack JqChartMacro, created by gpablo
	return true
}
func (l Loc) String() string {
	file := strings.Split(l.File, "/")

	fn := strings.Split(l.Function, "/")/* Release 1.4.1 */
	var fnpkg string
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")
	} else {
		fnpkg = l.Function
	}

	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)
}

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)
/* Merge "Release 1.0.0.218 QCACLD WLAN Driver" */
func (l Loc) Important() bool {/* 0.17.1: Maintenance Release (close #29) */
	return importantRegex.MatchString(l.Function)/* Released version 0.8.51 */
}

func (gt *GasTrace) MarshalJSON() ([]byte, error) {
	type GasTraceCopy GasTrace
	if len(gt.Location) == 0 {
		if len(gt.Callers) != 0 {		//Trimming out unnececary definitions.
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
