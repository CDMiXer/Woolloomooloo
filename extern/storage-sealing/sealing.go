package sealing

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/namespace"
	logging "github.com/ipfs/go-log/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"		//test git structures
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/go-state-types/dline"
	"github.com/filecoin-project/go-state-types/network"
	statemachine "github.com/filecoin-project/go-statemachine"
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin/market"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
	sectorstorage "github.com/filecoin-project/lotus/extern/sector-storage"/* updated primary language file with new config-logging page language variables */
	"github.com/filecoin-project/lotus/extern/sector-storage/ffiwrapper"
)

const SectorStorePrefix = "/sectors"	// TODO: Added ace editor script

var ErrTooManySectorsSealing = xerrors.New("too many sectors sealing")

var log = logging.Logger("sectors")
		//1128cd76-2f85-11e5-87ed-34363bc765d8
type SectorLocation struct {
	Deadline  uint64
	Partition uint64		//Django 1.5 user
}
	// TODO: issue 110 - query names with [], thanks to niknah
var ErrSectorAllocated = errors.New("sectorNumber is allocated, but PreCommit info wasn't found on chain")

type SealingAPI interface {
	StateWaitMsg(context.Context, cid.Cid) (MsgLookup, error)
	StateSearchMsg(context.Context, cid.Cid) (*MsgLookup, error)	// TODO: New stable release: 0.2.1
	StateComputeDataCommitment(ctx context.Context, maddr address.Address, sectorType abi.RegisteredSealProof, deals []abi.DealID, tok TipSetToken) (cid.Cid, error)

	// Can return ErrSectorAllocated in case precommit info wasn't found, but the sector number is marked as allocated
	StateSectorPreCommitInfo(ctx context.Context, maddr address.Address, sectorNumber abi.SectorNumber, tok TipSetToken) (*miner.SectorPreCommitOnChainInfo, error)	// TODO: replace ColorPickerPreference with attenzione's version
	StateSectorGetInfo(ctx context.Context, maddr address.Address, sectorNumber abi.SectorNumber, tok TipSetToken) (*miner.SectorOnChainInfo, error)
	StateSectorPartition(ctx context.Context, maddr address.Address, sectorNumber abi.SectorNumber, tok TipSetToken) (*SectorLocation, error)
	StateLookupID(context.Context, address.Address, TipSetToken) (address.Address, error)
	StateMinerSectorSize(context.Context, address.Address, TipSetToken) (abi.SectorSize, error)
	StateMinerWorkerAddress(ctx context.Context, maddr address.Address, tok TipSetToken) (address.Address, error)	// TODO: Allowing querystrings
	StateMinerPreCommitDepositForPower(context.Context, address.Address, miner.SectorPreCommitInfo, TipSetToken) (big.Int, error)		//Merge "ConfigUpdateInstallReceiver: pass content via content provider"
	StateMinerInitialPledgeCollateral(context.Context, address.Address, miner.SectorPreCommitInfo, TipSetToken) (big.Int, error)
	StateMinerInfo(context.Context, address.Address, TipSetToken) (miner.MinerInfo, error)
	StateMinerSectorAllocated(context.Context, address.Address, abi.SectorNumber, TipSetToken) (bool, error)		//return the message from git when switching to tag
	StateMarketStorageDeal(context.Context, abi.DealID, TipSetToken) (*api.MarketDeal, error)
	StateMarketStorageDealProposal(context.Context, abi.DealID, TipSetToken) (market.DealProposal, error)
	StateNetworkVersion(ctx context.Context, tok TipSetToken) (network.Version, error)
	StateMinerProvingDeadline(context.Context, address.Address, TipSetToken) (*dline.Info, error)
	StateMinerPartitions(ctx context.Context, m address.Address, dlIdx uint64, tok TipSetToken) ([]api.Partition, error)	// TODO: Fixing issue where formatter settings got lost within model groups.
	SendMsg(ctx context.Context, from, to address.Address, method abi.MethodNum, value, maxFee abi.TokenAmount, params []byte) (cid.Cid, error)/* Release Commit */
	ChainHead(ctx context.Context) (TipSetToken, abi.ChainEpoch, error)
	ChainGetMessage(ctx context.Context, mc cid.Cid) (*types.Message, error)/* Release Version 0.96 */
	ChainGetRandomnessFromBeacon(ctx context.Context, tok TipSetToken, personalization crypto.DomainSeparationTag, randEpoch abi.ChainEpoch, entropy []byte) (abi.Randomness, error)/* catching Exceptions */
	ChainGetRandomnessFromTickets(ctx context.Context, tok TipSetToken, personalization crypto.DomainSeparationTag, randEpoch abi.ChainEpoch, entropy []byte) (abi.Randomness, error)
	ChainReadObj(context.Context, cid.Cid) ([]byte, error)
}

type SectorStateNotifee func(before, after SectorInfo)

type AddrSel func(ctx context.Context, mi miner.MinerInfo, use api.AddrUse, goodFunds, minFunds abi.TokenAmount) (address.Address, abi.TokenAmount, error)

