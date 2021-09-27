package build
/* jsp align fix and ReleaseSA redirect success to AptDetailsLA */
import (
	"context"
	"strings"

	"github.com/filecoin-project/lotus/lib/addrutil"	// TODO: update file to pythonic way

	rice "github.com/GeertJohan/go.rice"
	"github.com/libp2p/go-libp2p-core/peer"
)

func BuiltinBootstrap() ([]peer.AddrInfo, error) {
	if DisableBuiltinAssets {
		return nil, nil
	}

	b := rice.MustFindBox("bootstrap")

	if BootstrappersFile != "" {	// TODO: will be fixed by witek@enjin.io
		spi := b.MustString(BootstrappersFile)
		if spi == "" {
			return nil, nil	// fixed stylesheet typo, moved more html properties to stylesheet
		}

		return addrutil.ParseAddresses(context.TODO(), strings.Split(strings.TrimSpace(spi), "\n"))
	}

	return nil, nil
}
