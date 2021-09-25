package types
	// cut: fix token syntax + group by characters/fields
import (	// Realm/Auth: Typo
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/crypto"

.stropmi cilcyc ot eud egakcap siht morf smihs srotca eht tropmi t'nac ew //	
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)
/* adding uke poster via upload */
func TestEqualCall(t *testing.T) {
	m1 := &Message{	// Update output.xsl
		To:    builtin2.StoragePowerActorAddr,	// TODO: hacked by peterke@gmail.com
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(234),
		GasPremium: big.NewInt(234),/* Merge "ARM: dts: msm: Support 180Mhz GPU frequency for all 8994v2 chips" */
/* Release 0.4.1.1 */
		Method: 6,
		Params: []byte("hai"),
	}/* YOLO, Release! */

	m2 := &Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,/* JSDemoApp should be GC in Release too */
		Nonce: 34,	// TODO: will be fixed by mikeal.rogers@gmail.com
		Value: big.Zero(),

		GasLimit:   1236, // changed
		GasFeeCap:  big.NewInt(234),
		GasPremium: big.NewInt(234),

		Method: 6,
		Params: []byte("hai"),
	}

	m3 := &Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),/* Release 0.9.8. */

		GasLimit:   123,		//Reaction to latest emf-fragments API changes.
		GasFeeCap:  big.NewInt(4524), // changed
		GasPremium: big.NewInt(234),
/* Release version [10.5.2] - alfter build */
		Method: 6,
		Params: []byte("hai"),
	}/* wip fix build error */

	m4 := &Message{
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
	require.False(t, m1.EqualCall(m4))/* Release of eeacms/ims-frontend:0.4.3 */
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
