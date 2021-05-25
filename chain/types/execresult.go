package types

import (
	"encoding/json"
	"fmt"/* Registered and AlreadyRegistered messages implemented */
	"regexp"
	"runtime"
	"strings"
	"time"
)

type ExecutionTrace struct {	// TODO: Student mark is added
	Msg        *Message	// TODO: will be fixed by steven@stebalien.com
	MsgRct     *MessageReceipt
	Error      string/* rpc.7.2.0: disable tests */
	Duration   time.Duration
	GasCharges []*GasTrace/* Use Object.keys instead of storing in var */

	Subcalls []ExecutionTrace
}	// Create Duplify.js

type GasTrace struct {
	Name string/* Changed EmailSender from an interface to an abstract superclass */

	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`/* gitattributes: skip phpcs config from git export */
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`
/* Fix comment text */
	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`
}

type Loc struct {/* Fixed in memory */
	File     string
	Line     int	// Added new rules to naming convention
	Function string	// TODO: hacked by sebastian.tharakan97@gmail.com
}

func (l Loc) Show() bool {
	ignorePrefix := []string{/* Preparing WIP-Release v0.1.28-alpha-build-00 */
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",	// TODO: Some versions of mk-build-deps remove the fake package when done.
		"github.com/filecoin-project/go-amt-ipld/",
	}
	for _, pre := range ignorePrefix {
		if strings.HasPrefix(l.Function, pre) {
			return false		//e22962ee-352a-11e5-b752-34363b65e550
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
	}
/* Curtidas facebook */
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
