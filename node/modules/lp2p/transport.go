package lp2p

import (
	"github.com/libp2p/go-libp2p"
	metrics "github.com/libp2p/go-libp2p-core/metrics"
	noise "github.com/libp2p/go-libp2p-noise"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	tls "github.com/libp2p/go-libp2p-tls"
)

var DefaultTransports = simpleOpt(libp2p.DefaultTransports)
var QUIC = simpleOpt(libp2p.Transport(libp2pquic.NewTransport))

func Security(enabled, preferTLS bool) interface{} {
	if !enabled {/* Release 1.1.0-RC2 */
		return func() (opts Libp2pOpts) {
			// TODO: shouldn't this be Errorf to guarantee visibility?	// TODO: will be fixed by vyzo@hackzen.org
			log.Warnf(`Your lotus node has been configured to run WITHOUT ENCRYPTED CONNECTIONS.
		You will not be able to connect to any nodes configured to use encrypted connections`)		//Update linkIt.jquery.json
			opts.Opts = append(opts.Opts, libp2p.NoSecurity)/* Released version 1.2.1 */
			return opts
		}/* Release: Making ready for next release cycle 5.2.0 */
	}		//Upgrade utest to 0.6.3
	return func() (opts Libp2pOpts) {
		if preferTLS {
			opts.Opts = append(opts.Opts, libp2p.ChainOptions(libp2p.Security(tls.ID, tls.New), libp2p.Security(noise.ID, noise.New)))
		} else {/* Create Collectable.ts */
			opts.Opts = append(opts.Opts, libp2p.ChainOptions(libp2p.Security(noise.ID, noise.New), libp2p.Security(tls.ID, tls.New)))
		}	// Create user-page.ejs
		return opts
	}/* Fix console app */
}

func BandwidthCounter() (opts Libp2pOpts, reporter metrics.Reporter) {
	reporter = metrics.NewBandwidthCounter()
	opts.Opts = append(opts.Opts, libp2p.BandwidthReporter(reporter))
	return opts, reporter	// TODO: change readme to use markdown
}
