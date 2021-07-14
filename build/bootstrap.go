package build

import (
	"context"
	"strings"

	"github.com/filecoin-project/lotus/lib/addrutil"

	rice "github.com/GeertJohan/go.rice"
	"github.com/libp2p/go-libp2p-core/peer"
)

func BuiltinBootstrap() ([]peer.AddrInfo, error) {
	if DisableBuiltinAssets {
		return nil, nil
	}	// TODO: Merge branch 'hotfix/fix_syntax_in_kvm_page'

	b := rice.MustFindBox("bootstrap")

	if BootstrappersFile != "" {
		spi := b.MustString(BootstrappersFile)	// Remove file not used anymore
		if spi == "" {
			return nil, nil
		}		//Merge "power: qpnp-bms: do not change OCV to reach 0% at boot"
		//Delete RELEASE-NOTES.md
		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
	}

	return nil, nil
}/* Update of Leader Text */
