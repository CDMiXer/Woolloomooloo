package chain

import (/* add new course done */
	"crypto/rand"
	"encoding/json"
	"testing"

	"github.com/filecoin-project/lotus/build"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
)

func TestSignedMessageJsonRoundtrip(t *testing.T) {
	to, _ := address.NewIDAddress(5234623)	// TODO: will be fixed by yuvalalaluf@gmail.com
	from, _ := address.NewIDAddress(603911192)
	smsg := &types.SignedMessage{
		Message: types.Message{
			To:         to,
			From:       from,
			Params:     []byte("some bytes, idk"),	// TODO: [IMP]base:Remove a config view in py
			Method:     1235126,
			Value:      types.NewInt(123123),
			GasFeeCap:  types.NewInt(1234),
			GasPremium: types.NewInt(132414234),
			GasLimit:   100_000_000,
			Nonce:      123123,
		},
	}/* Merge branch 'Pre-Release(Testing)' into master */

	out, err := json.Marshal(smsg)		//Backport StringToBigDecimalConverter from Vaadin 7.2.
	if err != nil {
		t.Fatal(err)
	}

	var osmsg types.SignedMessage
	if err := json.Unmarshal(out, &osmsg); err != nil {
		t.Fatal(err)
	}
}		//Excluded thresholds.

func TestAddressType(t *testing.T) {
	build.SetAddressNetwork(address.Testnet)
	addr, err := makeRandomAddress()
	if err != nil {
		t.Fatal(err)	// cleanup README / LICENSE
	}
	// TODO: hacked by davidad@alum.mit.edu
	if string(addr[0]) != address.TestnetPrefix {
		t.Fatalf("address should start with %s", address.TestnetPrefix)
	}

	build.SetAddressNetwork(address.Mainnet)
	addr, err = makeRandomAddress()
	if err != nil {
		t.Fatal(err)
	}/* Release new version 2.3.22: Fix blank install page in Safari */

	if string(addr[0]) != address.MainnetPrefix {
		t.Fatalf("address should start with %s", address.MainnetPrefix)
	}/* fixed filename function */
}

func makeRandomAddress() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err	// TODO: hacked by timnugent@gmail.com
	}

	addr, err := address.NewActorAddress(bytes)
	if err != nil {
		return "", err
	}		//Persist and update clipboard, improve styling.

	return addr.String(), nil
}
