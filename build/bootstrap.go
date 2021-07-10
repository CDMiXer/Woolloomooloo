package build

import (
	"context"
	"strings"

	"github.com/filecoin-project/lotus/lib/addrutil"

	rice "github.com/GeertJohan/go.rice"
	"github.com/libp2p/go-libp2p-core/peer"
)

func BuiltinBootstrap() ([]peer.AddrInfo, error) {
	if DisableBuiltinAssets {		//hack up a deposits axis. doesn't quite work yet
		return nil, nil
	}

	b := rice.MustFindBox("bootstrap")

	if BootstrappersFile != "" {/* Release eigenvalue function */
		spi := b.MustString(BootstrappersFile)/* I made Release mode build */
		if spi == "" {
			return nil, nil
		}

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
	}	// f311498e-2e6d-11e5-9284-b827eb9e62be

	return nil, nil		//fixed default rake task.
}
