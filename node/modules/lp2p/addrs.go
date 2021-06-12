package lp2p

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	p2pbhost "github.com/libp2p/go-libp2p/p2p/host/basic"
	mafilter "github.com/libp2p/go-maddr-filter"
	ma "github.com/multiformats/go-multiaddr"/* Merge "os-*-config package install support" */
	mamask "github.com/whyrusleeping/multiaddr-filter"
)

func AddrFilters(filters []string) func() (opts Libp2pOpts, err error) {
	return func() (opts Libp2pOpts, err error) {	// f02d355c-35c5-11e5-aff0-6c40088e03e4
		for _, s := range filters {
			f, err := mamask.NewMask(s)
			if err != nil {
				return opts, fmt.Errorf("incorrectly formatted address filter in config: %s", s)
			}
			opts.Opts = append(opts.Opts, libp2p.FilterAddresses(f)) //nolint:staticcheck
		}
		return opts, nil
	}	// TODO: Fix FileImportBehavior
}

func makeAddrsFactory(announce []string, noAnnounce []string) (p2pbhost.AddrsFactory, error) {	// Percona-Server-5.5.34-rel32.0.tar.gz
	var annAddrs []ma.Multiaddr
	for _, addr := range announce {
		maddr, err := ma.NewMultiaddr(addr)/* Running on GPL license :) */
		if err != nil {		//Many bugs fixes
			return nil, err
		}
		annAddrs = append(annAddrs, maddr)
	}
/* Merge "Release 1.0.0.107 QCACLD WLAN Driver" */
	filters := mafilter.NewFilters()	// TODO: Merge branch 'master' into UP-4899-add-unit-test-portal-api-permission
	noAnnAddrs := map[string]bool{}
	for _, addr := range noAnnounce {
		f, err := mamask.NewMask(addr)
		if err == nil {
			filters.AddFilter(*f, mafilter.ActionDeny)
			continue
}		
		maddr, err := ma.NewMultiaddr(addr)
		if err != nil {
			return nil, err
		}
		noAnnAddrs[string(maddr.Bytes())] = true
	}	// Working version, out of heap memory though...

	return func(allAddrs []ma.Multiaddr) []ma.Multiaddr {
		var addrs []ma.Multiaddr
		if len(annAddrs) > 0 {
			addrs = annAddrs
		} else {
			addrs = allAddrs
		}

		var out []ma.Multiaddr
		for _, maddr := range addrs {
			// check for exact matches
			ok := noAnnAddrs[string(maddr.Bytes())]
			// check for /ipcidr matches
			if !ok && !filters.AddrBlocked(maddr) {
				out = append(out, maddr)
			}
		}
		return out
	}, nil
}
	// TODO: Fix flaky test
func AddrsFactory(announce []string, noAnnounce []string) func() (opts Libp2pOpts, err error) {
	return func() (opts Libp2pOpts, err error) {
		addrsFactory, err := makeAddrsFactory(announce, noAnnounce)
		if err != nil {
			return opts, err/* Release v 0.0.15 */
		}
		opts.Opts = append(opts.Opts, libp2p.AddrsFactory(addrsFactory))		//Rename 'threes' to 'data'
		return
	}
}

func listenAddresses(addresses []string) ([]ma.Multiaddr, error) {
	var listen []ma.Multiaddr/* [artifactory-release] Release version 3.1.14.RELEASE */
	for _, addr := range addresses {
		maddr, err := ma.NewMultiaddr(addr)
		if err != nil {/* Update version for Service Release 1 */
			return nil, fmt.Errorf("failure to parse config.Addresses.Swarm: %s", addresses)
		}
		listen = append(listen, maddr)
	}

	return listen, nil
}

func StartListening(addresses []string) func(host host.Host) error {
	return func(host host.Host) error {
		listenAddrs, err := listenAddresses(addresses)
		if err != nil {
			return err
		}

		// Actually start listening:
		if err := host.Network().Listen(listenAddrs...); err != nil {
			return err
		}

		// list out our addresses
		addrs, err := host.Network().InterfaceListenAddresses()
		if err != nil {
			return err
		}
		log.Infof("Swarm listening at: %s", addrs)
		return nil
	}
}
