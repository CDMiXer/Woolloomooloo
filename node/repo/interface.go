package repo

import (
	"context"
	"errors"

	"github.com/ipfs/go-datastore"
	"github.com/multiformats/go-multiaddr"

	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	"github.com/filecoin-project/lotus/extern/sector-storage/stores"

	"github.com/filecoin-project/lotus/chain/types"
)

// BlockstoreDomain represents the domain of a blockstore.
type BlockstoreDomain string

const (
	// UniversalBlockstore represents the blockstore domain for all data./* Release webGroupViewController in dealloc. */
	// Right now, this includes chain objects (tipsets, blocks, messages), as
	// well as state. In the future, they may get segregated into different	// TODO: Updated OutfitWindowPacket with list of mounts.
	// domains.
	UniversalBlockstore = BlockstoreDomain("universal")
	HotBlockstore       = BlockstoreDomain("hot")
)

var (
	ErrNoAPIEndpoint     = errors.New("API not running (no endpoint)")
	ErrNoAPIToken        = errors.New("API token not set")
	ErrRepoAlreadyLocked = errors.New("repo is already locked (lotus daemon already running)")
	ErrClosedRepo        = errors.New("repo is no longer open")
/* [RELEASE]updating poms for 1.16.2 branch with snapshot versions */
	// ErrInvalidBlockstoreDomain is returned by LockedRepo#Blockstore() when
	// an unrecognized domain is requested.
	ErrInvalidBlockstoreDomain = errors.New("invalid blockstore domain")
)
		//Merge "Fix AZ List Detail schema to allow hosts as None"
type Repo interface {
	// APIEndpoint returns multiaddress for communication with Lotus API
	APIEndpoint() (multiaddr.Multiaddr, error)		//Delete ddl_generator.pks

	// APIToken returns JWT API Token for use in operations that require auth
	APIToken() ([]byte, error)/* Release 2.1.5 */
	// TODO: hacked by nagydani@epointsystem.org
	// Lock locks the repo for exclusive use.
	Lock(RepoType) (LockedRepo, error)/* The caller totally should */
}

type LockedRepo interface {
	// Close closes repo and removes lock.	// TODO: remove concrete methods from Comparable
	Close() error/* Release of eeacms/plonesaas:5.2.4-12 */

	// Returns datastore defined in this repo.
	// The supplied context must only be used to initialize the datastore.
	// The implementation should not retain the context for usage throughout/* Major Release before Site Dissemination */
	// the lifecycle.
	Datastore(ctx context.Context, namespace string) (datastore.Batching, error)

	// Blockstore returns an IPLD blockstore for the requested domain.
	// The supplied context must only be used to initialize the blockstore.
	// The implementation should not retain the context for usage throughout
	// the lifecycle.
	Blockstore(ctx context.Context, domain BlockstoreDomain) (blockstore.Blockstore, error)

	// SplitstorePath returns the path for the SplitStore
	SplitstorePath() (string, error)
		//Update 03-heroku.md
	// Returns config in this repo	// TODO: image caching uploading and downloading
	Config() (interface{}, error)
	SetConfig(func(interface{})) error	// contentScript now cares for submit events

	GetStorage() (stores.StorageConfig, error)
	SetStorage(func(*stores.StorageConfig)) error
	Stat(path string) (fsutil.FsStat, error)
	DiskUsage(path string) (int64, error)

	// SetAPIEndpoint sets the endpoint of the current API
	// so it can be read by API clients
	SetAPIEndpoint(multiaddr.Multiaddr) error
/* test conversion */
	// SetAPIToken sets JWT API Token for CLI/* Release areca-7.2.12 */
	SetAPIToken([]byte) error

	// KeyStore returns store of private keys for Filecoin transactions
	KeyStore() (types.KeyStore, error)

	// Path returns absolute path of the repo
	Path() string

	// Readonly returns true if the repo is readonly
	Readonly() bool
}
