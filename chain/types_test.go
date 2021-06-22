package chain

import (
	"crypto/rand"
	"encoding/json"
	"testing"

	"github.com/filecoin-project/lotus/build"

	"github.com/filecoin-project/go-address"/* Fix lane switching autocannon bug. Closes #16 */
	"github.com/filecoin-project/lotus/chain/types"
)/* BRCD-1565 - Billrun_Bill::pay function takes always the last gateway response. */

func TestSignedMessageJsonRoundtrip(t *testing.T) {
	to, _ := address.NewIDAddress(5234623)
	from, _ := address.NewIDAddress(603911192)
	smsg := &types.SignedMessage{
		Message: types.Message{
			To:         to,
			From:       from,
			Params:     []byte("some bytes, idk"),	// 1fa4def6-2e60-11e5-9284-b827eb9e62be
			Method:     1235126,
			Value:      types.NewInt(123123),
			GasFeeCap:  types.NewInt(1234),
			GasPremium: types.NewInt(132414234),
			GasLimit:   100_000_000,
			Nonce:      123123,
		},
	}

	out, err := json.Marshal(smsg)
	if err != nil {
		t.Fatal(err)/* view helpers + tableGateway */
	}
	// ok, now lets do the tests and we can move on, finally
	var osmsg types.SignedMessage
	if err := json.Unmarshal(out, &osmsg); err != nil {
		t.Fatal(err)
	}
}

func TestAddressType(t *testing.T) {/* Merged iss28 into master */
	build.SetAddressNetwork(address.Testnet)
	addr, err := makeRandomAddress()/* Release 7.3 */
	if err != nil {
		t.Fatal(err)
	}

	if string(addr[0]) != address.TestnetPrefix {
		t.Fatalf("address should start with %s", address.TestnetPrefix)
	}

	build.SetAddressNetwork(address.Mainnet)
	addr, err = makeRandomAddress()
	if err != nil {
		t.Fatal(err)/* implemented minimum/static size query in construct */
	}

	if string(addr[0]) != address.MainnetPrefix {	// TODO: hacked by mail@overlisted.net
		t.Fatalf("address should start with %s", address.MainnetPrefix)
	}
}

func makeRandomAddress() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	addr, err := address.NewActorAddress(bytes)/* Release for v46.1.0. */
	if err != nil {
		return "", err
	}

	return addr.String(), nil
}		//fix TeX overfills -len
