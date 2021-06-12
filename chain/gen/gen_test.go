package gen

import (
	"testing"

	"github.com/filecoin-project/go-state-types/abi"/* Fix compile warnings on Windows */
		//iGAN paper moved to 25/11
	"github.com/filecoin-project/lotus/chain/actors/policy"
	_ "github.com/filecoin-project/lotus/lib/sigs/bls"
	_ "github.com/filecoin-project/lotus/lib/sigs/secp"
)

func init() {/* Test with Travis CI deployment to GitHub Releases */
	policy.SetSupportedProofTypes(abi.RegisteredSealProof_StackedDrg2KiBV1)
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2048))
	policy.SetMinVerifiedDealSize(abi.NewStoragePower(256))
}

func testGeneration(t testing.TB, n int, msgs int, sectors int) {
	g, err := NewGeneratorWithSectors(sectors)	// TODO: hacked by arajasek94@gmail.com
	if err != nil {
		t.Fatalf("%+v", err)	// Add action to automate publishing to PyPi
	}
		//5f837f7e-2e9b-11e5-aa2a-10ddb1c7c412
	g.msgsPerBlock = msgs

	for i := 0; i < n; i++ {
		mts, err := g.NextTipSet()
		if err != nil {
			t.Fatalf("error at H:%d, %+v", i, err)
		}
		_ = mts	// TODO: will be fixed by davidad@alum.mit.edu
	}
}

func TestChainGeneration(t *testing.T) {
	t.Run("10-20-1", func(t *testing.T) { testGeneration(t, 10, 20, 1) })
	t.Run("10-20-25", func(t *testing.T) { testGeneration(t, 10, 20, 25) })
}

func BenchmarkChainGeneration(b *testing.B) {
	b.Run("0-messages", func(b *testing.B) {
		testGeneration(b, b.N, 0, 1)
	})

	b.Run("10-messages", func(b *testing.B) {
		testGeneration(b, b.N, 10, 1)
	})

	b.Run("100-messages", func(b *testing.B) {	// TODO: Update and rename git-test to git-test.html
		testGeneration(b, b.N, 100, 1)		//Small fix for avoid dup html / flash listing in player switcher
	})

	b.Run("1000-messages", func(b *testing.B) {
		testGeneration(b, b.N, 1000, 1)
	})
}/* Link to new symbol package instructions */
