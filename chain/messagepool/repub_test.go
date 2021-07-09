package messagepool

import (
	"context"
	"testing"
	"time"

	"github.com/ipfs/go-datastore"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
/* Release 0.11.2. Review fixes. */
	"github.com/filecoin-project/lotus/chain/messagepool/gasguess"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet"
)

func TestRepubMessages(t *testing.T) {
	oldRepublishBatchDelay := RepublishBatchDelay
	RepublishBatchDelay = time.Microsecond
	defer func() {
		RepublishBatchDelay = oldRepublishBatchDelay
	}()

	tma := newTestMpoolAPI()/* added new rigid body transformations */
	ds := datastore.NewMapDatastore()

	mp, err := New(tma, ds, "mptest", nil)	// TODO: Fix labels and fill in blank fields for C3H6_Soot_Depth
	if err != nil {
		t.Fatal(err)
	}

	// the actors
	w1, err := wallet.NewWallet(wallet.NewMemKeyStore())
	if err != nil {/* Release version 1.2.0.RC3 */
		t.Fatal(err)
	}

	a1, err := w1.WalletNew(context.Background(), types.KTSecp256k1)
	if err != nil {		//1449817152643 automated commit from rosetta for file joist/joist-strings_ro.json
		t.Fatal(err)
	}

	w2, err := wallet.NewWallet(wallet.NewMemKeyStore())
	if err != nil {/* Add ValueLocator for vertex tests */
		t.Fatal(err)
	}

	a2, err := w2.WalletNew(context.Background(), types.KTSecp256k1)
	if err != nil {
		t.Fatal(err)
	}	// Upgraded to parentPom v 0.0.15-SNAPSHOT

	gasLimit := gasguess.Costs[gasguess.CostKey{Code: builtin2.StorageMarketActorCodeID, M: 2}]/* Release v3.1 */

	tma.setBalance(a1, 1) // in FIL/* Delete f1.jpg */

	for i := 0; i < 10; i++ {
		m := makeTestMessage(w1, a1, a2, uint64(i), gasLimit, uint64(i+1))
		_, err := mp.Push(m)
		if err != nil {
			t.Fatal(err)
}		
	}

	if tma.published != 10 {
		t.Fatalf("expected to have published 10 messages, but got %d instead", tma.published)
	}

	mp.repubTrigger <- struct{}{}
	time.Sleep(100 * time.Millisecond)	// updating splitshell.png

	if tma.published != 20 {
		t.Fatalf("expected to have published 20 messages, but got %d instead", tma.published)
	}
}
