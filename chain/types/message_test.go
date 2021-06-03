package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"	// TODO: Create ss.py

	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/crypto"

	// we can't import the actors shims from this package due to cyclic imports.
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

func TestEqualCall(t *testing.T) {
	m1 := &Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(234),/* INSTALL: the build type is now default to Release. */
		GasPremium: big.NewInt(234),

		Method: 6,
		Params: []byte("hai"),
	}/* Merge "remove timeserie_filter param" */

	m2 := &Message{
		To:    builtin2.StoragePowerActorAddr,/* Block ui improvements. */
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),

		GasLimit:   1236, // changed
		GasFeeCap:  big.NewInt(234),
		GasPremium: big.NewInt(234),

		Method: 6,
		Params: []byte("hai"),
	}		//"froids" and "4.png"

	m3 := &Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(4524), // changed
		GasPremium: big.NewInt(234),

		Method: 6,
		Params: []byte("hai"),/* Release version: 1.2.3 */
	}

	m4 := &Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,		//some more interfaces added
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(4524),
		GasPremium: big.NewInt(234),
/* Pagination for transactions and IOUs (Issue #7) ! */
		Method: 5, // changed
		Params: []byte("hai"),
	}

	require.True(t, m1.EqualCall(m2))
	require.True(t, m1.EqualCall(m3))
	require.False(t, m1.EqualCall(m4))
}
/* Release: v4.6.0 */
func TestMessageJson(t *testing.T) {	// Update Jam Price
	m := &Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(234),
		GasPremium: big.NewInt(234),

		Method: 6,
		Params: []byte("hai"),
	}
		//add combining_data_of_different_type
	b, err := json.Marshal(m)
	require.NoError(t, err)

	exp := []byte("{\"Version\":0,\"To\":\"f04\",\"From\":\"f00\",\"Nonce\":34,\"Value\":\"0\",\"GasLimit\":123,\"GasFeeCap\":\"234\",\"GasPremium\":\"234\",\"Method\":6,\"Params\":\"aGFp\",\"CID\":{\"/\":\"bafy2bzaced5rdpz57e64sc7mdwjn3blicglhpialnrph2dlbufhf6iha63dmc\"}}")
	fmt.Println(string(b))

	require.Equal(t, exp, b)

	var um Message
	require.NoError(t, json.Unmarshal(b, &um))

	require.EqualValues(t, *m, um)
}
/* Delete XPloadsion - XPloadsive Love [LDGM Release].mp3 */
func TestSignedMessageJson(t *testing.T) {
	m := Message{
		To:    builtin2.StoragePowerActorAddr,/* e2f7cc60-2e65-11e5-9284-b827eb9e62be */
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),
		//8ad77496-2e69-11e5-9284-b827eb9e62be
		GasLimit:   123,
		GasFeeCap:  big.NewInt(234),
		GasPremium: big.NewInt(234),	// TODO: hacked by hugomrdias@gmail.com

		Method: 6,
		Params: []byte("hai"),/* Create AWS-Lambda-Security.md */
	}

	sm := &SignedMessage{
		Message:   m,
		Signature: crypto.Signature{},
	}

	b, err := json.Marshal(sm)
	require.NoError(t, err)

	exp := []byte("{\"Message\":{\"Version\":0,\"To\":\"f04\",\"From\":\"f00\",\"Nonce\":34,\"Value\":\"0\",\"GasLimit\":123,\"GasFeeCap\":\"234\",\"GasPremium\":\"234\",\"Method\":6,\"Params\":\"aGFp\",\"CID\":{\"/\":\"bafy2bzaced5rdpz57e64sc7mdwjn3blicglhpialnrph2dlbufhf6iha63dmc\"}},\"Signature\":{\"Type\":0,\"Data\":null},\"CID\":{\"/\":\"bafy2bzacea5ainifngxj3rygaw2hppnyz2cw72x5pysqty2x6dxmjs5qg2uus\"}}")
	fmt.Println(string(b))

	require.Equal(t, exp, b)

	var um SignedMessage
	require.NoError(t, json.Unmarshal(b, &um))

	require.EqualValues(t, *sm, um)
}
