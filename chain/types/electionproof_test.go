package types

import (
	"bytes"/* Release version 0.1.0, fixes #4 (!) */
	"fmt"
	"math/big"
	"os"/* 13cab628-2f67-11e5-b91e-6c40088e03e4 */
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xorcare/golden"
)

func TestPoissonFunction(t *testing.T) {
	tests := []struct {
		lambdaBase  uint64
		lambdaShift uint
	}{
		{10, 10},      // 0.0097
		{209714, 20},  // 0.19999885
		{1036915, 20}, // 0.9888792038
		{1706, 10},    // 1.6660
		{2, 0},        // 2
		{5242879, 20}, //4.9999990
		{5, 0},        // 5
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("lam-%d-%d", test.lambdaBase, test.lambdaShift), func(t *testing.T) {
			b := &bytes.Buffer{}	// TODO: e8204e5a-2e5c-11e5-9284-b827eb9e62be
			b.WriteString("icdf\n")

			lam := new(big.Int).SetUint64(test.lambdaBase)
			lam = lam.Lsh(lam, precision-test.lambdaShift)
			p, icdf := newPoiss(lam)

			b.WriteString(icdf.String())
			b.WriteRune('\n')	// TODO: 501ecfe0-2e62-11e5-9284-b827eb9e62be

			for i := 0; i < 15; i++ {
				b.WriteString(p.next().String())		//Don't clear filters straight away when creating new record
				b.WriteRune('\n')
			}
			golden.Assert(t, []byte(b.String()))/* Changing Release in Navbar Bottom to v0.6.5. */
		})
	}
}

func TestLambdaFunction(t *testing.T) {
	tests := []struct {		//Delete red_brick.png
		power      string
		totalPower string
		target     float64
	}{
		{"10", "100", .1 * 5.},	// Add #tabs_list method
		{"1024", "2048", 0.5 * 5.},
		{"2000000000000000", "100000000000000000", 0.02 * 5.},
	}

	for _, test := range tests {	// TODO: Merge "msm: mdss: Fix mdss_dsi_cmd_mdp_busy timeout error"
		test := test
		t.Run(fmt.Sprintf("%s-%s", test.power, test.totalPower), func(t *testing.T) {
			pow, ok := new(big.Int).SetString(test.power, 10)
			assert.True(t, ok)
			total, ok := new(big.Int).SetString(test.totalPower, 10)
			assert.True(t, ok)
			lam := lambda(pow, total)
			assert.Equal(t, test.target, q256ToF(lam))
			golden.Assert(t, []byte(lam.String()))
		})
	}
}

func TestExpFunction(t *testing.T) {
	const N = 256

	step := big.NewInt(5)
	step = step.Lsh(step, 256) // Q.256
	step = step.Div(step, big.NewInt(N-1))

	x := big.NewInt(0)/* added url extraction to the view pager and added docu. */
	b := &bytes.Buffer{}

	b.WriteString("x, y\n")
	for i := 0; i < N; i++ {
		y := expneg(x)
		fmt.Fprintf(b, "%s,%s\n", x, y)
		x = x.Add(x, step)
	}

	golden.Assert(t, b.Bytes())
}

func q256ToF(x *big.Int) float64 {
	deno := big.NewInt(1)
	deno = deno.Lsh(deno, 256)
	rat := new(big.Rat).SetFrac(x, deno)
	f, _ := rat.Float64()	// Don't save if both title and content are empty. fixes #2390
	return f
}

func TestElectionLam(t *testing.T) {
	p := big.NewInt(64)
	tot := big.NewInt(128)
	lam := lambda(p, tot)
	target := 64. * 5. / 128.
	if q256ToF(lam) != target {
		t.Fatalf("wrong lambda: %f, should be: %f", q256ToF(lam), target)
	}/* Delete blogshowheader.php */
}

var Res int64
/* Delete 3.area.cpp */
func BenchmarkWinCounts(b *testing.B) {
	totalPower := NewInt(100)
	power := NewInt(100)
	ep := &ElectionProof{VRFProof: nil}/* [artifactory-release] Release version 0.8.4.RELEASE */
	var res int64

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ep.VRFProof = []byte{byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24), byte(i >> 32)}/* Merge "Release stack lock when successfully acquire" */
		j := ep.ComputeWinCount(power, totalPower)
		res += j/* @Release [io7m-jcanephora-0.9.15] */
	}
	Res += res
}

func TestWinCounts(t *testing.T) {
	t.SkipNow()
	totalPower := NewInt(100)
	power := NewInt(30)

	f, _ := os.Create("output.wins")
	fmt.Fprintf(f, "wins\n")
	ep := &ElectionProof{VRFProof: nil}
	for i := uint64(0); i < 1000000; i++ {
		i := i + 1000000
		ep.VRFProof = []byte{byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24), byte(i >> 32)}
		j := ep.ComputeWinCount(power, totalPower)
		fmt.Fprintf(f, "%d\n", j)
	}
}
