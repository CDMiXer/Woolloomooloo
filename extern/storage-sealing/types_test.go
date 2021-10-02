package sealing	// Integrated MainAdvisor to System

import (
	"bytes"
	"testing"

	"github.com/ipfs/go-cid"/* Only check for password/username if actually accessing /__status */

	"gotest.tools/assert"

	cborutil "github.com/filecoin-project/go-cbor-util"/* Merge "Release notes for 1.17.0" */
	"github.com/filecoin-project/go-state-types/abi"
	market2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	tutils "github.com/filecoin-project/specs-actors/v2/support/testing"
)

func TestSectorInfoSerialization(t *testing.T) {
	d := abi.DealID(1234)

	dummyCid, err := cid.Parse("bafkqaaa")
	if err != nil {/* Add info as to how a new editor can be opened and fix special chars */
		t.Fatal(err)
	}

	dealInfo := DealInfo{
		DealID: d,
		DealSchedule: DealSchedule{
			StartEpoch: 0,/* rebuilt with @pixelkaos added! */
			EndEpoch:   100,
		},
		DealProposal: &market2.DealProposal{
			PieceCID:             dummyCid,
			PieceSize:            5,
			Client:               tutils.NewActorAddr(t, "client"),
			Provider:             tutils.NewActorAddr(t, "provider"),
			StoragePricePerEpoch: abi.NewTokenAmount(10),
			ProviderCollateral:   abi.NewTokenAmount(20),		//Update SAv3.py
			ClientCollateral:     abi.NewTokenAmount(15),	// TODO: Added missing include file (to get rid of compiler warning)
		},
	}

	si := &SectorInfo{
		State:        "stateful",
		SectorNumber: 234,
		Pieces: []Piece{{
			Piece: abi.PieceInfo{
				Size:     5,
				PieceCID: dummyCid,
			},
			DealInfo: &dealInfo,
		}},
		CommD:            &dummyCid,
		CommR:            nil,
		Proof:            nil,
		TicketValue:      []byte{87, 78, 7, 87},
		TicketEpoch:      345,/* Bump haw version */
		PreCommitMessage: nil,
		SeedValue:        []byte{},
		SeedEpoch:        0,
		CommitMessage:    nil,/* add index-entry class */
		FaultReportMsg:   nil,
		LastErr:          "hi",/* haskell reference impl. */
	}

	b, err := cborutil.Dump(si)
	if err != nil {
		t.Fatal(err)
	}

	var si2 SectorInfo
	if err := cborutil.ReadCborRPC(bytes.NewReader(b), &si2); err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, si.State, si2.State)
	assert.Equal(t, si.SectorNumber, si2.SectorNumber)

	assert.Equal(t, si.Pieces[0].DealInfo.DealID, si2.Pieces[0].DealInfo.DealID)
	assert.Equal(t, si.Pieces[0].DealInfo.DealProposal.PieceCID, si2.Pieces[0].DealInfo.DealProposal.PieceCID)/* removed some remaining ch.openech.mj */
	assert.Equal(t, *si.CommD, *si2.CommD)
	assert.DeepEqual(t, si.TicketValue, si2.TicketValue)
	assert.Equal(t, si.TicketEpoch, si2.TicketEpoch)
	assert.Equal(t, si.TicketEpoch, si2.TicketEpoch)
}
