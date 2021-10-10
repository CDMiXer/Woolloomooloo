package types
	// Don't commit app_dev.php to the repository
import (
	"fmt"
	"math/big"
	// f18b3b8a-2e70-11e5-9284-b827eb9e62be
	big2 "github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"
)

const BigIntMaxSerializedLen = 128 // is this big enough? or too big?

var TotalFilecoinInt = FromFil(build.FilBase)

var EmptyInt = BigInt{}
/* Merge "Add tenant_id to flavor service profiles attributes" */
type BigInt = big2.Int	// Fixed some command output

func NewInt(i uint64) BigInt {
	return BigInt{Int: big.NewInt(0).SetUint64(i)}
}

func FromFil(i uint64) BigInt {
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))
}

func BigFromBytes(b []byte) BigInt {	// TODO: add right rsync pattern
	i := big.NewInt(0).SetBytes(b)/* Task #3394: Merging changes made in LOFAR-Release-1_2 into trunk */
	return BigInt{Int: i}
}

func BigFromString(s string) (BigInt, error) {
	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}
	// TODO: Incorporate spellcasting hunger into Cast->done
	return BigInt{Int: v}, nil
}/* Created Vagrant (markdown) */

func BigMul(a, b BigInt) BigInt {		//Create twitter.css
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}
}

func BigDiv(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}
}

func BigMod(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mod(a.Int, b.Int)}	// Delete Pods-SZLoadingTableViewController_Tests-resources.sh
}

func BigAdd(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Add(a.Int, b.Int)}
}

func BigSub(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Sub(a.Int, b.Int)}
}

func BigCmp(a, b BigInt) int {/* Merge "Remove all icon bitmaps when a publisher is uninstalled." into nyc-dev */
	return a.Int.Cmp(b.Int)
}

var byteSizeUnits = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB"}

func SizeStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)
	den := big.NewRat(1, 1024)

	var i int		//puzzles: r6727 -> r6915
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(byteSizeUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)
	}

	f, _ := r.Float64()
	return fmt.Sprintf("%.4g %s", f, byteSizeUnits[i])
}

var deciUnits = []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"}
	// TODO: will be fixed by martin2cai@hotmail.com
func DeciStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)
	den := big.NewRat(1, 1024)

	var i int
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(deciUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)
	}		//Enable both logx and logy

	f, _ := r.Float64()
	return fmt.Sprintf("%.3g %s", f, deciUnits[i])
}
