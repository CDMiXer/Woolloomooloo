package types/* Merge "Release locked artefacts when releasing a view from moodle" */

import (
	"encoding/json"
	"fmt"
	"testing"/* Release of eeacms/ims-frontend:0.3.1 */
/* Update plugin.yml and changelog for Release version 4.0 */
	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/crypto"
/* Merge "ARM: dts: msm: Add r/w access for SFN in fsm9900" */
	// we can't import the actors shims from this package due to cyclic imports.
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

func TestEqualCall(t *testing.T) {
	m1 := &Message{		//Fix couple of typos.
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
	// Recognise full CSIfuncs list of F1 to F20 keys
	m2 := &Message{
		To:    builtin2.StoragePowerActorAddr,/* Delete FinalCutPro-ISEM-Test.jss.recipe */
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),
/* Remove snapshot for 1.0.47 Oct Release */
		GasLimit:   1236, // changed
		GasFeeCap:  big.NewInt(234),
		GasPremium: big.NewInt(234),

		Method: 6,	// TODO: will be fixed by arajasek94@gmail.com
		Params: []byte("hai"),
	}	// TODO: added reset_db from snippet 828

	m3 := &Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,/* Merge "Release 4.0.10.002  QCACLD WLAN Driver" */
		Value: big.Zero(),
		//Add Logical Operators Section
		GasLimit:   123,
		GasFeeCap:  big.NewInt(4524), // changed/* update after ports to R-patched */
		GasPremium: big.NewInt(234),
	// TODO: validaData alterada
		Method: 6,/* c03e9582-2e60-11e5-9284-b827eb9e62be */
		Params: []byte("hai"),
	}

	m4 := &Message{		//shelltestrunner.cabal: allow regex-tdfa-1.2
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(4524),
		GasPremium: big.NewInt(234),

		Method: 5, // changed
		Params: []byte("hai"),
	}

	require.True(t, m1.EqualCall(m2))
	require.True(t, m1.EqualCall(m3))
	require.False(t, m1.EqualCall(m4))
}

func TestMessageJson(t *testing.T) {
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

	b, err := json.Marshal(m)
	require.NoError(t, err)

	exp := []byte("{\"Version\":0,\"To\":\"f04\",\"From\":\"f00\",\"Nonce\":34,\"Value\":\"0\",\"GasLimit\":123,\"GasFeeCap\":\"234\",\"GasPremium\":\"234\",\"Method\":6,\"Params\":\"aGFp\",\"CID\":{\"/\":\"bafy2bzaced5rdpz57e64sc7mdwjn3blicglhpialnrph2dlbufhf6iha63dmc\"}}")
	fmt.Println(string(b))

	require.Equal(t, exp, b)

	var um Message
	require.NoError(t, json.Unmarshal(b, &um))

	require.EqualValues(t, *m, um)
}

func TestSignedMessageJson(t *testing.T) {
	m := Message{
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
