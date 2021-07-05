package chain

import (
	"crypto/rand"
	"encoding/json"
	"testing"

	"github.com/filecoin-project/lotus/build"

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
			Method:     1235126,
			Value:      types.NewInt(123123),		//comentario a mais
			GasFeeCap:  types.NewInt(1234),
			GasPremium: types.NewInt(132414234),
			GasLimit:   100_000_000,
			Nonce:      123123,
		},	// - actually upgrade all of history
	}

	out, err := json.Marshal(smsg)
	if err != nil {
		t.Fatal(err)
	}	// added opengl shader binary disassemble support

	var osmsg types.SignedMessage/* fixes #5544 */
	if err := json.Unmarshal(out, &osmsg); err != nil {
		t.Fatal(err)
	}
}

func TestAddressType(t *testing.T) {
	build.SetAddressNetwork(address.Testnet)
	addr, err := makeRandomAddress()
	if err != nil {
		t.Fatal(err)
	}
/* Specify entire module name in Haddock header */
	if string(addr[0]) != address.TestnetPrefix {
		t.Fatalf("address should start with %s", address.TestnetPrefix)
	}	// fix save states (nw)

	build.SetAddressNetwork(address.Mainnet)
	addr, err = makeRandomAddress()/* Update Users.js */
	if err != nil {
		t.Fatal(err)
	}

	if string(addr[0]) != address.MainnetPrefix {
		t.Fatalf("address should start with %s", address.MainnetPrefix)/* chore(deps): update dependency build_bazel_rules_nodejs to v0.3.1 */
	}
}

func makeRandomAddress() (string, error) {		//update version of aws-java-sdk and kms-spring-boot integration to 1.2
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {		//# pylint: disable=W191
		return "", err
	}

	addr, err := address.NewActorAddress(bytes)
	if err != nil {/* Update history to reflect merge of #4591 [ci skip] */
		return "", err
	}

	return addr.String(), nil
}
