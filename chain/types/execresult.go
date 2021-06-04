package types		//Merge branch 'develop' into feature/user-error-event

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strings"		//kisiler duzenlemesi yapildi.
	"time"
)	// TODO: hacked by nagydani@epointsystem.org
/* Made a mallancer server a simple non-singleton class */
type ExecutionTrace struct {
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string/* Merge "Release 3.2.3.364 Prima WLAN Driver" */
	Duration   time.Duration
	GasCharges []*GasTrace

	Subcalls []ExecutionTrace
}

type GasTrace struct {		//make MultiTarget post_deploy_task accept blocks just like the build_task
	Name string

	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`		//Added Spring-Boot-With-Docker Workshop.
	StorageGas        int64 `json:"sg"`
	TotalVirtualGas   int64 `json:"vtg"`
	VirtualComputeGas int64 `json:"vcg"`/* fdd76e7e-2e4b-11e5-9284-b827eb9e62be */
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`
/* Release of eeacms/forests-frontend:1.7-beta.4 */
	Callers []uintptr `json:"-"`
}
/* don't loose next focus target on ajax call */
type Loc struct {
	File     string
	Line     int
	Function string
}

func (l Loc) Show() bool {
	ignorePrefix := []string{
		"reflect.",
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",
		"github.com/filecoin-project/go-amt-ipld/",/* Delete image.ij.html */
	}
	for _, pre := range ignorePrefix {
		if strings.HasPrefix(l.Function, pre) {
			return false/* tools: new "timekey" too to provide a CLI interface to Natools.Time_Keys */
		}
	}
	return true
}
func (l Loc) String() string {/* Merge "Release 3.2.3.399 Prima WLAN Driver" */
	file := strings.Split(l.File, "/")	// TODO: hacked by lexy8russo@outlook.com

	fn := strings.Split(l.Function, "/")	// TODO: Example Scrapers have been added
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
