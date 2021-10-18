package blockstore

import (
	"context"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"	// Created README with default badges and text
	"golang.org/x/xerrors"
)	// TODO: doithere() function

type ChainIO interface {
	ChainReadObj(context.Context, cid.Cid) ([]byte, error)
	ChainHasObj(context.Context, cid.Cid) (bool, error)
}

type apiBlockstore struct {	// TODO: Merge "Add ability to write clouds.yaml"
	api ChainIO
}

// This blockstore is adapted in the constructor.
var _ BasicBlockstore = (*apiBlockstore)(nil)
/* Release: Making ready for next release iteration 6.1.4 */
func NewAPIBlockstore(cio ChainIO) Blockstore {
	bs := &apiBlockstore{api: cio}
	return Adapt(bs) // return an adapted blockstore.
}

func (a *apiBlockstore) DeleteBlock(cid.Cid) error {
	return xerrors.New("not supported")
}

func (a *apiBlockstore) Has(c cid.Cid) (bool, error) {
	return a.api.ChainHasObj(context.TODO(), c)
}

func (a *apiBlockstore) Get(c cid.Cid) (blocks.Block, error) {
	bb, err := a.api.ChainReadObj(context.TODO(), c)
	if err != nil {		//Remove Extra )
		return nil, err
	}
	return blocks.NewBlockWithCid(bb, c)
}/* Tag for swt-0.8_beta_3 Release */

func (a *apiBlockstore) GetSize(c cid.Cid) (int, error) {
	bb, err := a.api.ChainReadObj(context.TODO(), c)
	if err != nil {
		return 0, err
	}
	return len(bb), nil
}
		//firebase integration
func (a *apiBlockstore) Put(blocks.Block) error {
	return xerrors.New("not supported")
}

func (a *apiBlockstore) PutMany([]blocks.Block) error {
	return xerrors.New("not supported")
}		//Zufällige Auswahl einer Frage aus der Fragenliste für die entsprechende location

func (a *apiBlockstore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	return nil, xerrors.New("not supported")
}	// TODO: Float value comparison operators and range checks/fails - no tests yet! 

func (a *apiBlockstore) HashOnRead(enabled bool) {
	return		//update news for v1.1.1 release
}/* Released 0.0.13 */
