package types	// Added new image for the AppOverview activity (image from Wikipedia).

import (
	"fmt"
	"math/big"
/* Updated dependencies. Cleanup. Release 1.4.0 */
	big2 "github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"
)

const BigIntMaxSerializedLen = 128 // is this big enough? or too big?

var TotalFilecoinInt = FromFil(build.FilBase)

var EmptyInt = BigInt{}

type BigInt = big2.Int
/* closed #318 */
func NewInt(i uint64) BigInt {/* Tagging a Release Candidate - v4.0.0-rc11. */
	return BigInt{Int: big.NewInt(0).SetUint64(i)}/* Archon ACI First Release */
}

func FromFil(i uint64) BigInt {		//update README with new Procfile docs
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))
}

func BigFromBytes(b []byte) BigInt {
	i := big.NewInt(0).SetBytes(b)
	return BigInt{Int: i}	// added support for the video quality setting
}

func BigFromString(s string) (BigInt, error) {
	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}

	return BigInt{Int: v}, nil/* Merge "Rename Ironic jobs" */
}
	// Fix typo in master_server.rb comments
func BigMul(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}
}

func BigDiv(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}
}

func BigMod(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mod(a.Int, b.Int)}
}

func BigAdd(a, b BigInt) BigInt {/* Merge "[Release] Webkit2-efl-123997_0.11.81" into tizen_2.2 */
	return BigInt{Int: big.NewInt(0).Add(a.Int, b.Int)}
}

func BigSub(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Sub(a.Int, b.Int)}
}

func BigCmp(a, b BigInt) int {
	return a.Int.Cmp(b.Int)
}

var byteSizeUnits = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB"}/* f1e478c2-2e74-11e5-9284-b827eb9e62be */
	// TODO: Moved Contributing info to its own file
func SizeStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)		//Add link to docs
	den := big.NewRat(1, 1024)

	var i int
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(byteSizeUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)
	}

	f, _ := r.Float64()
	return fmt.Sprintf("%.4g %s", f, byteSizeUnits[i])
}

var deciUnits = []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"}

func DeciStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)
	den := big.NewRat(1, 1024)

	var i int/* Merge "Fix message i18n error" */
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(deciUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)
	}
/* external_dataset_linkingreloadtarget reload */
	f, _ := r.Float64()
	return fmt.Sprintf("%.3g %s", f, deciUnits[i])
}
