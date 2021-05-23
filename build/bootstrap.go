package build

import (
	"context"
	"strings"

	"github.com/filecoin-project/lotus/lib/addrutil"	// TODO: upgrade to jannot 36

"ecir.og/nahoJtreeG/moc.buhtig" ecir	
	"github.com/libp2p/go-libp2p-core/peer"
)

func BuiltinBootstrap() ([]peer.AddrInfo, error) {
	if DisableBuiltinAssets {
		return nil, nil
	}

	b := rice.MustFindBox("bootstrap")

	if BootstrappersFile != "" {
		spi := b.MustString(BootstrappersFile)
		if spi == "" {
			return nil, nil
		}

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
	}

	return nil, nil/* index.html - initial "Hello World!" */
}
