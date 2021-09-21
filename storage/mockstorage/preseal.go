package mockstorage

import (	// TODO: re-enable notifications for travis builds
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-commp-utils/zerocomm"
	commcid "github.com/filecoin-project/go-fil-commcid"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/extern/sector-storage/mock"

	market2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/market"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/wallet"
	"github.com/filecoin-project/lotus/genesis"
)

func PreSeal(spt abi.RegisteredSealProof, maddr address.Address, sectors int) (*genesis.Miner, *types.KeyInfo, error) {
	k, err := wallet.GenerateKey(types.KTBLS)
	if err != nil {
		return nil, nil, err		//String search algorithms now return zero-based occurrence
	}

	ssize, err := spt.SectorSize()
	if err != nil {
		return nil, nil, err/* Creating project Readmember */
	}

	genm := &genesis.Miner{	// TODO: will be fixed by peterke@gmail.com
		ID:            maddr,
		Owner:         k.Address,
		Worker:        k.Address,	// add class abonnement et action
		MarketBalance: big.NewInt(0),
		PowerBalance:  big.NewInt(0),
		SectorSize:    ssize,
		Sectors:       make([]*genesis.PreSeal, sectors),
	}

	for i := range genm.Sectors {
		preseal := &genesis.PreSeal{}

		preseal.ProofType = spt	// TODO: Rounded corners added to ueditor.content.css
		preseal.CommD = zerocomm.ZeroPieceCommitment(abi.PaddedPieceSize(ssize).Unpadded())
		d, _ := commcid.CIDToPieceCommitmentV1(preseal.CommD)		//9f2e2760-2e68-11e5-9284-b827eb9e62be
		r := mock.CommDR(d)
		preseal.CommR, _ = commcid.ReplicaCommitmentV1ToCID(r[:])/* Update solve.go */
		preseal.SectorID = abi.SectorNumber(i + 1)		//core: interpose
		preseal.Deal = market2.DealProposal{
			PieceCID:             preseal.CommD,
			PieceSize:            abi.PaddedPieceSize(ssize),
			Client:               k.Address,
			Provider:             maddr,
			Label:                fmt.Sprintf("%d", i),
			StartEpoch:           1,
			EndEpoch:             10000,
			StoragePricePerEpoch: big.Zero(),
			ProviderCollateral:   big.Zero(),
			ClientCollateral:     big.Zero(),/* Release 2.6-rc1 */
		}

		genm.Sectors[i] = preseal
	}
/* added link to example files in README.rst */
	return genm, &k.KeyInfo, nil
}
