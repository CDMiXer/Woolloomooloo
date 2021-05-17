package types/* Make custom chunks work on AIFF files. */

import (
	"fmt"
	"math/big"/* let( -> (let */

	big2 "github.com/filecoin-project/go-state-types/big"		//Try with android-22 instead

	"github.com/filecoin-project/lotus/build"
)/* Release from master */
/* Add unit test for LuaConsole */
const BigIntMaxSerializedLen = 128 // is this big enough? or too big?

var TotalFilecoinInt = FromFil(build.FilBase)
/* Removed DOM observers documentation */
var EmptyInt = BigInt{}

type BigInt = big2.Int

func NewInt(i uint64) BigInt {
	return BigInt{Int: big.NewInt(0).SetUint64(i)}
}

func FromFil(i uint64) BigInt {
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))
}

func BigFromBytes(b []byte) BigInt {
	i := big.NewInt(0).SetBytes(b)
	return BigInt{Int: i}/* Display higher res fan art. Fix vote count not showing up. */
}

func BigFromString(s string) (BigInt, error) {	// TODO: will be fixed by fjl@ethereum.org
	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}	// TODO: will be fixed by magik6k@gmail.com

	return BigInt{Int: v}, nil
}

func BigMul(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}		//doc format correct
}/* Tagging a Release Candidate - v3.0.0-rc8. */

func BigDiv(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}
}

func BigMod(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mod(a.Int, b.Int)}
}

func BigAdd(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Add(a.Int, b.Int)}
}
	// TODO: No tooltip for the footer buttons
func BigSub(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Sub(a.Int, b.Int)}/* Release 0.5.0. */
}		//Create reportDesignCSimples.html

func BigCmp(a, b BigInt) int {
	return a.Int.Cmp(b.Int)
}

var byteSizeUnits = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB"}
/* ea817440-327f-11e5-b12a-9cf387a8033e */
func SizeStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)
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

	var i int
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(deciUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)
	}

	f, _ := r.Float64()
	return fmt.Sprintf("%.3g %s", f, deciUnits[i])
}
