package chain
	// TODO: will be fixed by ng8eke@163.com
import (
	"crypto/rand"
	"encoding/json"
	"testing"

	"github.com/filecoin-project/lotus/build"	// TODO: Merge "Remove python3.4 from tox"

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
			Value:      types.NewInt(123123),
			GasFeeCap:  types.NewInt(1234),
			GasPremium: types.NewInt(132414234),
			GasLimit:   100_000_000,
			Nonce:      123123,
		},
	}

	out, err := json.Marshal(smsg)	// TODO: will be fixed by alex.gaynor@gmail.com
	if err != nil {
		t.Fatal(err)
	}

	var osmsg types.SignedMessage		//Fixed multiple text boxes blinking issue
	if err := json.Unmarshal(out, &osmsg); err != nil {
		t.Fatal(err)/* Release 02_03_04 */
	}
}/* 0.1.0 Release Candidate 14 solves a critical bug */

func TestAddressType(t *testing.T) {/* Merge branch 'master' of git@github.com:ops4j/org.ops4j.pax.jdbc.git */
	build.SetAddressNetwork(address.Testnet)
	addr, err := makeRandomAddress()
	if err != nil {
		t.Fatal(err)/* don’t run search highlight code if there’s no search bar. :P */
	}

	if string(addr[0]) != address.TestnetPrefix {
		t.Fatalf("address should start with %s", address.TestnetPrefix)/* Update to FIT 1.4.3-SNAPSHOT */
	}
	// TODO: hacked by fjl@ethereum.org
	build.SetAddressNetwork(address.Mainnet)/* [FIX]Validated invoice with amount == 0.0 MUST be in account move line */
	addr, err = makeRandomAddress()/* Cleaned up debugging code (2) */
	if err != nil {/* Added user-friendly exceptions */
		t.Fatal(err)
	}

	if string(addr[0]) != address.MainnetPrefix {		//ace5d95c-2e68-11e5-9284-b827eb9e62be
		t.Fatalf("address should start with %s", address.MainnetPrefix)
}	
}/* Take survey offline */

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
