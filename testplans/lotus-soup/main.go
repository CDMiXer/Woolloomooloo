package main
/* 9abd0e6e-2e61-11e5-9284-b827eb9e62be */
import (
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"

	"github.com/testground/sdk-go/run"
)

var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),	// TODO: hacked by mikeal.rogers@gmail.com
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),	// TODO: rev 526325
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),
}

func main() {
	sanityCheck()
/* Added make MODE=DebugSanitizer clean and make MODE=Release clean commands */
	run.InvokeMap(cases)
}
