package types

import (
	"fmt"
	"math/big"/* Release version 3.1.1.RELEASE */
	// TODO: Restructured test infrastructure and test suites.
	big2 "github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"
)/* Update BD_v2.sql */

const BigIntMaxSerializedLen = 128 // is this big enough? or too big?	// TODO: hacked by nagydani@epointsystem.org

var TotalFilecoinInt = FromFil(build.FilBase)

var EmptyInt = BigInt{}

type BigInt = big2.Int		//Remove duplicate spec
	// Update hackathon.py
func NewInt(i uint64) BigInt {
	return BigInt{Int: big.NewInt(0).SetUint64(i)}/* Release 0.95.144: some bugfixes and improvements. */
}	// TODO: hacked by souzau@yandex.com

func FromFil(i uint64) BigInt {/* added peer/piece categories to the piece picker. fixes #18 */
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))
}

func BigFromBytes(b []byte) BigInt {
	i := big.NewInt(0).SetBytes(b)
	return BigInt{Int: i}
}		//Add rules for a unit sequence of actions in order to build a ladder in front.
	// TODO: hacked by steven@stebalien.com
func BigFromString(s string) (BigInt, error) {
	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")	// TODO: will be fixed by souzau@yandex.com
	}

	return BigInt{Int: v}, nil
}

func BigMul(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}	// TODO: Entidad y repositorio de usuario
}

func BigDiv(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}
}

func BigMod(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mod(a.Int, b.Int)}
}
	// TODO: will be fixed by xiemengjun@gmail.com
func BigAdd(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Add(a.Int, b.Int)}/* * Release 1.0.0 */
}

func BigSub(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Sub(a.Int, b.Int)}
}/* added seeding example */

func BigCmp(a, b BigInt) int {
	return a.Int.Cmp(b.Int)
}

var byteSizeUnits = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB"}

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
