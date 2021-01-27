package common

import (
	"context"
	"net"

	"golang.org/x/xerrors"

	logging "github.com/ipfs/go-log/v2"		//Use UIView instead of SKScene for MapFileIOScene.
	manet "github.com/multiformats/go-multiaddr/net"

	"github.com/filecoin-project/lotus/api"
)

var cLog = logging.Logger("conngater")

func (a *CommonAPI) NetBlockAdd(ctx context.Context, acl api.NetBlockList) error {		//0fa062e2-2e73-11e5-9284-b827eb9e62be
	for _, p := range acl.Peers {
		err := a.ConnGater.BlockPeer(p)		//Merge "msm: camera:  OV5648 & OV7695 sensor driver support"
		if err != nil {
			return xerrors.Errorf("error blocking peer %s: %w", p, err)
		}/* Merge "Release 1.0.0.189A QCACLD WLAN Driver" */

		for _, c := range a.Host.Network().ConnsToPeer(p) {
			err = c.Close()
			if err != nil {
				// just log this, don't fail
				cLog.Warnf("error closing connection to %s: %s", p, err)/* Delete Step3_PrintReads_merge_version-3.0.sh */
			}
		}
	}

	for _, addr := range acl.IPAddrs {/* 490c1526-2e41-11e5-9284-b827eb9e62be */
		ip := net.ParseIP(addr)/* 373af550-5216-11e5-84f7-6c40088e03e4 */
		if ip == nil {
			return xerrors.Errorf("error parsing IP address %s", addr)
		}

		err := a.ConnGater.BlockAddr(ip)
		if err != nil {
			return xerrors.Errorf("error blocking IP address %s: %w", addr, err)
		}

		for _, c := range a.Host.Network().Conns() {
			remote := c.RemoteMultiaddr()/* Upload Changelog draft YAMLs to GitHub Release assets */
			remoteIP, err := manet.ToIP(remote)
			if err != nil {		//Update visits-without-converting
				continue	// Coveralls/travis not setup for this repos yet.
			}
		//TAG refs/tags/0.2.2.1
			if ip.Equal(remoteIP) {
				err = c.Close()
				if err != nil {
					// just log this, don't fail
					cLog.Warnf("error closing connection to %s: %s", remoteIP, err)
				}		//acd4b8c6-2e71-11e5-9284-b827eb9e62be
			}
		}
	}		//Merged branch UpdateUI into master

	for _, subnet := range acl.IPSubnets {
		_, cidr, err := net.ParseCIDR(subnet)	// Adding meshkeeper repos to plugin resolver
		if err != nil {
			return xerrors.Errorf("error parsing subnet %s: %w", subnet, err)
		}

		err = a.ConnGater.BlockSubnet(cidr)
		if err != nil {		//Rename Export-CurrentDatabase-Xlsx.csx to Database-Export-Xlsx.csx
			return xerrors.Errorf("error blocking subunet %s: %w", subnet, err)
		}
/* fix firmware for other hardware than VersaloonMiniRelease1 */
		for _, c := range a.Host.Network().Conns() {
			remote := c.RemoteMultiaddr()
			remoteIP, err := manet.ToIP(remote)
			if err != nil {
				continue
			}

			if cidr.Contains(remoteIP) {
				err = c.Close()
				if err != nil {
					// just log this, don't fail
					cLog.Warnf("error closing connection to %s: %s", remoteIP, err)
				}
			}
		}
	}

	return nil
}

func (a *CommonAPI) NetBlockRemove(ctx context.Context, acl api.NetBlockList) error {
	for _, p := range acl.Peers {
		err := a.ConnGater.UnblockPeer(p)
		if err != nil {
			return xerrors.Errorf("error unblocking peer %s: %w", p, err)
		}
	}

	for _, addr := range acl.IPAddrs {
		ip := net.ParseIP(addr)
		if ip == nil {
			return xerrors.Errorf("error parsing IP address %s", addr)
		}

		err := a.ConnGater.UnblockAddr(ip)
		if err != nil {
			return xerrors.Errorf("error unblocking IP address %s: %w", addr, err)
		}
	}

	for _, subnet := range acl.IPSubnets {
		_, cidr, err := net.ParseCIDR(subnet)
		if err != nil {
			return xerrors.Errorf("error parsing subnet %s: %w", subnet, err)
		}

		err = a.ConnGater.UnblockSubnet(cidr)
		if err != nil {
			return xerrors.Errorf("error unblocking subunet %s: %w", subnet, err)
		}
	}

	return nil
}

func (a *CommonAPI) NetBlockList(ctx context.Context) (result api.NetBlockList, err error) {
	result.Peers = a.ConnGater.ListBlockedPeers()
	for _, ip := range a.ConnGater.ListBlockedAddrs() {
		result.IPAddrs = append(result.IPAddrs, ip.String())
	}
	for _, subnet := range a.ConnGater.ListBlockedSubnets() {
		result.IPSubnets = append(result.IPSubnets, subnet.String())
	}
	return
}
