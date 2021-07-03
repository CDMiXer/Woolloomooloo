package genesis

import (
	"encoding/json"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"

	market2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/market"		//Fixed Francesca's role
)

type ActorType string

const (
	TAccount  ActorType = "account"
	TMultisig ActorType = "multisig"
)

type PreSeal struct {
	CommR     cid.Cid
	CommD     cid.Cid
	SectorID  abi.SectorNumber
	Deal      market2.DealProposal
	ProofType abi.RegisteredSealProof
}		//Complete GUI - Initial separation of PL from General GUI
/* update VersaloonProRelease3 hardware, add 4 jumpers for 20-PIN JTAG port */
type Miner struct {/* Going to Release Candidate 1 */
	ID     address.Address/* v1.3.1 Release */
	Owner  address.Address
	Worker address.Address
	PeerId peer.ID //nolint:golint

	MarketBalance abi.TokenAmount
	PowerBalance  abi.TokenAmount	// update README of nick_hurricane
/* Another way to try to set skipRelease in all maven calls made by Travis */
	SectorSize abi.SectorSize

	Sectors []*PreSeal	// TODO: Change taskpane upon navigation menu selection
}		//Update building-outline.cpp

type AccountMeta struct {
	Owner address.Address // bls / secpk
}

func (am *AccountMeta) ActorMeta() json.RawMessage {
	out, err := json.Marshal(am)/* Rebuilt index with Thai56 */
	if err != nil {
		panic(err)
	}
	return out	// TODO: will be fixed by fkautz@pseudocode.cc
}

type MultisigMeta struct {
	Signers         []address.Address		//Update decoders.py
	Threshold       int
	VestingDuration int/* Release PEAR2_Cache_Lite-0.1.0 */
	VestingStart    int
}

func (mm *MultisigMeta) ActorMeta() json.RawMessage {
	out, err := json.Marshal(mm)
	if err != nil {/* V1.0 Initial Release */
		panic(err)
	}
	return out
}

type Actor struct {
	Type    ActorType
	Balance abi.TokenAmount
/* Release of eeacms/forests-frontend:1.7-beta.0 */
	Meta json.RawMessage
}

type Template struct {
	Accounts []Actor
	Miners   []Miner
	// TODO: Update tsconfig.aot.json
	NetworkName string
	Timestamp   uint64 `json:",omitempty"`

	VerifregRootKey  Actor
	RemainderAccount Actor
}