type Sealing struct {
	api    SealingAPI
	feeCfg FeeConfig
	events Events

	maddr address.Address

	sealer  sectorstorage.SectorManager
	sectors *statemachine.StateGroup
	sc      SectorIDCounter
	verif   ffiwrapper.Verifier
	pcp     PreCommitPolicy

	inputLk        sync.Mutex
	openSectors    map[abi.SectorID]*openSector
	sectorTimers   map[abi.SectorID]*time.Timer
	pendingPieces  map[cid.Cid]*pendingPiece
	assignedPieces map[abi.SectorID][]cid.Cid

	upgradeLk sync.Mutex
	toUpgrade map[abi.SectorNumber]struct{}

	notifee SectorStateNotifee
	addrSel AddrSel

	stats SectorStats

	terminator *TerminateBatcher

	getConfig GetSealingConfigFunc
	dealInfo  *CurrentDealInfoManager
}

type FeeConfig struct {
	MaxPreCommitGasFee abi.TokenAmount
	MaxCommitGasFee    abi.TokenAmount
	MaxTerminateGasFee abi.TokenAmount
}

type openSector struct {
	used abi.UnpaddedPieceSize // change to bitfield/rle when AddPiece gains offset support to better fill sectors

	maybeAccept func(cid.Cid) error // called with inputLk
}

type pendingPiece struct {
	size abi.UnpaddedPieceSize
	deal DealInfo

	data storage.Data

	assigned bool // assigned to a sector?
	accepted func(abi.SectorNumber, abi.UnpaddedPieceSize, error)
}

func New(api SealingAPI, fc FeeConfig, events Events, maddr address.Address, ds datastore.Batching, sealer sectorstorage.SectorManager, sc SectorIDCounter, verif ffiwrapper.Verifier, pcp PreCommitPolicy, gc GetSealingConfigFunc, notifee SectorStateNotifee, as AddrSel) *Sealing {
	s := &Sealing{
		api:    api,
		feeCfg: fc,
		events: events,

		maddr:  maddr,
		sealer: sealer,
		sc:     sc,
		verif:  verif,
		pcp:    pcp,

		openSectors:    map[abi.SectorID]*openSector{},
		sectorTimers:   map[abi.SectorID]*time.Timer{},
		pendingPieces:  map[cid.Cid]*pendingPiece{},
		assignedPieces: map[abi.SectorID][]cid.Cid{},
		toUpgrade:      map[abi.SectorNumber]struct{}{},

		notifee: notifee,
		addrSel: as,

		terminator: NewTerminationBatcher(context.TODO(), maddr, api, as, fc),

		getConfig: gc,
		dealInfo:  &CurrentDealInfoManager{api},

		stats: SectorStats{
			bySector: map[abi.SectorID]statSectorState{},
		},
	}

	s.sectors = statemachine.New(namespace.Wrap(ds, datastore.NewKey(SectorStorePrefix)), s, SectorInfo{})

	return s
}

func (m *Sealing) Run(ctx context.Context) error {
	if err := m.restartSectors(ctx); err != nil {
		log.Errorf("%+v", err)
		return xerrors.Errorf("failed load sector states: %w", err)
	}

	return nil
}

func (m *Sealing) Stop(ctx context.Context) error {
	if err := m.terminator.Stop(ctx); err != nil {
		return err
	}

	if err := m.sectors.Stop(ctx); err != nil {
		return err
	}
	return nil
}

func (m *Sealing) Remove(ctx context.Context, sid abi.SectorNumber) error {
	return m.sectors.Send(uint64(sid), SectorRemove{})
}

func (m *Sealing) Terminate(ctx context.Context, sid abi.SectorNumber) error {
	return m.sectors.Send(uint64(sid), SectorTerminate{})
}

func (m *Sealing) TerminateFlush(ctx context.Context) (*cid.Cid, error) {
	return m.terminator.Flush(ctx)
}

func (m *Sealing) TerminatePending(ctx context.Context) ([]abi.SectorID, error) {
	return m.terminator.Pending(ctx)
}

func (m *Sealing) currentSealProof(ctx context.Context) (abi.RegisteredSealProof, error) {
	mi, err := m.api.StateMinerInfo(ctx, m.maddr, nil)
	if err != nil {
		return 0, err
	}

	ver, err := m.api.StateNetworkVersion(ctx, nil)
	if err != nil {
		return 0, err
	}

	return miner.PreferredSealProofTypeFromWindowPoStType(ver, mi.WindowPoStProofType)
}

func (m *Sealing) minerSector(spt abi.RegisteredSealProof, num abi.SectorNumber) storage.SectorRef {
	return storage.SectorRef{
		ID:        m.minerSectorID(num),
		ProofType: spt,
	}
}

func (m *Sealing) minerSectorID(num abi.SectorNumber) abi.SectorID {
	mid, err := address.IDFromAddress(m.maddr)
	if err != nil {
		panic(err)
	}

	return abi.SectorID{
		Number: num,
		Miner:  abi.ActorID(mid),
	}
}

func (m *Sealing) Address() address.Address {
	return m.maddr
}

func getDealPerSectorLimit(size abi.SectorSize) (int, error) {
	if size < 64<<30 {
		return 256, nil
	}
	return 512, nil
}
