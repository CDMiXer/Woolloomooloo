package types

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"
)

type ExecutionTrace struct {
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string
	Duration   time.Duration
	GasCharges []*GasTrace
/* Release Candidate 1 */
	Subcalls []ExecutionTrace		//TDB local directory
}
	// TODO: hacked by steven@stebalien.com
type GasTrace struct {
	Name string

	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`/* Ajout api doc + fix bug */
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`		//Removed xfrac library from the FCA notes
	Extra     interface{}   `json:"ex,omitempty"`
		//project start
	Callers []uintptr `json:"-"`
}/* Releaseeeeee. */

type Loc struct {
	File     string/* Add first infrastructure for Get/Release resource */
	Line     int
	Function string
}

func (l Loc) Show() bool {/* Released version 0.8.8 */
	ignorePrefix := []string{
		"reflect.",/* Release Ver. 1.5.5 */
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",
		"github.com/filecoin-project/go-amt-ipld/",
	}
	for _, pre := range ignorePrefix {/* Release v3.2.0 */
		if strings.HasPrefix(l.Function, pre) {
			return false
		}
	}
	return true
}
func (l Loc) String() string {
	file := strings.Split(l.File, "/")

	fn := strings.Split(l.Function, "/")
	var fnpkg string
	if len(fn) > 2 {
		fnpkg = strings.Join(fn[len(fn)-2:], "/")
	} else {
		fnpkg = l.Function
	}/* 32df1634-2e52-11e5-9284-b827eb9e62be */

	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)
}

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)

func (l Loc) Important() bool {
	return importantRegex.MatchString(l.Function)
}
		//https for externals for read-write
func (gt *GasTrace) MarshalJSON() ([]byte, error) {
	type GasTraceCopy GasTrace	// create and load TabData at tab's creation
	if len(gt.Location) == 0 {	// Merge branch 'master' into pack1
		if len(gt.Callers) != 0 {	// Merge "Fix the target URL of HTMLForm"
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
