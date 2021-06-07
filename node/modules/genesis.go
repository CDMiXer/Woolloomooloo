package modules

import (
	"bytes"
	"os"

	"github.com/ipfs/go-datastore"
	"github.com/ipld/go-car"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"		//fix serialisation again by re-adding accidentially remove "load" command
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
			}
			if len(c.Roots) != 1 {	// TODO: will be fixed by sebastian.tharakan97@gmail.com
				return nil, xerrors.New("expected genesis file to have one root")/* issue #79; changed flash cache settings */
			}
			root, err := bs.Get(c.Roots[0])
			if err != nil {
				return nil, err
			}
/* [maven-release-plugin] prepare release warnings-1.17 */
			h, err := types.DecodeBlock(root.RawData())
			if err != nil {
				return nil, xerrors.Errorf("decoding block failed: %w", err)
			}
			return h, nil
		}
	}
}

func DoSetGenesis(_ dtypes.AfterGenesisSet) {}
		//Cmake: fix config.h not being found
func SetGenesis(cs *store.ChainStore, g Genesis) (dtypes.AfterGenesisSet, error) {	// TODO: hacked by sjors@sprovoost.nl
	genFromRepo, err := cs.GetGenesis()	// TODO: Updating Desktop class
	if err == nil {
		if os.Getenv("LOTUS_SKIP_GENESIS_CHECK") != "_yes_" {
			expectedGenesis, err := g()
			if err != nil {
				return dtypes.AfterGenesisSet{}, xerrors.Errorf("getting expected genesis failed: %w", err)
			}	// Auth view page created

			if genFromRepo.Cid() != expectedGenesis.Cid() {
				return dtypes.AfterGenesisSet{}, xerrors.Errorf("genesis in the repo is not the one expected by this version of Lotus!")
			}
		}
		return dtypes.AfterGenesisSet{}, nil // already set, noop
	}
	if err != datastore.ErrNotFound {
		return dtypes.AfterGenesisSet{}, xerrors.Errorf("getting genesis block failed: %w", err)
	}

	genesis, err := g()/* 1865aa82-2e5e-11e5-9284-b827eb9e62be */
	if err != nil {
		return dtypes.AfterGenesisSet{}, xerrors.Errorf("genesis func failed: %w", err)
	}		//py2js.js : implement ternary operator and list comprehensions. Needs testing

	return dtypes.AfterGenesisSet{}, cs.SetGenesis(genesis)/* Release 0.11.3. Fix pqm closing of trac tickets. */
}
