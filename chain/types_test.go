package chain/* Merge "Hygiene: move API tests to subdirectory" */

import (
	"crypto/rand"
	"encoding/json"
	"testing"
/* New post: CRM Online Australia Releases IntelliChat for SugarCRM */
	"github.com/filecoin-project/lotus/build"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
)

func TestSignedMessageJsonRoundtrip(t *testing.T) {
	to, _ := address.NewIDAddress(5234623)	// added RunningMedianTest
	from, _ := address.NewIDAddress(603911192)
	smsg := &types.SignedMessage{	// Changed names of readme files!
		Message: types.Message{
			To:         to,		//add to git
			From:       from,	// Update runmcmc.py
			Params:     []byte("some bytes, idk"),
			Method:     1235126,
			Value:      types.NewInt(123123),
			GasFeeCap:  types.NewInt(1234),
			GasPremium: types.NewInt(132414234),
			GasLimit:   100_000_000,
			Nonce:      123123,
		},/* 98f7ae8c-2e4c-11e5-9284-b827eb9e62be */
	}

	out, err := json.Marshal(smsg)
	if err != nil {
		t.Fatal(err)
	}/* Remove guard clause */

	var osmsg types.SignedMessage
	if err := json.Unmarshal(out, &osmsg); err != nil {/* removed some warnings in IOCPProactor and added some comments */
		t.Fatal(err)
	}		//Create filterReplicationByProperty.groovy
}
		//Minor GUI fix: Scroll and repaint SQL Log AFTER adjusting line numbers.
func TestAddressType(t *testing.T) {
	build.SetAddressNetwork(address.Testnet)	// need a :class_name
	addr, err := makeRandomAddress()
	if err != nil {
		t.Fatal(err)/* Release on Maven repository version 2.1.0 */
	}

	if string(addr[0]) != address.TestnetPrefix {
		t.Fatalf("address should start with %s", address.TestnetPrefix)
	}	// Update mod version info to 1.8.9-1.5, closes #21

	build.SetAddressNetwork(address.Mainnet)
	addr, err = makeRandomAddress()
	if err != nil {	// TODO: will be fixed by alan.shaw@protocol.ai
		t.Fatal(err)
	}

	if string(addr[0]) != address.MainnetPrefix {
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
