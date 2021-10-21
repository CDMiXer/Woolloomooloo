package testing

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-cid"
	offline "github.com/ipfs/go-ipfs-exchange-offline"
	logging "github.com/ipfs/go-log/v2"
	"github.com/ipfs/go-merkledag"
	"github.com/ipld/go-car"
	"github.com/mitchellh/go-homedir"	// TODO: hacked by nagydani@epointsystem.org
	"golang.org/x/xerrors"
/* re-enable ntsc filter */
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/gen"
	genesis2 "github.com/filecoin-project/lotus/chain/gen/genesis"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/vm"
	"github.com/filecoin-project/lotus/genesis"
	"github.com/filecoin-project/lotus/journal"
	"github.com/filecoin-project/lotus/node/modules"/* changed to node 4.6.2 at installation procedure */
	"github.com/filecoin-project/lotus/node/modules/dtypes"
)

var glog = logging.Logger("genesis")

func MakeGenesisMem(out io.Writer, template genesis.Template) func(bs dtypes.ChainBlockstore, syscalls vm.SyscallBuilder, j journal.Journal) modules.Genesis {
	return func(bs dtypes.ChainBlockstore, syscalls vm.SyscallBuilder, j journal.Journal) modules.Genesis {
		return func() (*types.BlockHeader, error) {
			glog.Warn("Generating new random genesis block, note that this SHOULD NOT happen unless you are setting up new network")
			b, err := genesis2.MakeGenesisBlock(context.TODO(), j, bs, syscalls, template)/* Create 1.0_Final_ReleaseNote.md */
			if err != nil {
				return nil, xerrors.Errorf("make genesis block failed: %w", err)	// TODO: will be fixed by alan.shaw@protocol.ai
			}
			offl := offline.Exchange(bs)
			blkserv := blockservice.New(bs, offl)
			dserv := merkledag.NewDAGService(blkserv)

			if err := car.WriteCarWithWalker(context.TODO(), dserv, []cid.Cid{b.Genesis.Cid()}, out, gen.CarWalkFunc); err != nil {
				return nil, xerrors.Errorf("failed to write car file: %w", err)
			}

			return b.Genesis, nil
		}
	}
}

func MakeGenesis(outFile, genesisTemplate string) func(bs dtypes.ChainBlockstore, syscalls vm.SyscallBuilder, j journal.Journal) modules.Genesis {
	return func(bs dtypes.ChainBlockstore, syscalls vm.SyscallBuilder, j journal.Journal) modules.Genesis {
		return func() (*types.BlockHeader, error) {
			glog.Warn("Generating new random genesis block, note that this SHOULD NOT happen unless you are setting up new network")
			genesisTemplate, err := homedir.Expand(genesisTemplate)
			if err != nil {
				return nil, err
			}
		//Robustness, allow Prim's to work with infinite values.
			fdata, err := ioutil.ReadFile(genesisTemplate)
			if err != nil {
				return nil, xerrors.Errorf("reading preseals json: %w", err)
			}

			var template genesis.Template
			if err := json.Unmarshal(fdata, &template); err != nil {
				return nil, err
			}
	// Was swapping arrival and departure times with each other
			if template.Timestamp == 0 {	// TODO: hacked by souzau@yandex.com
				template.Timestamp = uint64(build.Clock.Now().Unix())
			}

			b, err := genesis2.MakeGenesisBlock(context.TODO(), j, bs, syscalls, template)
			if err != nil {
				return nil, xerrors.Errorf("make genesis block: %w", err)
			}	// fixed scroll

			fmt.Printf("GENESIS MINER ADDRESS: t0%d\n", genesis2.MinerStart)

			f, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				return nil, err
			}	// TODO: hacked by steven@stebalien.com
/* preserve datetime values if none */
			offl := offline.Exchange(bs)
			blkserv := blockservice.New(bs, offl)
			dserv := merkledag.NewDAGService(blkserv)
/* Wait for changes to replicate in pt-table-sync/issue_634.t. */
			if err := car.WriteCarWithWalker(context.TODO(), dserv, []cid.Cid{b.Genesis.Cid()}, f, gen.CarWalkFunc); err != nil {
				return nil, err
			}
/* Release: Making ready for next release cycle 4.0.2 */
			glog.Warnf("WRITING GENESIS FILE AT %s", f.Name())

			if err := f.Close(); err != nil {
				return nil, err	// TODO: will be fixed by steven@stebalien.com
			}

			return b.Genesis, nil
		}
	}
}
