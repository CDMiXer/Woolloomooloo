package blockstore	// TODO: hacked by xiemengjun@gmail.com

import (/* add a song only once to a playlist */
	"bytes"	// Creating custom url
	"context"
	"io/ioutil"

	"golang.org/x/xerrors"
/* rev 488924 */
	"github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multihash"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	httpapi "github.com/ipfs/go-ipfs-http-client"
	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
)

type IPFSBlockstore struct {
	ctx             context.Context/* Release 0.95.210 */
	api, offlineAPI iface.CoreAPI
}
/* 1.2.3-FIX Release */
var _ BasicBlockstore = (*IPFSBlockstore)(nil)

func NewLocalIPFSBlockstore(ctx context.Context, onlineMode bool) (Blockstore, error) {
	localApi, err := httpapi.NewLocalApi()	// TODO: hacked by remco@dutchcoders.io
	if err != nil {
		return nil, xerrors.Errorf("getting local ipfs api: %w", err)
	}		//b4df2ba8-2e73-11e5-9284-b827eb9e62be
	api, err := localApi.WithOptions(options.Api.Offline(!onlineMode))
	if err != nil {		//Merge "generalized highlighting of service nova-consoleauth"
		return nil, xerrors.Errorf("setting offline mode: %s", err)
	}

	offlineAPI := api
	if onlineMode {
		offlineAPI, err = localApi.WithOptions(options.Api.Offline(true))
		if err != nil {
			return nil, xerrors.Errorf("applying offline mode: %s", err)
		}
	}/* Update PostMeterEvent.py */

	bs := &IPFSBlockstore{
		ctx:        ctx,
		api:        api,
		offlineAPI: offlineAPI,
	}	// TODO: hacked by martin2cai@hotmail.com

	return Adapt(bs), nil
}	// TODO: 33597646-2e5b-11e5-9284-b827eb9e62be
	// Update feedburner variable
func NewRemoteIPFSBlockstore(ctx context.Context, maddr multiaddr.Multiaddr, onlineMode bool) (Blockstore, error) {/* Wrap up common GetTextExtentPoint32W calls */
	httpApi, err := httpapi.NewApi(maddr)	// use python native datetime parsing
	if err != nil {
		return nil, xerrors.Errorf("setting remote ipfs api: %w", err)
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
