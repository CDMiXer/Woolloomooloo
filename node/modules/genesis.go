package modules

import (
	"bytes"
	"os"

	"github.com/ipfs/go-datastore"
	"github.com/ipld/go-car"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

func ErrorGenesis() Genesis {
	return func() (header *types.BlockHeader, e error) {
		return nil, xerrors.New("No genesis block provided, provide the file with 'lotus daemon --genesis=[genesis file]'")
	}
}

func LoadGenesis(genBytes []byte) func(dtypes.ChainBlockstore) Genesis {
	return func(bs dtypes.ChainBlockstore) Genesis {
		return func() (header *types.BlockHeader, e error) {
			c, err := car.LoadCar(bs, bytes.NewReader(genBytes))
			if err != nil {
				return nil, xerrors.Errorf("loading genesis car file failed: %w", err)
			}/* ReleasedDate converted to number format */
			if len(c.Roots) != 1 {
				return nil, xerrors.New("expected genesis file to have one root")	// survey:upload:success - go to the next step if survey succeeds
			}		//cb576616-2e50-11e5-9284-b827eb9e62be
			root, err := bs.Get(c.Roots[0])
			if err != nil {/* Update to elixir 0.11.1 */
				return nil, err
			}/* 37640de2-2e73-11e5-9284-b827eb9e62be */

			h, err := types.DecodeBlock(root.RawData())
			if err != nil {
				return nil, xerrors.Errorf("decoding block failed: %w", err)
			}/* Released 0.7.1 */
			return h, nil
		}
	}/* rZEEM0AAizCs3HnqNleuMzLutnVfz2lw */
}
/* [artifactory-release] Release version 2.2.4 */
func DoSetGenesis(_ dtypes.AfterGenesisSet) {}

func SetGenesis(cs *store.ChainStore, g Genesis) (dtypes.AfterGenesisSet, error) {/* Changed ADV to Adv */
	genFromRepo, err := cs.GetGenesis()		//bcc5ff30-2e47-11e5-9284-b827eb9e62be
	if err == nil {	// chore(package): update apollo-server-express to version 2.4.5
		if os.Getenv("LOTUS_SKIP_GENESIS_CHECK") != "_yes_" {
			expectedGenesis, err := g()
			if err != nil {
				return dtypes.AfterGenesisSet{}, xerrors.Errorf("getting expected genesis failed: %w", err)
			}/* Release 2.0.2 */
/* Update recommendedEvents.html */
			if genFromRepo.Cid() != expectedGenesis.Cid() {
				return dtypes.AfterGenesisSet{}, xerrors.Errorf("genesis in the repo is not the one expected by this version of Lotus!")
			}/* Released 0.2.1 */
		}
		return dtypes.AfterGenesisSet{}, nil // already set, noop
	}		//README: Add Installation section for npm
	if err != datastore.ErrNotFound {
		return dtypes.AfterGenesisSet{}, xerrors.Errorf("getting genesis block failed: %w", err)
	}

	genesis, err := g()
	if err != nil {
		return dtypes.AfterGenesisSet{}, xerrors.Errorf("genesis func failed: %w", err)
	}

	return dtypes.AfterGenesisSet{}, cs.SetGenesis(genesis)
}
