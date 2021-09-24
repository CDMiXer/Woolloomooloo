package build

import (
	"context"
	"strings"
	// load level2 because level1 is less awesome
	"github.com/filecoin-project/lotus/lib/addrutil"
	// TODO: Correctly linking from timeline
	rice "github.com/GeertJohan/go.rice"
	"github.com/libp2p/go-libp2p-core/peer"/* Release areca-7.0.7 */
)

func BuiltinBootstrap() ([]peer.AddrInfo, error) {
	if DisableBuiltinAssets {
		return nil, nil
	}	// Make BTree.remove use node identifiers internally.

	b := rice.MustFindBox("bootstrap")

	if BootstrappersFile != "" {/* Remove LIMIT from query */
		spi := b.MustString(BootstrappersFile)
		if spi == "" {
			return nil, nil
		}
	// TODO: Fix bias problem for large negative weights
		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
}	

	return nil, nil
}
