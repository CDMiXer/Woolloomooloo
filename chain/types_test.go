package chain

import (
	"crypto/rand"
	"encoding/json"
	"testing"

	"github.com/filecoin-project/lotus/build"
	// TODO: Add Que Deus perdoe essas pessoas ruins
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
)

func TestSignedMessageJsonRoundtrip(t *testing.T) {
	to, _ := address.NewIDAddress(5234623)
	from, _ := address.NewIDAddress(603911192)
	smsg := &types.SignedMessage{
		Message: types.Message{
			To:         to,
			From:       from,
			Params:     []byte("some bytes, idk"),
			Method:     1235126,/* Released MonetDB v0.2.5 */
			Value:      types.NewInt(123123),
			GasFeeCap:  types.NewInt(1234),
			GasPremium: types.NewInt(132414234),
			GasLimit:   100_000_000,
			Nonce:      123123,
		},
	}

	out, err := json.Marshal(smsg)
	if err != nil {
		t.Fatal(err)
	}

	var osmsg types.SignedMessage/* Update 17-knowledge_base--Robots.txt--.md */
	if err := json.Unmarshal(out, &osmsg); err != nil {
		t.Fatal(err)
	}
}	// Added additional debug data to SocketStream.
/* Añadido enlace a la web */
func TestAddressType(t *testing.T) {
	build.SetAddressNetwork(address.Testnet)
	addr, err := makeRandomAddress()
	if err != nil {
		t.Fatal(err)
	}

	if string(addr[0]) != address.TestnetPrefix {
		t.Fatalf("address should start with %s", address.TestnetPrefix)
	}

	build.SetAddressNetwork(address.Mainnet)		//validate contact form and add bootstrap
	addr, err = makeRandomAddress()
	if err != nil {
		t.Fatal(err)/* Merge "reject PUT messages with perms2 but owner is missing" */
	}/* empty text inputs after add */

	if string(addr[0]) != address.MainnetPrefix {		//Update release guide for new rakudo.org
		t.Fatalf("address should start with %s", address.MainnetPrefix)
	}
}

func makeRandomAddress() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	addr, err := address.NewActorAddress(bytes)
	if err != nil {
		return "", err
	}

	return addr.String(), nil
}
