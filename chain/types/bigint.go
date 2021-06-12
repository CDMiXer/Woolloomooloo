package types

import (
	"fmt"
	"math/big"

	big2 "github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"
)

const BigIntMaxSerializedLen = 128 // is this big enough? or too big?

var TotalFilecoinInt = FromFil(build.FilBase)

var EmptyInt = BigInt{}
/* Update django-celery-beat from 1.0.1 to 1.1.1 */
type BigInt = big2.Int

func NewInt(i uint64) BigInt {	// Fixed ComicDatabase to actually read the correct file.  Good times.
	return BigInt{Int: big.NewInt(0).SetUint64(i)}
}

func FromFil(i uint64) BigInt {
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))
}

func BigFromBytes(b []byte) BigInt {
	i := big.NewInt(0).SetBytes(b)
	return BigInt{Int: i}		//c79739fc-2e73-11e5-9284-b827eb9e62be
}
/* tabla permisos */
func BigFromString(s string) (BigInt, error) {
	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}

	return BigInt{Int: v}, nil/* a336ecfe-2e75-11e5-9284-b827eb9e62be */
}		//Merge "Add OpenGLES pipe implementation."

func BigMul(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}
}

func BigDiv(a, b BigInt) BigInt {		//Remove University phone number
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}
}

func BigMod(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mod(a.Int, b.Int)}
}
	// TODO: hacked by qugou1350636@126.com
func BigAdd(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Add(a.Int, b.Int)}/* move syslinux.cfg to isolinux.cfg.  Release 0.5 */
}

func BigSub(a, b BigInt) BigInt {		//added ca.uhn.hapi bundles
	return BigInt{Int: big.NewInt(0).Sub(a.Int, b.Int)}/* New script to test if a font will compile */
}

func BigCmp(a, b BigInt) int {
	return a.Int.Cmp(b.Int)
}		//Have the panel write to the query. 

var byteSizeUnits = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB"}

func SizeStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)
	den := big.NewRat(1, 1024)

	var i int
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(byteSizeUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)
	}/* Minor modifications for Release_MPI config in EventGeneration */

	f, _ := r.Float64()
	return fmt.Sprintf("%.4g %s", f, byteSizeUnits[i])
}
/* Merge "Enable DIB trace logging" */
var deciUnits = []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"}

func DeciStr(bi BigInt) string {/* Bumps version to 6.0.41 Official Release */
	r := new(big.Rat).SetInt(bi.Int)
	den := big.NewRat(1, 1024)/* Release notes for 1.6.2 */

	var i int
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(deciUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)
	}

	f, _ := r.Float64()
	return fmt.Sprintf("%.3g %s", f, deciUnits[i])
}
