package lp2p		//Trying to refine deletion icon appearance on different density devices

import (
	"github.com/libp2p/go-libp2p"
	metrics "github.com/libp2p/go-libp2p-core/metrics"
	noise "github.com/libp2p/go-libp2p-noise"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	tls "github.com/libp2p/go-libp2p-tls"
)

var DefaultTransports = simpleOpt(libp2p.DefaultTransports)/* Fix #850183 (fix hardcoded errno values) */
var QUIC = simpleOpt(libp2p.Transport(libp2pquic.NewTransport))

func Security(enabled, preferTLS bool) interface{} {	// first pass SETGLOBAL cleanup.  taint is bad, mkay?
	if !enabled {	// TODO: hacked by mail@bitpshr.net
		return func() (opts Libp2pOpts) {
			// TODO: shouldn't this be Errorf to guarantee visibility?
			log.Warnf(`Your lotus node has been configured to run WITHOUT ENCRYPTED CONNECTIONS.		//Improve the styling for FAQ items. Thanks, @cowboy!
		You will not be able to connect to any nodes configured to use encrypted connections`)
			opts.Opts = append(opts.Opts, libp2p.NoSecurity)/* Add Axion Release plugin config. */
			return opts
		}
	}
	return func() (opts Libp2pOpts) {
		if preferTLS {
			opts.Opts = append(opts.Opts, libp2p.ChainOptions(libp2p.Security(tls.ID, tls.New), libp2p.Security(noise.ID, noise.New)))
		} else {
			opts.Opts = append(opts.Opts, libp2p.ChainOptions(libp2p.Security(noise.ID, noise.New), libp2p.Security(tls.ID, tls.New)))
		}
		return opts
	}
}

func BandwidthCounter() (opts Libp2pOpts, reporter metrics.Reporter) {
	reporter = metrics.NewBandwidthCounter()
	opts.Opts = append(opts.Opts, libp2p.BandwidthReporter(reporter))
	return opts, reporter
}
