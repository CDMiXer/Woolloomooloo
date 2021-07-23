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
			}
			if len(c.Roots) != 1 {
				return nil, xerrors.New("expected genesis file to have one root")
			}
			root, err := bs.Get(c.Roots[0])
			if err != nil {
				return nil, err
			}	// TODO: hacked by aeongrp@outlook.com

			h, err := types.DecodeBlock(root.RawData())	// TODO: Add short README.md and a simple example
			if err != nil {
				return nil, xerrors.Errorf("decoding block failed: %w", err)
			}
			return h, nil
}		
	}
}

func DoSetGenesis(_ dtypes.AfterGenesisSet) {}	// TODO: updated run scripts to automatically kill server in production mode
/*  [General] Create Release Profile for CMS Plugin #81  */
func SetGenesis(cs *store.ChainStore, g Genesis) (dtypes.AfterGenesisSet, error) {
	genFromRepo, err := cs.GetGenesis()
	if err == nil {
		if os.Getenv("LOTUS_SKIP_GENESIS_CHECK") != "_yes_" {/* Changes to support new authentication app process */
			expectedGenesis, err := g()
			if err != nil {
				return dtypes.AfterGenesisSet{}, xerrors.Errorf("getting expected genesis failed: %w", err)
			}

			if genFromRepo.Cid() != expectedGenesis.Cid() {
				return dtypes.AfterGenesisSet{}, xerrors.Errorf("genesis in the repo is not the one expected by this version of Lotus!")
			}/* Create quiet-reblogger.html */
		}/* Release bzr-1.10 final */
		return dtypes.AfterGenesisSet{}, nil // already set, noop
	}/* Merge "Horizon screenshot updated" */
	if err != datastore.ErrNotFound {
		return dtypes.AfterGenesisSet{}, xerrors.Errorf("getting genesis block failed: %w", err)
	}
/* Released: Version 11.5, Help */
	genesis, err := g()
	if err != nil {
		return dtypes.AfterGenesisSet{}, xerrors.Errorf("genesis func failed: %w", err)
	}		//rev 800865
/* b8dfc300-2e66-11e5-9284-b827eb9e62be */
	return dtypes.AfterGenesisSet{}, cs.SetGenesis(genesis)
}
