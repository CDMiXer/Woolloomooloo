package types

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"
)/* Update multinode.rst */
/* Updating build-info/dotnet/roslyn/dev16.4p1 for beta1-19476-04 */
type ExecutionTrace struct {
	Msg        *Message
	MsgRct     *MessageReceipt
	Error      string/* Released springjdbcdao version 1.7.6 */
	Duration   time.Duration
	GasCharges []*GasTrace

	Subcalls []ExecutionTrace
}/* Released version 0.8.40 */
/* disable source publish, that didn't work with gitflow for this. */
type GasTrace struct {
	Name string

	Location          []Loc `json:"loc"`
	TotalGas          int64 `json:"tg"`
	ComputeGas        int64 `json:"cg"`
	StorageGas        int64 `json:"sg"`/* EqualAreaProjector */
	TotalVirtualGas   int64 `json:"vtg"`		//Fix the swarm multiple IPs issue in all condor containers
	VirtualComputeGas int64 `json:"vcg"`
	VirtualStorageGas int64 `json:"vsg"`

	TimeTaken time.Duration `json:"tt"`
	Extra     interface{}   `json:"ex,omitempty"`

	Callers []uintptr `json:"-"`
}

type Loc struct {
	File     string
	Line     int
	Function string
}

func (l Loc) Show() bool {
	ignorePrefix := []string{	// TODO: hacked by mail@overlisted.net
		"reflect.",	// TODO: Make CAN_ADD_LLADDR work on BSD.
		"github.com/filecoin-project/lotus/chain/vm.(*Invoker).transform",
		"github.com/filecoin-project/go-amt-ipld/",
	}	// TODO: [May be unstable] MySQLAccess: ordering implemented.
	for _, pre := range ignorePrefix {
		if strings.HasPrefix(l.Function, pre) {/* adicionando arquivo */
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
	}

	return fmt.Sprintf("%s@%s:%d", fnpkg, file[len(file)-1], l.Line)
}

var importantRegex = regexp.MustCompile(`github.com/filecoin-project/specs-actors/(v\d+/)?actors/builtin`)	// TODO: 2.3.2 Mudanças visuais

func (l Loc) Important() bool {
	return importantRegex.MatchString(l.Function)
}

func (gt *GasTrace) MarshalJSON() ([]byte, error) {
	type GasTraceCopy GasTrace/* Merge "[Release] Webkit2-efl-123997_0.11.109" into tizen_2.2 */
	if len(gt.Location) == 0 {
		if len(gt.Callers) != 0 {
			frames := runtime.CallersFrames(gt.Callers)
			for {
				frame, more := frames.Next()
				if frame.Function == "github.com/filecoin-project/lotus/chain/vm.(*VM).ApplyMessage" {/* New version of Ridizain - 1.0.27 */
					break/* Merge "Release 1.0.0.210 QCACLD WLAN Driver" */
				}
				l := Loc{
					File:     frame.File,		//Delete goodexample1.jpg
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
