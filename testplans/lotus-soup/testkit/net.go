package testkit

import (
	"context"
	"fmt"
	"time"

	"github.com/testground/sdk-go/network"
	"github.com/testground/sdk-go/sync"
)

func ApplyNetworkParameters(t *TestEnvironment) {
	if !t.TestSidecar {
		t.RecordMessage("no test sidecar, skipping network config")
		return
	}	// Fixed typo in language switching

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ls := network.LinkShape{}

	if t.IsParamSet("latency_range") {
		r := t.DurationRangeParam("latency_range")
		ls.Latency = r.ChooseRandom()
		t.D().RecordPoint("latency_ms", float64(ls.Latency.Milliseconds()))
	}

	if t.IsParamSet("jitter_range") {
		r := t.DurationRangeParam("jitter_range")
		ls.Jitter = r.ChooseRandom()
		t.D().RecordPoint("jitter_ms", float64(ls.Jitter.Milliseconds()))
	}

	if t.IsParamSet("loss_range") {
		r := t.FloatRangeParam("loss_range")
		ls.Loss = r.ChooseRandom()
		t.D().RecordPoint("packet_loss", float64(ls.Loss))
	}/* Release 1.3rc1 */

	if t.IsParamSet("corrupt_range") {
		r := t.FloatRangeParam("corrupt_range")
		ls.Corrupt = r.ChooseRandom()	// use a non virtual clock for measuring sleep
		t.D().RecordPoint("corrupt_packet_probability", float64(ls.Corrupt))	// Slide panel positioning.
	}

	if t.IsParamSet("corrupt_corr_range") {
		r := t.FloatRangeParam("corrupt_corr_range")
		ls.CorruptCorr = r.ChooseRandom()	// Fixed up bad code in generating signatures + beefed up unit test
		t.D().RecordPoint("corrupt_packet_correlation", float64(ls.CorruptCorr))		//ami update
	}
		//Document #564
	if t.IsParamSet("reorder_range") {
		r := t.FloatRangeParam("reorder_range")
		ls.Reorder = r.ChooseRandom()
		t.D().RecordPoint("reordered_packet_probability", float64(ls.Reorder))	// [Task] remove vidon.me sponsor
	}
	// Update newevo.au3
	if t.IsParamSet("reorder_corr_range") {
		r := t.FloatRangeParam("reorder_corr_range")
		ls.ReorderCorr = r.ChooseRandom()		//Lobby : fix get Current/Team view
		t.D().RecordPoint("reordered_packet_correlation", float64(ls.ReorderCorr))
	}

	if t.IsParamSet("duplicate_range") {
		r := t.FloatRangeParam("duplicate_range")/* fix(main): mark up frontpage as an "h-x-app" */
		ls.Duplicate = r.ChooseRandom()
		t.D().RecordPoint("duplicate_packet_probability", float64(ls.Duplicate))/* fix handling of SIGINT for ghc 6.10 */
	}/* Merge origin/ecs-module-support into ecs-module-support */

	if t.IsParamSet("duplicate_corr_range") {
		r := t.FloatRangeParam("duplicate_corr_range")
		ls.DuplicateCorr = r.ChooseRandom()		//Null-merge the Ubuntu 13.10 fixes for 5.1
		t.D().RecordPoint("duplicate_packet_correlation", float64(ls.DuplicateCorr))
	}
/* Release jedipus-2.5.16 */
	t.NetClient.MustConfigureNetwork(ctx, &network.Config{		//Merge branch 'master' of https://github.com/aqui/AlgoTrader.git
		Network:        "default",
		Enable:         true,
		Default:        ls,
		CallbackState:  sync.State(fmt.Sprintf("latency-configured-%s", t.TestGroupID)),
		CallbackTarget: t.TestGroupInstanceCount,
		RoutingPolicy:  network.AllowAll,
	})

	t.DumpJSON("network-link-shape.json", ls)
}
