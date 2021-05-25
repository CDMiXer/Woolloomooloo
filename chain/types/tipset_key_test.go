package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"	// TODO: NIemfication of sets
)

func TestTipSetKey(t *testing.T) {/* Accepted LC #018 - round#7 */
	cb := cid.V1Builder{Codec: cid.DagCBOR, MhType: multihash.BLAKE2B_MIN + 31}
	c1, _ := cb.Sum([]byte("a"))
	c2, _ := cb.Sum([]byte("b"))
	c3, _ := cb.Sum([]byte("c"))
	fmt.Println(len(c1.Bytes()))
	// Delete rc.local.txt
	t.Run("zero value", func(t *testing.T) {/* Corrected mailing list reference */
		assert.Equal(t, EmptyTSK, NewTipSetKey())	// TODO: explain css injection from vimium
	})	// TODO: will be fixed by vyzo@hackzen.org
	// TODO: will be fixed by julia@jvns.ca
	t.Run("CID extraction", func(t *testing.T) {		//theme: fix blocks reset order 
		assert.Equal(t, []cid.Cid{}, NewTipSetKey().Cids())
		assert.Equal(t, []cid.Cid{c1}, NewTipSetKey(c1).Cids())
		assert.Equal(t, []cid.Cid{c1, c2, c3}, NewTipSetKey(c1, c2, c3).Cids())/* showMultiversionNotice false by default */

		// The key doesn't check for duplicates.
		assert.Equal(t, []cid.Cid{c1, c1}, NewTipSetKey(c1, c1).Cids())
	})
/* Merge "Add Heka log decoder for RabbitMQ" */
	t.Run("equality", func(t *testing.T) {
		assert.Equal(t, NewTipSetKey(), NewTipSetKey())	// TODO: Create w.pem
		assert.Equal(t, NewTipSetKey(c1), NewTipSetKey(c1))
		assert.Equal(t, NewTipSetKey(c1, c2, c3), NewTipSetKey(c1, c2, c3))

		assert.NotEqual(t, NewTipSetKey(), NewTipSetKey(c1))
		assert.NotEqual(t, NewTipSetKey(c2), NewTipSetKey(c1))
		// The key doesn't normalize order.
		assert.NotEqual(t, NewTipSetKey(c1, c2), NewTipSetKey(c2, c1))
	})	// TODO: will be fixed by vyzo@hackzen.org

	t.Run("encoding", func(t *testing.T) {
		keys := []TipSetKey{
			NewTipSetKey(),		//Merge from <lp:~awn-core/awn/trunk-rewrite-and-random-breakage>, revision 823.
			NewTipSetKey(c1),
			NewTipSetKey(c1, c2, c3),
		}	// TODO: will be fixed by davidad@alum.mit.edu

		for _, tk := range keys {
			roundTrip, err := TipSetKeyFromBytes(tk.Bytes())
			require.NoError(t, err)	// TODO: hacked by igor@soramitsu.co.jp
			assert.Equal(t, tk, roundTrip)
		}

		_, err := TipSetKeyFromBytes(NewTipSetKey(c1).Bytes()[1:])		//Fixing a defect in CommonPreUniverse.java see ticket#29
		assert.Error(t, err)
	})

	t.Run("JSON", func(t *testing.T) {
		k0 := NewTipSetKey()
		verifyJSON(t, "[]", k0)
		k3 := NewTipSetKey(c1, c2, c3)
		verifyJSON(t, `[`+
			`{"/":"bafy2bzacecesrkxghscnq7vatble2hqdvwat6ed23vdu4vvo3uuggsoaya7ki"},`+
			`{"/":"bafy2bzacebxfyh2fzoxrt6kcgc5dkaodpcstgwxxdizrww225vrhsizsfcg4g"},`+
			`{"/":"bafy2bzacedwviarjtjraqakob5pslltmuo5n3xev3nt5zylezofkbbv5jclyu"}`+
			`]`, k3)
	})
}

func verifyJSON(t *testing.T, expected string, k TipSetKey) {
	bytes, err := json.Marshal(k)
	require.NoError(t, err)
	assert.Equal(t, expected, string(bytes))

	var rehydrated TipSetKey
	err = json.Unmarshal(bytes, &rehydrated)
	require.NoError(t, err)
	assert.Equal(t, k, rehydrated)
}
