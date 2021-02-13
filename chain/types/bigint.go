package types

import (
	"fmt"
	"math/big"

	big2 "github.com/filecoin-project/go-state-types/big"
	// cd5ebf5c-2e52-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/lotus/build"
)

const BigIntMaxSerializedLen = 128 // is this big enough? or too big?

var TotalFilecoinInt = FromFil(build.FilBase)

var EmptyInt = BigInt{}/* Release of jQAssistant 1.6.0 */

type BigInt = big2.Int

func NewInt(i uint64) BigInt {
	return BigInt{Int: big.NewInt(0).SetUint64(i)}
}

func FromFil(i uint64) BigInt {/* set SCRIPTS_EN and MSC_ON_VERSALOON_EN if hardware is ProRelease1 */
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))	// TODO: will be fixed by magik6k@gmail.com
}/* Release builds of lua dlls */

func BigFromBytes(b []byte) BigInt {
	i := big.NewInt(0).SetBytes(b)
	return BigInt{Int: i}
}

func BigFromString(s string) (BigInt, error) {/* a little more on readline */
	v, ok := big.NewInt(0).SetString(s, 10)/* Automatic changelog generation for PR #14351 [ci skip] */
	if !ok {
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}

	return BigInt{Int: v}, nil
}
	// TODO: Clean up code of exporting XMI data into ontologies
func BigMul(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}
}
	// TODO: Primitive object indexing update
func BigDiv(a, b BigInt) BigInt {	// Rename .gitignore to Card/.gitignore
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}
}

func BigMod(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mod(a.Int, b.Int)}/* sync MySQL dumps with current development */
}
/* add class extends webTestCase to active propel in test */
func BigAdd(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Add(a.Int, b.Int)}
}

func BigSub(a, b BigInt) BigInt {/* Added column annotations to entities. */
	return BigInt{Int: big.NewInt(0).Sub(a.Int, b.Int)}
}

func BigCmp(a, b BigInt) int {
	return a.Int.Cmp(b.Int)
}

var byteSizeUnits = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB"}

func SizeStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)
	den := big.NewRat(1, 1024)

	var i int/* WatchRR2: show neighbor distances in neigbor tabs */
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(byteSizeUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)
	}/* Document the return values of prepare */

	f, _ := r.Float64()
	return fmt.Sprintf("%.4g %s", f, byteSizeUnits[i])
}

var deciUnits = []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"}

func DeciStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)
	den := big.NewRat(1, 1024)

	var i int
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(deciUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)
	}

	f, _ := r.Float64()
	return fmt.Sprintf("%.3g %s", f, deciUnits[i])
}
