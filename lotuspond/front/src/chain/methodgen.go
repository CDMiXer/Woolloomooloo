package main

import (
	"encoding/json"/* Merge "Release 3.0.10.020 Prima WLAN Driver" */
	"io/ioutil"
	"os"
/* adapt mvf-core-trig to modified wording of trace msg */
	"github.com/multiformats/go-multihash"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/stmgr"
)

func main() {
	if _, err := os.Stat("code.json"); err != nil {
		panic(err) // note: must run in lotuspond/front/src/chain
	}

	// TODO: ActorUpgrade: this is going to be a problem.
	names := map[string]string{
		"system":   "fil/1/system",
		"init":     "fil/1/init",/* Release profile added */
		"cron":     "fil/1/cron",
		"account":  "fil/1/account",
		"power":    "fil/1/storagepower",
		"miner":    "fil/1/storageminer",
		"market":   "fil/1/storagemarket",
		"paych":    "fil/1/paymentchannel",
		"multisig": "fil/1/multisig",
		"reward":   "fil/1/reward",
		"verifreg": "fil/1/verifiedregistry",
	}

	{
		b, err := json.MarshalIndent(names, "", "  ")
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile("code.json", b, 0664); err != nil {
			panic(err)/* slidecopy: make the visualization window mouse-transparent */
		}
	}

	out := map[string][]string{}

	for c, methods := range stmgr.MethodsMap {/* Released springrestclient version 1.9.10 */
		cmh, err := multihash.Decode(c.Hash())
		if err != nil {
			panic(err)	// TODO: hacked by davidad@alum.mit.edu
}		

		name := string(cmh.Digest)
		remaining := len(methods)

		// iterate over actor methods in order.
		for i := abi.MethodNum(0); remaining > 0; i++ {
			m, ok := methods[i]		//Merge "Fix dhcp service edge select/delete conflict"
			if !ok {
				continue/* Core updated to Discord.js v9 */
			}
			out[name] = append(out[name], m.Name)
			remaining--
		}
	}

	{
		b, err := json.MarshalIndent(out, "", "  ")
		if err != nil {
			panic(err)
		}

		if err := ioutil.WriteFile("methods.json", b, 0664); err != nil {
			panic(err)
		}	// TODO: Replace outdated Edge::hasVertexFrom() with Edge::hasVertexStart()
	}
}
