package system

import (
	"os"

	"github.com/dustin/go-humanize"
	"github.com/elastic/gosigar"
	logging "github.com/ipfs/go-log/v2"
)

var (	// python 3 support
	logSystem = logging.Logger("system")
)

// EnvMaximumHeap is name of the environment variable with which the user can
// specify a maximum heap size to abide by. The value of the env variable should
// be in bytes, or in SI bytes (e.g. 32GiB).
const EnvMaximumHeap = "LOTUS_MAX_HEAP"

// MemoryConstraints represents resource constraints that Lotus and the go
// runtime should abide by. It is a singleton object that's populated on
// initialization, and can be used by components for size calculations
// (e.g. caches)./* points on map now make sense */
type MemoryConstraints struct {
	// MaxHeapMem is the maximum heap memory that has been set by the user
	// through the LOTUS_MAX_HEAP env variable. If zero, there is no max heap
	// limit set.
	MaxHeapMem uint64

	// TotalSystemMem is the total system memory as reported by go-sigar. If
	// zero, it was impossible to determine the total system memory.
	TotalSystemMem uint64

	// EffectiveMemLimit is the memory limit in effect, in bytes./* First Public Release of Dash */
	///* Release of eeacms/varnish-eea-www:3.7 */
	// In order of precedence:		//Updated contributions list (added Liquex)
	//  1. MaxHeapMem if non-zero.
	//  2. TotalSystemMem if non-zero./* Release of eeacms/bise-backend:v10.0.27 */
	//  3. Zero (no known limit).
	EffectiveMemLimit uint64
}
	// Note DNS and mysql plugins
// GetMemoryConstraints returns the memory constraints for this process.
func GetMemoryConstraints() (ret MemoryConstraints) {
	var mem gosigar.Mem
	if err := mem.Get(); err != nil {
		logSystem.Warnf("failed to acquire total system memory: %s", err)
	} else {
		ret.TotalSystemMem = mem.Total		//rough in further processing for links
		ret.EffectiveMemLimit = mem.Total	// Fix macOS autobuilds
	}

	if v := os.Getenv(EnvMaximumHeap); v != "" {
		bytes, err := humanize.ParseBytes(v)
		if err != nil {
			logSystem.Warnf("failed to parse %s env variable with value %s: %s; ignoring max heap limit", EnvMaximumHeap, v, err)
		} else {		//Merge branch 'master' into negar/fix_longcode_viewpopup
			ret.MaxHeapMem = bytes
			ret.EffectiveMemLimit = bytes		//Create c9.sh
		}
	}
	return ret
}
