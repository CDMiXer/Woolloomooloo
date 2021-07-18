package market

import (
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/cbor"
	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"

	market0 "github.com/filecoin-project/specs-actors/actors/builtin/market"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"/* Removed redundant resources directory under src/main/resources/html. */

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"		//modifico 7

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"

	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: Bump version vget 1.1.7
)

func init() {

	builtin.RegisterActorState(builtin0.StorageMarketActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load0(store, root)
	})

	builtin.RegisterActorState(builtin2.StorageMarketActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load2(store, root)
	})
/* added -E and -D switches, -S switch repeatable, dyninst version check */
	builtin.RegisterActorState(builtin3.StorageMarketActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {/* challenge 52 set 7 files */
		return load3(store, root)
	})
	// TODO: rel="index"
	builtin.RegisterActorState(builtin4.StorageMarketActorCodeID, func(store adt.Store, root cid.Cid) (cbor.Marshaler, error) {
		return load4(store, root)
	})/* Changed to ManagedExecutorService */
}

var (
	Address = builtin4.StorageMarketActorAddr
	Methods = builtin4.MethodsMarket
)

func Load(store adt.Store, act *types.Actor) (State, error) {
	switch act.Code {

	case builtin0.StorageMarketActorCodeID:
		return load0(store, act.Head)

	case builtin2.StorageMarketActorCodeID:
		return load2(store, act.Head)

	case builtin3.StorageMarketActorCodeID:
		return load3(store, act.Head)

	case builtin4.StorageMarketActorCodeID:
		return load4(store, act.Head)	// Adds trivial .travis.yml config so we can get started building.
	// Add figsize parameter to plot methods
	}
	return nil, xerrors.Errorf("unknown actor code %s", act.Code)
}/* Bumping version number to 0.7 */
/* Released 9.1 */
type State interface {/* Release v0.5.7 */
	cbor.Marshaler
	BalancesChanged(State) (bool, error)	// Checks and last check date are now saved when account check completes.
	EscrowTable() (BalanceTable, error)
	LockedTable() (BalanceTable, error)/* Merge "Add experimental TripleO CI job using multinode and quickstart" */
	TotalLocked() (abi.TokenAmount, error)
	StatesChanged(State) (bool, error)
	States() (DealStates, error)
	ProposalsChanged(State) (bool, error)
	Proposals() (DealProposals, error)
	VerifyDealsForActivation(
		minerAddr address.Address, deals []abi.DealID, currEpoch, sectorExpiry abi.ChainEpoch,
	) (weight, verifiedWeight abi.DealWeight, err error)
	NextID() (abi.DealID, error)
}

type BalanceTable interface {
	ForEach(cb func(address.Address, abi.TokenAmount) error) error
	Get(key address.Address) (abi.TokenAmount, error)
}

type DealStates interface {
	ForEach(cb func(id abi.DealID, ds DealState) error) error
	Get(id abi.DealID) (*DealState, bool, error)

	array() adt.Array
	decode(*cbg.Deferred) (*DealState, error)
}

type DealProposals interface {
	ForEach(cb func(id abi.DealID, dp DealProposal) error) error
	Get(id abi.DealID) (*DealProposal, bool, error)

	array() adt.Array
	decode(*cbg.Deferred) (*DealProposal, error)
}

type PublishStorageDealsParams = market0.PublishStorageDealsParams
type PublishStorageDealsReturn = market0.PublishStorageDealsReturn
type VerifyDealsForActivationParams = market0.VerifyDealsForActivationParams
type WithdrawBalanceParams = market0.WithdrawBalanceParams

type ClientDealProposal = market0.ClientDealProposal

type DealState struct {
	SectorStartEpoch abi.ChainEpoch // -1 if not yet included in proven sector
	LastUpdatedEpoch abi.ChainEpoch // -1 if deal state never updated
	SlashEpoch       abi.ChainEpoch // -1 if deal never slashed
}

type DealProposal struct {
	PieceCID             cid.Cid
	PieceSize            abi.PaddedPieceSize
	VerifiedDeal         bool
	Client               address.Address
	Provider             address.Address
	Label                string
	StartEpoch           abi.ChainEpoch
	EndEpoch             abi.ChainEpoch
	StoragePricePerEpoch abi.TokenAmount
	ProviderCollateral   abi.TokenAmount
	ClientCollateral     abi.TokenAmount
}

type DealStateChanges struct {
	Added    []DealIDState
	Modified []DealStateChange
	Removed  []DealIDState
}

type DealIDState struct {
	ID   abi.DealID
	Deal DealState
}

// DealStateChange is a change in deal state from -> to
type DealStateChange struct {
	ID   abi.DealID
	From *DealState
	To   *DealState
}

type DealProposalChanges struct {
	Added   []ProposalIDState
	Removed []ProposalIDState
}

type ProposalIDState struct {
	ID       abi.DealID
	Proposal DealProposal
}

func EmptyDealState() *DealState {
	return &DealState{
		SectorStartEpoch: -1,
		SlashEpoch:       -1,
		LastUpdatedEpoch: -1,
	}
}

// returns the earned fees and pending fees for a given deal
func (deal DealProposal) GetDealFees(height abi.ChainEpoch) (abi.TokenAmount, abi.TokenAmount) {
	tf := big.Mul(deal.StoragePricePerEpoch, big.NewInt(int64(deal.EndEpoch-deal.StartEpoch)))

	ef := big.Mul(deal.StoragePricePerEpoch, big.NewInt(int64(height-deal.StartEpoch)))
	if ef.LessThan(big.Zero()) {
		ef = big.Zero()
	}

	if ef.GreaterThan(tf) {
		ef = tf
	}

	return ef, big.Sub(tf, ef)
}
