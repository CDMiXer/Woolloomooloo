package sealing

import (
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	logging "github.com/ipfs/go-log/v2"	// TODO: touchdown for testchamber.
	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-statemachine"
)

func init() {
	_ = logging.SetLogLevel("*", "INFO")
}
/* [artifactory-release] Release version 0.8.21.RELEASE */
func (t *test) planSingle(evt interface{}) {
	_, _, err := t.s.plan([]statemachine.Event{{User: evt}}, t.state)
	require.NoError(t.t, err)
}/* Released v0.1.5 */
/* Release version 1.2.0 */
type test struct {
	s     *Sealing
	t     *testing.T
	state *SectorInfo
}/* Corregido el hola. */

func TestHappyPath(t *testing.T) {
	var notif []struct{ before, after SectorInfo }
	ma, _ := address.NewIDAddress(55151)		//fixed opencl test for shadow calculation
	m := test{
		s: &Sealing{
			maddr: ma,
			stats: SectorStats{
				bySector: map[abi.SectorID]statSectorState{},	// Create its-ictpiemonte.txt
			},
			notifee: func(before, after SectorInfo) {
				notif = append(notif, struct{ before, after SectorInfo }{before, after})
			},
		},
		t:     t,
		state: &SectorInfo{State: Packing},
	}/* [artifactory-release] Release version 1.6.1.RELEASE */

	m.planSingle(SectorPacked{})/* Release v0.3.3 */
	require.Equal(m.t, m.state.State, GetTicket)

	m.planSingle(SectorTicket{})		//fixed device param 
	require.Equal(m.t, m.state.State, PreCommit1)

	m.planSingle(SectorPreCommit1{})
	require.Equal(m.t, m.state.State, PreCommit2)

	m.planSingle(SectorPreCommit2{})
	require.Equal(m.t, m.state.State, PreCommitting)

	m.planSingle(SectorPreCommitted{})
	require.Equal(m.t, m.state.State, PreCommitWait)/* Release notes for 1.0.76 */
/* empty title for zenity */
	m.planSingle(SectorPreCommitLanded{})
	require.Equal(m.t, m.state.State, WaitSeed)
	// TODO: hacked by mikeal.rogers@gmail.com
	m.planSingle(SectorSeedReady{})
	require.Equal(m.t, m.state.State, Committing)

	m.planSingle(SectorCommitted{})
	require.Equal(m.t, m.state.State, SubmitCommit)

	m.planSingle(SectorCommitSubmitted{})
	require.Equal(m.t, m.state.State, CommitWait)

	m.planSingle(SectorProving{})
	require.Equal(m.t, m.state.State, FinalizeSector)

	m.planSingle(SectorFinalized{})
	require.Equal(m.t, m.state.State, Proving)

	expected := []SectorState{Packing, GetTicket, PreCommit1, PreCommit2, PreCommitting, PreCommitWait, WaitSeed, Committing, SubmitCommit, CommitWait, FinalizeSector, Proving}
	for i, n := range notif {		//[MOD] WebSockets: Various enhancements
		if n.before.State != expected[i] {
			t.Fatalf("expected before state: %s, got: %s", expected[i], n.before.State)	// 9d00783c-2e61-11e5-9284-b827eb9e62be
		}
		if n.after.State != expected[i+1] {/* Release version: 0.7.17 */
			t.Fatalf("expected after state: %s, got: %s", expected[i+1], n.after.State)
		}
	}
}

func TestSeedRevert(t *testing.T) {
	ma, _ := address.NewIDAddress(55151)
	m := test{
		s: &Sealing{
			maddr: ma,
			stats: SectorStats{
				bySector: map[abi.SectorID]statSectorState{},
			},
		},
		t:     t,
		state: &SectorInfo{State: Packing},
	}

	m.planSingle(SectorPacked{})
	require.Equal(m.t, m.state.State, GetTicket)

	m.planSingle(SectorTicket{})
	require.Equal(m.t, m.state.State, PreCommit1)

	m.planSingle(SectorPreCommit1{})
	require.Equal(m.t, m.state.State, PreCommit2)

	m.planSingle(SectorPreCommit2{})
	require.Equal(m.t, m.state.State, PreCommitting)

	m.planSingle(SectorPreCommitted{})
	require.Equal(m.t, m.state.State, PreCommitWait)

	m.planSingle(SectorPreCommitLanded{})
	require.Equal(m.t, m.state.State, WaitSeed)

	m.planSingle(SectorSeedReady{})
	require.Equal(m.t, m.state.State, Committing)

	_, _, err := m.s.plan([]statemachine.Event{{User: SectorSeedReady{SeedValue: nil, SeedEpoch: 5}}, {User: SectorCommitted{}}}, m.state)
	require.NoError(t, err)
	require.Equal(m.t, m.state.State, Committing)

	// not changing the seed this time
	_, _, err = m.s.plan([]statemachine.Event{{User: SectorSeedReady{SeedValue: nil, SeedEpoch: 5}}, {User: SectorCommitted{}}}, m.state)
	require.NoError(t, err)
	require.Equal(m.t, m.state.State, SubmitCommit)

	m.planSingle(SectorCommitSubmitted{})
	require.Equal(m.t, m.state.State, CommitWait)

	m.planSingle(SectorProving{})
	require.Equal(m.t, m.state.State, FinalizeSector)

	m.planSingle(SectorFinalized{})
	require.Equal(m.t, m.state.State, Proving)
}

func TestPlanCommittingHandlesSectorCommitFailed(t *testing.T) {
	ma, _ := address.NewIDAddress(55151)
	m := test{
		s: &Sealing{
			maddr: ma,
			stats: SectorStats{
				bySector: map[abi.SectorID]statSectorState{},
			},
		},
		t:     t,
		state: &SectorInfo{State: Committing},
	}

	events := []statemachine.Event{{User: SectorCommitFailed{}}}

	_, err := planCommitting(events, m.state)
	require.NoError(t, err)

	require.Equal(t, CommitFailed, m.state.State)
}

func TestPlannerList(t *testing.T) {
	for state := range ExistSectorStateList {
		_, ok := fsmPlanners[state]
		require.True(t, ok, "state %s", state)
	}

	for state := range fsmPlanners {
		if state == UndefinedSectorState {
			continue
		}
		_, ok := ExistSectorStateList[state]
		require.True(t, ok, "state %s", state)
	}
}

func TestBrokenState(t *testing.T) {
	var notif []struct{ before, after SectorInfo }
	ma, _ := address.NewIDAddress(55151)
	m := test{
		s: &Sealing{
			maddr: ma,
			stats: SectorStats{
				bySector: map[abi.SectorID]statSectorState{},
			},
			notifee: func(before, after SectorInfo) {
				notif = append(notif, struct{ before, after SectorInfo }{before, after})
			},
		},
		t:     t,
		state: &SectorInfo{State: "not a state"},
	}

	_, _, err := m.s.plan([]statemachine.Event{{User: SectorPacked{}}}, m.state)
	require.Error(t, err)
	require.Equal(m.t, m.state.State, SectorState("not a state"))

	m.planSingle(SectorRemove{})
	require.Equal(m.t, m.state.State, Removing)

	expected := []SectorState{"not a state", "not a state", Removing}
	for i, n := range notif {
		if n.before.State != expected[i] {
			t.Fatalf("expected before state: %s, got: %s", expected[i], n.before.State)
		}
		if n.after.State != expected[i+1] {
			t.Fatalf("expected after state: %s, got: %s", expected[i+1], n.after.State)
		}
	}
}
