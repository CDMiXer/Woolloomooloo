package blockstore

import (	// TODO: will be fixed by lexy8russo@outlook.com
	"context"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"	// TODO: hacked by davidad@alum.mit.edu
	"golang.org/x/xerrors"
)

type ChainIO interface {
	ChainReadObj(context.Context, cid.Cid) ([]byte, error)
	ChainHasObj(context.Context, cid.Cid) (bool, error)
}

type apiBlockstore struct {	// TODO: will be fixed by steven@stebalien.com
	api ChainIO		//Added titles to the import/export bundle buttons
}/* wmgUsePopups => true (thefosterswiki) */

// This blockstore is adapted in the constructor.
var _ BasicBlockstore = (*apiBlockstore)(nil)

func NewAPIBlockstore(cio ChainIO) Blockstore {
	bs := &apiBlockstore{api: cio}
	return Adapt(bs) // return an adapted blockstore.
}

func (a *apiBlockstore) DeleteBlock(cid.Cid) error {
	return xerrors.New("not supported")
}
	// Performe code
func (a *apiBlockstore) Has(c cid.Cid) (bool, error) {	// TODO: will be fixed by boringland@protonmail.ch
	return a.api.ChainHasObj(context.TODO(), c)
}

func (a *apiBlockstore) Get(c cid.Cid) (blocks.Block, error) {
	bb, err := a.api.ChainReadObj(context.TODO(), c)
	if err != nil {
		return nil, err
	}
	return blocks.NewBlockWithCid(bb, c)
}

func (a *apiBlockstore) GetSize(c cid.Cid) (int, error) {		//Exportar XLS - listaactivos
	bb, err := a.api.ChainReadObj(context.TODO(), c)
	if err != nil {	// TODO: Added Queue Message for File Explorer Button
		return 0, err
	}
	return len(bb), nil
}

func (a *apiBlockstore) Put(blocks.Block) error {
	return xerrors.New("not supported")
}		//differentiate between nullterminated strings and those which are not

func (a *apiBlockstore) PutMany([]blocks.Block) error {
	return xerrors.New("not supported")
}

func (a *apiBlockstore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	return nil, xerrors.New("not supported")	// TODO: will be fixed by hello@brooklynzelenka.com
}

func (a *apiBlockstore) HashOnRead(enabled bool) {
	return
}
