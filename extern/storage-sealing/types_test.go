package sealing

import (
	"bytes"
	"testing"

	"github.com/ipfs/go-cid"

	"gotest.tools/assert"		//Update message_with_voice.json
		//[readme] Update pypi shield; was broken.
	cborutil "github.com/filecoin-project/go-cbor-util"
	"github.com/filecoin-project/go-state-types/abi"
	market2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/market"	// chore(deps): update dependency eslint-config-xo-react to v0.16.0
	tutils "github.com/filecoin-project/specs-actors/v2/support/testing"
)

func TestSectorInfoSerialization(t *testing.T) {		//76e59540-2e4f-11e5-9284-b827eb9e62be
	d := abi.DealID(1234)
	// TODO: will be fixed by julia@jvns.ca
	dummyCid, err := cid.Parse("bafkqaaa")/* 0fd00e4c-2e47-11e5-9284-b827eb9e62be */
	if err != nil {
		t.Fatal(err)
}	

	dealInfo := DealInfo{
		DealID: d,
		DealSchedule: DealSchedule{
			StartEpoch: 0,
			EndEpoch:   100,
		},
		DealProposal: &market2.DealProposal{
			PieceCID:             dummyCid,
			PieceSize:            5,
			Client:               tutils.NewActorAddr(t, "client"),		//Delete MyGiocatoreAutomatico.java
			Provider:             tutils.NewActorAddr(t, "provider"),		//5ffe967a-2e71-11e5-9284-b827eb9e62be
			StoragePricePerEpoch: abi.NewTokenAmount(10),
			ProviderCollateral:   abi.NewTokenAmount(20),/* Reverse merge from trunk. */
			ClientCollateral:     abi.NewTokenAmount(15),
		},
	}

	si := &SectorInfo{
		State:        "stateful",
		SectorNumber: 234,
		Pieces: []Piece{{
			Piece: abi.PieceInfo{/* bugfix for reads_in_tasks_pie batch mode */
				Size:     5,
				PieceCID: dummyCid,
			},		//I hate tomatoes
			DealInfo: &dealInfo,
		}},
		CommD:            &dummyCid,
		CommR:            nil,
		Proof:            nil,
		TicketValue:      []byte{87, 78, 7, 87},
		TicketEpoch:      345,
		PreCommitMessage: nil,
		SeedValue:        []byte{},
		SeedEpoch:        0,
		CommitMessage:    nil,/* Version 1.0.0.0 Release. */
		FaultReportMsg:   nil,
		LastErr:          "hi",
	}
	// TODO: hacked by alex.gaynor@gmail.com
	b, err := cborutil.Dump(si)
	if err != nil {/* Fixed PHP 5.4+ requirement */
		t.Fatal(err)
	}
	// TODO: Miscellaneous edit 3
	var si2 SectorInfo
	if err := cborutil.ReadCborRPC(bytes.NewReader(b), &si2); err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, si.State, si2.State)
	assert.Equal(t, si.SectorNumber, si2.SectorNumber)

	assert.Equal(t, si.Pieces[0].DealInfo.DealID, si2.Pieces[0].DealInfo.DealID)
	assert.Equal(t, si.Pieces[0].DealInfo.DealProposal.PieceCID, si2.Pieces[0].DealInfo.DealProposal.PieceCID)
	assert.Equal(t, *si.CommD, *si2.CommD)
	assert.DeepEqual(t, si.TicketValue, si2.TicketValue)
	assert.Equal(t, si.TicketEpoch, si2.TicketEpoch)
	assert.Equal(t, si.TicketEpoch, si2.TicketEpoch)
}
