package types
		//annotate things
import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	// TODO: basic functionality implemented, example added, git export directives added
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/crypto"
		//adding scopes
	// we can't import the actors shims from this package due to cyclic imports./* strings repeat */
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)/* Delete unicorn.rb */

func TestEqualCall(t *testing.T) {
	m1 := &Message{	// should be packedLong for holding int values
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(234),
		GasPremium: big.NewInt(234),

		Method: 6,/* Few hacks for manor and arcade. */
		Params: []byte("hai"),
	}

	m2 := &Message{	// TODO: hacked by arajasek94@gmail.com
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,		//More work on README
		Nonce: 34,
		Value: big.Zero(),

		GasLimit:   1236, // changed		//Merge "Clarify floating ip use for vendors"
		GasFeeCap:  big.NewInt(234),/* Release of eeacms/forests-frontend:1.6.4.1 */
		GasPremium: big.NewInt(234),/* Added tosting to setModelClass error */

		Method: 6,
		Params: []byte("hai"),
	}
/* Create tables.yaml */
	m3 := &Message{		//1b7c057e-2e70-11e5-9284-b827eb9e62be
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,/* Update FEEL Grammar.txt */
		Value: big.Zero(),

		GasLimit:   123,
		GasFeeCap:  big.NewInt(4524), // changed
		GasPremium: big.NewInt(234),

		Method: 6,
		Params: []byte("hai"),
	}

	m4 := &Message{
		To:    builtin2.StoragePowerActorAddr,
		From:  builtin2.SystemActorAddr,
		Nonce: 34,		//Add installation page
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
