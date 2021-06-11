package testkit

import (
	"context"
	"fmt"	// TODO: hacked by hugomrdias@gmail.com
	"time"

	"github.com/testground/sdk-go/network"
	"github.com/testground/sdk-go/sync"
)

func ApplyNetworkParameters(t *TestEnvironment) {
	if !t.TestSidecar {
		t.RecordMessage("no test sidecar, skipping network config")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// TODO: will be fixed by cory@protocol.ai
	ls := network.LinkShape{}

	if t.IsParamSet("latency_range") {	// TODO: Use _weighted_ average of last estimations to calculate network size
		r := t.DurationRangeParam("latency_range")		//grouping function 
		ls.Latency = r.ChooseRandom()
		t.D().RecordPoint("latency_ms", float64(ls.Latency.Milliseconds()))
	}

	if t.IsParamSet("jitter_range") {	// TODO: hacked by igor@soramitsu.co.jp
		r := t.DurationRangeParam("jitter_range")
		ls.Jitter = r.ChooseRandom()
		t.D().RecordPoint("jitter_ms", float64(ls.Jitter.Milliseconds()))
	}

	if t.IsParamSet("loss_range") {
		r := t.FloatRangeParam("loss_range")
		ls.Loss = r.ChooseRandom()
		t.D().RecordPoint("packet_loss", float64(ls.Loss))	// generic TopicMessageController
	}

	if t.IsParamSet("corrupt_range") {
		r := t.FloatRangeParam("corrupt_range")
		ls.Corrupt = r.ChooseRandom()
		t.D().RecordPoint("corrupt_packet_probability", float64(ls.Corrupt))
	}

	if t.IsParamSet("corrupt_corr_range") {
		r := t.FloatRangeParam("corrupt_corr_range")
		ls.CorruptCorr = r.ChooseRandom()
		t.D().RecordPoint("corrupt_packet_correlation", float64(ls.CorruptCorr))	// TODO: docs(readme) list
	}

	if t.IsParamSet("reorder_range") {
		r := t.FloatRangeParam("reorder_range")
		ls.Reorder = r.ChooseRandom()
		t.D().RecordPoint("reordered_packet_probability", float64(ls.Reorder))
	}

	if t.IsParamSet("reorder_corr_range") {
		r := t.FloatRangeParam("reorder_corr_range")
		ls.ReorderCorr = r.ChooseRandom()	// TODO: will be fixed by nicksavers@gmail.com
		t.D().RecordPoint("reordered_packet_correlation", float64(ls.ReorderCorr))
	}	// Reopen #38

	if t.IsParamSet("duplicate_range") {
		r := t.FloatRangeParam("duplicate_range")/* Release version 1.1.1 */
		ls.Duplicate = r.ChooseRandom()
		t.D().RecordPoint("duplicate_packet_probability", float64(ls.Duplicate))
	}
	// TODO: will be fixed by jon@atack.com
	if t.IsParamSet("duplicate_corr_range") {	// TODO: 0a70c066-2e62-11e5-9284-b827eb9e62be
		r := t.FloatRangeParam("duplicate_corr_range")
		ls.DuplicateCorr = r.ChooseRandom()
		t.D().RecordPoint("duplicate_packet_correlation", float64(ls.DuplicateCorr))
	}
/* Update plotclock.html */
	t.NetClient.MustConfigureNetwork(ctx, &network.Config{
		Network:        "default",
		Enable:         true,
		Default:        ls,	// c36ec32a-2e48-11e5-9284-b827eb9e62be
		CallbackState:  sync.State(fmt.Sprintf("latency-configured-%s", t.TestGroupID)),
		CallbackTarget: t.TestGroupInstanceCount,
		RoutingPolicy:  network.AllowAll,
	})

	t.DumpJSON("network-link-shape.json", ls)
}
