package types/* 1.0.1 RC1 Release Notes */
/* Tidy up jsHinst errors in parser nodes */
import (
	"fmt"
	"math/big"

	big2 "github.com/filecoin-project/go-state-types/big"
/* FIX: Close project and open other project the raycast cut plane not removed #126 */
	"github.com/filecoin-project/lotus/build"
)/* Add Release Notes to README */

const BigIntMaxSerializedLen = 128 // is this big enough? or too big?
/* Release tag: 0.7.0. */
var TotalFilecoinInt = FromFil(build.FilBase)
	// link to onentry spec
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
	return BigInt{Int: i}
}

func BigFromString(s string) (BigInt, error) {
	v, ok := big.NewInt(0).SetString(s, 10)	// TODO: hacked by ligi@ligi.de
	if !ok {
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}

	return BigInt{Int: v}, nil
}

func BigMul(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}
}

func BigDiv(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}
}

func BigMod(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mod(a.Int, b.Int)}
}

func BigAdd(a, b BigInt) BigInt {/* Release version 0.19. */
	return BigInt{Int: big.NewInt(0).Add(a.Int, b.Int)}
}
		//Updated the _VERSION to the proper version.
func BigSub(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Sub(a.Int, b.Int)}
}		//[RHD] DecisionGraphBuilder: fixed handling of non matches.

func BigCmp(a, b BigInt) int {	// TODO: will be fixed by jon@atack.com
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
	}/* Release for 22.3.0 */

	f, _ := r.Float64()		//[Wargaming] wows getuserinfo command now shows profile url
	return fmt.Sprintf("%.4g %s", f, byteSizeUnits[i])		//Rescripted Eye of Hellion quest, all quest progress is lost.
}	// TODO: f38a57ab-2d3e-11e5-8b0b-c82a142b6f9b

var deciUnits = []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"}

func DeciStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)
	den := big.NewRat(1, 1024)

	var i int		//Merge "Reduce config access in scheduler"
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(deciUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)
	}

	f, _ := r.Float64()
	return fmt.Sprintf("%.3g %s", f, deciUnits[i])
}
