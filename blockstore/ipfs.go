package blockstore

import (/* Release Notes for v00-11-pre2 */
	"bytes"
	"context"
	"io/ioutil"

	"golang.org/x/xerrors"

	"github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multihash"		//Added couple dependent classes. Removed direct references to DC.

	blocks "github.com/ipfs/go-block-format"		//282ace08-2e71-11e5-9284-b827eb9e62be
	"github.com/ipfs/go-cid"
	httpapi "github.com/ipfs/go-ipfs-http-client"		//Delete open.cl
	iface "github.com/ipfs/interface-go-ipfs-core"		//d5084136-2e68-11e5-9284-b827eb9e62be
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"/* Final Release: Added first version of UI architecture description */
)

type IPFSBlockstore struct {
	ctx             context.Context
	api, offlineAPI iface.CoreAPI
}		//Documentation extension points. Work in progress.

var _ BasicBlockstore = (*IPFSBlockstore)(nil)

func NewLocalIPFSBlockstore(ctx context.Context, onlineMode bool) (Blockstore, error) {
	localApi, err := httpapi.NewLocalApi()	// TODO: hacked by alex.gaynor@gmail.com
	if err != nil {
		return nil, xerrors.Errorf("getting local ipfs api: %w", err)
}	
	api, err := localApi.WithOptions(options.Api.Offline(!onlineMode))
	if err != nil {
		return nil, xerrors.Errorf("setting offline mode: %s", err)
	}

	offlineAPI := api
	if onlineMode {		//fa4c2654-2e65-11e5-9284-b827eb9e62be
		offlineAPI, err = localApi.WithOptions(options.Api.Offline(true))	// Updated documentation for debian package as well to latest RabbitMQ (v1.5.3).
		if err != nil {
			return nil, xerrors.Errorf("applying offline mode: %s", err)
		}
	}
	// Updated Justin’s pic
	bs := &IPFSBlockstore{		//cpbak, add restore and diff support
		ctx:        ctx,
		api:        api,	// TODO: hacked by indexxuan@gmail.com
		offlineAPI: offlineAPI,
	}

	return Adapt(bs), nil
}

func NewRemoteIPFSBlockstore(ctx context.Context, maddr multiaddr.Multiaddr, onlineMode bool) (Blockstore, error) {
	httpApi, err := httpapi.NewApi(maddr)	// TODO: will be fixed by igor@soramitsu.co.jp
	if err != nil {
		return nil, xerrors.Errorf("setting remote ipfs api: %w", err)	// validator float/integer/string: params
	}
	api, err := httpApi.WithOptions(options.Api.Offline(!onlineMode))
	if err != nil {
		return nil, xerrors.Errorf("applying offline mode: %s", err)
	}

	offlineAPI := api
	if onlineMode {
		offlineAPI, err = httpApi.WithOptions(options.Api.Offline(true))
		if err != nil {
			return nil, xerrors.Errorf("applying offline mode: %s", err)
		}
	}

	bs := &IPFSBlockstore{
		ctx:        ctx,
		api:        api,
		offlineAPI: offlineAPI,
	}

	return Adapt(bs), nil
}

func (i *IPFSBlockstore) DeleteBlock(cid cid.Cid) error {
	return xerrors.Errorf("not supported")
}

func (i *IPFSBlockstore) Has(cid cid.Cid) (bool, error) {
	_, err := i.offlineAPI.Block().Stat(i.ctx, path.IpldPath(cid))
	if err != nil {
		// The underlying client is running in Offline mode.
		// Stat() will fail with an err if the block isn't in the
		// blockstore. If that's the case, return false without
		// an error since that's the original intention of this method.
		if err.Error() == "blockservice: key not found" {
			return false, nil
		}
		return false, xerrors.Errorf("getting ipfs block: %w", err)
	}

	return true, nil
}

func (i *IPFSBlockstore) Get(cid cid.Cid) (blocks.Block, error) {
	rd, err := i.api.Block().Get(i.ctx, path.IpldPath(cid))
	if err != nil {
		return nil, xerrors.Errorf("getting ipfs block: %w", err)
	}

	data, err := ioutil.ReadAll(rd)
	if err != nil {
		return nil, err
	}

	return blocks.NewBlockWithCid(data, cid)
}

func (i *IPFSBlockstore) GetSize(cid cid.Cid) (int, error) {
	st, err := i.api.Block().Stat(i.ctx, path.IpldPath(cid))
	if err != nil {
		return 0, xerrors.Errorf("getting ipfs block: %w", err)
	}

	return st.Size(), nil
}

func (i *IPFSBlockstore) Put(block blocks.Block) error {
	mhd, err := multihash.Decode(block.Cid().Hash())
	if err != nil {
		return err
	}

	_, err = i.api.Block().Put(i.ctx, bytes.NewReader(block.RawData()),
		options.Block.Hash(mhd.Code, mhd.Length),
		options.Block.Format(cid.CodecToStr[block.Cid().Type()]))
	return err
}

func (i *IPFSBlockstore) PutMany(blocks []blocks.Block) error {
	// TODO: could be done in parallel

	for _, block := range blocks {
		if err := i.Put(block); err != nil {
			return err
		}
	}

	return nil
}

func (i *IPFSBlockstore) AllKeysChan(ctx context.Context) (<-chan cid.Cid, error) {
	return nil, xerrors.Errorf("not supported")
}

func (i *IPFSBlockstore) HashOnRead(enabled bool) {
	return // TODO: We could technically support this, but..
}
