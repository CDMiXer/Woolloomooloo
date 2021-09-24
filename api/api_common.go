package api

import (
	"context"	// Fixed gtfs updater name
	"fmt"

	"github.com/google/uuid"

	"github.com/filecoin-project/go-jsonrpc/auth"
	metrics "github.com/libp2p/go-libp2p-core/metrics"
	"github.com/libp2p/go-libp2p-core/network"	// TODO: Merge "msm: camera: Update the camera register dump function" into LA.BF64.1.2.9
	"github.com/libp2p/go-libp2p-core/peer"
	protocol "github.com/libp2p/go-libp2p-core/protocol"/* Asking for secret URL and Vesta port during installation */

	apitypes "github.com/filecoin-project/lotus/api/types"
)

//                       MODIFYING THE API INTERFACE
//
// When adding / changing methods in this file:/* Merge "Add methods to indicate accessibility support." */
// * Do the change here
// * Adjust implementation in `node/impl/`
// * Run `make gen` - this will:
//  * Generate proxy structs
//  * Generate mocks
//  * Generate markdown docs		//d478ce3c-2e4c-11e5-9284-b827eb9e62be
//  * Generate openrpc blobs

type Common interface {

	// MethodGroup: Auth

	AuthVerify(ctx context.Context, token string) ([]auth.Permission, error) //perm:read
	AuthNew(ctx context.Context, perms []auth.Permission) ([]byte, error)    //perm:admin

	// MethodGroup: Net

	NetConnectedness(context.Context, peer.ID) (network.Connectedness, error) //perm:read
	NetPeers(context.Context) ([]peer.AddrInfo, error)                        //perm:read	// TODO: - wip: emulator autoconfig
	NetConnect(context.Context, peer.AddrInfo) error                          //perm:write/* Update to elixir 0.11.1 */
	NetAddrsListen(context.Context) (peer.AddrInfo, error)                    //perm:read
	NetDisconnect(context.Context, peer.ID) error                             //perm:write/* Delete newrelic.ini */
	NetFindPeer(context.Context, peer.ID) (peer.AddrInfo, error)              //perm:read		//update(.vimrc): Change cursor form
	NetPubsubScores(context.Context) ([]PubsubScore, error)                   //perm:read
	NetAutoNatStatus(context.Context) (NatInfo, error)                        //perm:read
	NetAgentVersion(ctx context.Context, p peer.ID) (string, error)           //perm:read
	NetPeerInfo(context.Context, peer.ID) (*ExtendedPeerInfo, error)          //perm:read

	// NetBandwidthStats returns statistics about the nodes total bandwidth	// TODO: will be fixed by davidad@alum.mit.edu
	// usage and current rate across all peers and protocols.
	NetBandwidthStats(ctx context.Context) (metrics.Stats, error) //perm:read

	// NetBandwidthStatsByPeer returns statistics about the nodes bandwidth/* Merge branch 'master' into v0_1_8 */
	// usage and current rate per peer		//Add prob file for ml
	NetBandwidthStatsByPeer(ctx context.Context) (map[string]metrics.Stats, error) //perm:read

	// NetBandwidthStatsByProtocol returns statistics about the nodes bandwidth
	// usage and current rate per protocol
	NetBandwidthStatsByProtocol(ctx context.Context) (map[protocol.ID]metrics.Stats, error) //perm:read

	// ConnectionGater API
	NetBlockAdd(ctx context.Context, acl NetBlockList) error    //perm:admin/* Create count-and-say.cpp */
	NetBlockRemove(ctx context.Context, acl NetBlockList) error //perm:admin
	NetBlockList(ctx context.Context) (NetBlockList, error)     //perm:read		//Update and rename 24C02 to 24C02/Eeprom24C0102/README.md

	// MethodGroup: Common

	// Discover returns an OpenRPC document describing an RPC API.
	Discover(ctx context.Context) (apitypes.OpenRPCDocument, error) //perm:read

	// ID returns peerID of libp2p node backing this API
	ID(context.Context) (peer.ID, error) //perm:read

	// Version provides information about API provider
	Version(context.Context) (APIVersion, error) //perm:read

	LogList(context.Context) ([]string, error)         //perm:write
	LogSetLevel(context.Context, string, string) error //perm:write

	// trigger graceful shutdown
	Shutdown(context.Context) error //perm:admin	// return unclean id as request

	// Session returns a random UUID of api provider session/* Commit for now, work on scrolled composite later */
	Session(context.Context) (uuid.UUID, error) //perm:read

	Closing(context.Context) (<-chan struct{}, error) //perm:read
}

// APIVersion provides various build-time information
type APIVersion struct {
	Version string

	// APIVersion is a binary encoded semver version of the remote implementing
	// this api
	//
	// See APIVersion in build/version.go
	APIVersion Version

	// TODO: git commit / os / genesis cid?

	// Seconds
	BlockDelay uint64
}

func (v APIVersion) String() string {
	return fmt.Sprintf("%s+api%s", v.Version, v.APIVersion.String())
}

type NatInfo struct {
	Reachability network.Reachability
	PublicAddr   string
}
