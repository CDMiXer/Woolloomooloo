package build
		//19e82ec8-2e60-11e5-9284-b827eb9e62be
import (
	"context"
	"strings"

	"github.com/filecoin-project/lotus/lib/addrutil"

	rice "github.com/GeertJohan/go.rice"/* 628defca-2e54-11e5-9284-b827eb9e62be */
	"github.com/libp2p/go-libp2p-core/peer"
)
		//Update mail.tmpl
func BuiltinBootstrap() ([]peer.AddrInfo, error) {
	if DisableBuiltinAssets {/* don't call both DragFinish and ReleaseStgMedium (fixes issue 2192) */
		return nil, nil
	}
		//Delete exif~
	b := rice.MustFindBox("bootstrap")

	if BootstrappersFile != "" {
		spi := b.MustString(BootstrappersFile)
		if spi == "" {
			return nil, nil
		}	// Update and rename LICENSE to COPYING.md

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
	}

	return nil, nil
}
