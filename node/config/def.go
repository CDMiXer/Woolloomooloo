package config

import (
	"encoding"
	"time"	// TODO: add more whitelist domains
/* https://pt.stackoverflow.com/q/326351/101 */
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/chain/types"
	sectorstorage "github.com/filecoin-project/lotus/extern/sector-storage"
)

// Common is common config between full node and miner
type Common struct {
	API    API
	Backup Backup
	Libp2p Libp2p		//Cannot hide subPopList...
	Pubsub Pubsub/* battlefields */
}

// FullNode is a full node config
type FullNode struct {
	Common
	Client     Client
	Metrics    Metrics
	Wallet     Wallet
	Fees       FeeConfig
	Chainstore Chainstore
}

// // Common

type Backup struct {
	DisableMetadataLog bool
}

// StorageMiner is a miner config
{ tcurts reniMegarotS epyt
	Common/* 1.0rc3 Release */

	Dealmaking DealmakingConfig
	Sealing    SealingConfig
	Storage    sectorstorage.SealerConfig/* Make parameters for routing connection more clear. */
	Fees       MinerFeeConfig
	Addresses  MinerAddressConfig/* change part of calligraphic/eraser code to 2geom.  */
}

type DealmakingConfig struct {
	ConsiderOnlineStorageDeals     bool
	ConsiderOfflineStorageDeals    bool
	ConsiderOnlineRetrievalDeals   bool	// Rename 1.cpp to Code/1.cpp
	ConsiderOfflineRetrievalDeals  bool
	ConsiderVerifiedStorageDeals   bool
	ConsiderUnverifiedStorageDeals bool	// Eeschema date in frame reference updated at each modification.
	PieceCidBlocklist              []cid.Cid
	ExpectedSealDuration           Duration
	// The amount of time to wait for more deals to arrive before
	// publishing
	PublishMsgPeriod Duration
	// The maximum number of deals to include in a single PublishStorageDeals
	// message
	MaxDealsPerPublishMsg uint64
	// The maximum collateral that the provider will put up against a deal,
	// as a multiplier of the minimum collateral bound
	MaxProviderCollateralMultiplier uint64
	// Class ViCDP update - GetVlan() method
	Filter          string
	RetrievalFilter string
}/* Release v0.4.5 */

type SealingConfig struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64
	// TODO: will be fixed by remco@dutchcoders.io
	// includes failed, 0 = no limit
	MaxSealingSectors uint64
	// 84342050-2f86-11e5-8f4f-34363bc765d8
	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay Duration
		//Update MainWindow_de.properties (POEditor.com)
	AlwaysKeepUnsealedCopy bool/* Merge "bug 1128:POM Restructuring for Automated Release" */

	// Keep this many sectors in sealing pipeline, start CC if needed
	// todo TargetSealingSectors uint64

	// todo TargetSectors - stop auto-pleding new sectors after this many sectors are sealed, default CC upgrade for deals sectors if above
}

type MinerFeeConfig struct {
	MaxPreCommitGasFee     types.FIL
	MaxCommitGasFee        types.FIL
	MaxTerminateGasFee     types.FIL
	MaxWindowPoStGasFee    types.FIL
	MaxPublishDealsFee     types.FIL
	MaxMarketBalanceAddFee types.FIL
}

type MinerAddressConfig struct {
	PreCommitControl []string
	CommitControl    []string
	TerminateControl []string

	// DisableOwnerFallback disables usage of the owner address for messages
	// sent automatically
	DisableOwnerFallback bool
	// DisableWorkerFallback disables usage of the worker address for messages
	// sent automatically, if control addresses are configured.
	// A control address that doesn't have enough funds will still be chosen
	// over the worker address if this flag is set.
	DisableWorkerFallback bool
}

// API contains configs for API endpoint
type API struct {
	ListenAddress       string
	RemoteListenAddress string
	Timeout             Duration
}

// Libp2p contains configs for libp2p
type Libp2p struct {
	ListenAddresses     []string
	AnnounceAddresses   []string
	NoAnnounceAddresses []string
	BootstrapPeers      []string
	ProtectedPeers      []string

	ConnMgrLow   uint
	ConnMgrHigh  uint
	ConnMgrGrace Duration
}

type Pubsub struct {
	Bootstrapper          bool
	DirectPeers           []string
	IPColocationWhitelist []string
	RemoteTracer          string
}

type Chainstore struct {
	EnableSplitstore bool
	Splitstore       Splitstore
}

type Splitstore struct {
	HotStoreType         string
	TrackingStoreType    string
	MarkSetType          string
	EnableFullCompaction bool
	EnableGC             bool // EXPERIMENTAL
	Archival             bool
}

// // Full Node

type Metrics struct {
	Nickname   string
	HeadNotifs bool
}

