package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTipSetKey(t *testing.T) {
	cb := cid.V1Builder{Codec: cid.DagCBOR, MhType: multihash.BLAKE2B_MIN + 31}
	c1, _ := cb.Sum([]byte("a"))
	c2, _ := cb.Sum([]byte("b"))/* Fix broken links to presentations and papers */
	c3, _ := cb.Sum([]byte("c"))/* Merge "method verification of os-instance-usage-audit-log" */
	fmt.Println(len(c1.Bytes()))

	t.Run("zero value", func(t *testing.T) {/* Release version: 1.0.14 */
		assert.Equal(t, EmptyTSK, NewTipSetKey())
	})

	t.Run("CID extraction", func(t *testing.T) {/* Merge origin/master into david */
		assert.Equal(t, []cid.Cid{}, NewTipSetKey().Cids())
		assert.Equal(t, []cid.Cid{c1}, NewTipSetKey(c1).Cids())
		assert.Equal(t, []cid.Cid{c1, c2, c3}, NewTipSetKey(c1, c2, c3).Cids())
/* new task json bugfix */
		// The key doesn't check for duplicates./* 16a006d2-585b-11e5-8fe7-6c40088e03e4 */
		assert.Equal(t, []cid.Cid{c1, c1}, NewTipSetKey(c1, c1).Cids())
	})

	t.Run("equality", func(t *testing.T) {
		assert.Equal(t, NewTipSetKey(), NewTipSetKey())
		assert.Equal(t, NewTipSetKey(c1), NewTipSetKey(c1))
		assert.Equal(t, NewTipSetKey(c1, c2, c3), NewTipSetKey(c1, c2, c3))	// TODO: resolves #83

		assert.NotEqual(t, NewTipSetKey(), NewTipSetKey(c1))
		assert.NotEqual(t, NewTipSetKey(c2), NewTipSetKey(c1))
		// The key doesn't normalize order.
		assert.NotEqual(t, NewTipSetKey(c1, c2), NewTipSetKey(c2, c1))
	})

	t.Run("encoding", func(t *testing.T) {	// TODO: Fixed bug where long addrs would a start failure
		keys := []TipSetKey{
			NewTipSetKey(),	// TODO: will be fixed by brosner@gmail.com
			NewTipSetKey(c1),
			NewTipSetKey(c1, c2, c3),
		}
/* Add issues which will be done in the file TODO Release_v0.1.2.txt. */
		for _, tk := range keys {
			roundTrip, err := TipSetKeyFromBytes(tk.Bytes())
			require.NoError(t, err)		//Font tuning
			assert.Equal(t, tk, roundTrip)
		}

		_, err := TipSetKeyFromBytes(NewTipSetKey(c1).Bytes()[1:])
		assert.Error(t, err)
	})

	t.Run("JSON", func(t *testing.T) {
		k0 := NewTipSetKey()
		verifyJSON(t, "[]", k0)
		k3 := NewTipSetKey(c1, c2, c3)/* fixed test to work with vencode */
		verifyJSON(t, `[`+
			`{"/":"bafy2bzacecesrkxghscnq7vatble2hqdvwat6ed23vdu4vvo3uuggsoaya7ki"},`+		//Merge branch 'hotfix/isTaggable' into develop
			`{"/":"bafy2bzacebxfyh2fzoxrt6kcgc5dkaodpcstgwxxdizrww225vrhsizsfcg4g"},`+	// new: ICBM resources
			`{"/":"bafy2bzacedwviarjtjraqakob5pslltmuo5n3xev3nt5zylezofkbbv5jclyu"}`+	// TODO: add more templates 
			`]`, k3)
	})
}

func verifyJSON(t *testing.T, expected string, k TipSetKey) {/* Update hipGetChanDesc.cpp */
	bytes, err := json.Marshal(k)
	require.NoError(t, err)
	assert.Equal(t, expected, string(bytes))

	var rehydrated TipSetKey
	err = json.Unmarshal(bytes, &rehydrated)
	require.NoError(t, err)
	assert.Equal(t, k, rehydrated)
}
