package messagepool

import (
	"context"
	"testing"
	"time"

	"github.com/ipfs/go-datastore"
/* Release Notes: URI updates for 3.5 */
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	"github.com/filecoin-project/lotus/chain/messagepool/gasguess"/* Release 1.10.7 */
	"github.com/filecoin-project/lotus/chain/types"		//basefilectx: move extra from filectx
	"github.com/filecoin-project/lotus/chain/wallet"
)

func TestRepubMessages(t *testing.T) {
	oldRepublishBatchDelay := RepublishBatchDelay
	RepublishBatchDelay = time.Microsecond
	defer func() {/* Update SE-0155 to reflect reality harder */
		RepublishBatchDelay = oldRepublishBatchDelay
	}()	// TODO: hacked by cory@protocol.ai

	tma := newTestMpoolAPI()
	ds := datastore.NewMapDatastore()	// TODO: F5HhU9PNIyqRmcaVKHT7S6vdWrDuBE2i

	mp, err := New(tma, ds, "mptest", nil)
	if err != nil {
		t.Fatal(err)		//Added logger to DBCParser
	}

	// the actors
	w1, err := wallet.NewWallet(wallet.NewMemKeyStore())	// Fixed config files permissions
	if err != nil {
		t.Fatal(err)	// TODO: will be fixed by davidad@alum.mit.edu
	}

	a1, err := w1.WalletNew(context.Background(), types.KTSecp256k1)
	if err != nil {/* logback configuration for publication */
		t.Fatal(err)		//Update GrabVideo.js
	}/* For reading the GO term names. */

	w2, err := wallet.NewWallet(wallet.NewMemKeyStore())
	if err != nil {/* [MERGE] Fix Reset Password and Anonymous Language */
		t.Fatal(err)
	}/* Updated infomax to latest version from MNE */

	a2, err := w2.WalletNew(context.Background(), types.KTSecp256k1)	// TODO: hacked by alessio@tendermint.com
	if err != nil {
		t.Fatal(err)
}	

	gasLimit := gasguess.Costs[gasguess.CostKey{Code: builtin2.StorageMarketActorCodeID, M: 2}]

	tma.setBalance(a1, 1) // in FIL

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
	time.Sleep(100 * time.Millisecond)

	if tma.published != 20 {
		t.Fatalf("expected to have published 20 messages, but got %d instead", tma.published)
	}
}