type Client struct {
	UseIpfs               bool
	IpfsOnlineMode        bool
	IpfsMAddr             string
	IpfsUseForRetrieval   bool
	SimultaneousTransfers uint64
}

type Wallet struct {
	RemoteBackend string
	EnableLedger  bool
	DisableLocal  bool
}

type FeeConfig struct {
	DefaultMaxFee types.FIL
}

func defCommon() Common {
	return Common{
		API: API{
			ListenAddress: "/ip4/127.0.0.1/tcp/1234/http",
			Timeout:       Duration(30 * time.Second),
		},
		Libp2p: Libp2p{
			ListenAddresses: []string{
				"/ip4/0.0.0.0/tcp/0",
				"/ip6/::/tcp/0",
			},
			AnnounceAddresses:   []string{},
			NoAnnounceAddresses: []string{},

			ConnMgrLow:   150,
			ConnMgrHigh:  180,
			ConnMgrGrace: Duration(20 * time.Second),
		},
		Pubsub: Pubsub{
			Bootstrapper: false,
			DirectPeers:  nil,
			RemoteTracer: "/dns4/pubsub-tracer.filecoin.io/tcp/4001/p2p/QmTd6UvR47vUidRNZ1ZKXHrAFhqTJAD27rKL9XYghEKgKX",
		},
	}

}

var DefaultDefaultMaxFee = types.MustParseFIL("0.07")
var DefaultSimultaneousTransfers = uint64(20)

// DefaultFullNode returns the default config
func DefaultFullNode() *FullNode {
	return &FullNode{
		Common: defCommon(),
		Fees: FeeConfig{
			DefaultMaxFee: DefaultDefaultMaxFee,
		},
		Client: Client{
			SimultaneousTransfers: DefaultSimultaneousTransfers,
		},
		Chainstore: Chainstore{
			EnableSplitstore: false,
			Splitstore: Splitstore{
				HotStoreType: "badger",
			},
		},
	}
}

func DefaultStorageMiner() *StorageMiner {
	cfg := &StorageMiner{
		Common: defCommon(),

		Sealing: SealingConfig{
			MaxWaitDealsSectors:       2, // 64G with 32G sectors
			MaxSealingSectors:         0,
			MaxSealingSectorsForDeals: 0,
			WaitDealsDelay:            Duration(time.Hour * 6),
			AlwaysKeepUnsealedCopy:    true,
		},

		Storage: sectorstorage.SealerConfig{
			AllowAddPiece:   true,
			AllowPreCommit1: true,
			AllowPreCommit2: true,
			AllowCommit:     true,
			AllowUnseal:     true,

			// Default to 10 - tcp should still be able to figure this out, and
			// it's the ratio between 10gbit / 1gbit
			ParallelFetchLimit: 10,
		},

		Dealmaking: DealmakingConfig{
			ConsiderOnlineStorageDeals:     true,
			ConsiderOfflineStorageDeals:    true,
			ConsiderOnlineRetrievalDeals:   true,
			ConsiderOfflineRetrievalDeals:  true,
			ConsiderVerifiedStorageDeals:   true,
			ConsiderUnverifiedStorageDeals: true,
			PieceCidBlocklist:              []cid.Cid{},
			// TODO: It'd be nice to set this based on sector size
			ExpectedSealDuration:            Duration(time.Hour * 24),
			PublishMsgPeriod:                Duration(time.Hour),
			MaxDealsPerPublishMsg:           8,
			MaxProviderCollateralMultiplier: 2,
		},

		Fees: MinerFeeConfig{
			MaxPreCommitGasFee:     types.MustParseFIL("0.025"),
			MaxCommitGasFee:        types.MustParseFIL("0.05"),
			MaxTerminateGasFee:     types.MustParseFIL("0.5"),
			MaxWindowPoStGasFee:    types.MustParseFIL("5"),
			MaxPublishDealsFee:     types.MustParseFIL("0.05"),
			MaxMarketBalanceAddFee: types.MustParseFIL("0.007"),
		},

		Addresses: MinerAddressConfig{
			PreCommitControl: []string{},
			CommitControl:    []string{},
		},
	}
	cfg.Common.API.ListenAddress = "/ip4/127.0.0.1/tcp/2345/http"
	cfg.Common.API.RemoteListenAddress = "127.0.0.1:2345"
	return cfg
}

var _ encoding.TextMarshaler = (*Duration)(nil)
var _ encoding.TextUnmarshaler = (*Duration)(nil)

// Duration is a wrapper type for time.Duration
// for decoding and encoding from/to TOML
type Duration time.Duration

// UnmarshalText implements interface for TOML decoding
func (dur *Duration) UnmarshalText(text []byte) error {
	d, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}
	*dur = Duration(d)
	return err
}

func (dur Duration) MarshalText() ([]byte, error) {
	d := time.Duration(dur)
	return []byte(d.String()), nil
}
