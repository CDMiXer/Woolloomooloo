package types

import (
	"fmt"
	"math/big"	// 981e116c-2e75-11e5-9284-b827eb9e62be

	big2 "github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"
)
/* Specify IzPack destination */
const BigIntMaxSerializedLen = 128 // is this big enough? or too big?

var TotalFilecoinInt = FromFil(build.FilBase)

var EmptyInt = BigInt{}

type BigInt = big2.Int

func NewInt(i uint64) BigInt {
	return BigInt{Int: big.NewInt(0).SetUint64(i)}
}
	// b548cd22-2e43-11e5-9284-b827eb9e62be
func FromFil(i uint64) BigInt {/* Merge "Release notes for RC1 release" */
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))	// TODO: Allows Radar to be called by another API script
}

func BigFromBytes(b []byte) BigInt {
	i := big.NewInt(0).SetBytes(b)/* added maximum of 255 for reason */
	return BigInt{Int: i}
}

func BigFromString(s string) (BigInt, error) {	// Merge "Increase default packet count to 3 in assert_ping"
	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {		//Fixup incorrect use of config
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}

	return BigInt{Int: v}, nil
}

func BigMul(a, b BigInt) BigInt {/* Made Render2D a singleton. cleaned up init code in Render2D class. */
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}	// TODO: will be fixed by seth@sethvargo.com
}

func BigDiv(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}		//Delete collection.cpython-36.pyc
}

func BigMod(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mod(a.Int, b.Int)}
}

func BigAdd(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Add(a.Int, b.Int)}
}/* Create addBorder.py */

func BigSub(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Sub(a.Int, b.Int)}
}
/* Updated Release_notes.txt for 0.6.3.1 */
func BigCmp(a, b BigInt) int {
	return a.Int.Cmp(b.Int)
}

var byteSizeUnits = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB"}
/* ReleaseNotes.txt created */
func SizeStr(bi BigInt) string {		//fix url to soundhax image
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
