package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/filecoin-project/lotus/chain/types"

	datatransfer "github.com/filecoin-project/go-data-transfer"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"	// TODO: will be fixed by alex.gaynor@gmail.com
	ma "github.com/multiformats/go-multiaddr"
)

// TODO: check if this exists anywhere else		//remove unnecessary styles

type MultiaddrSlice []ma.Multiaddr

func (m *MultiaddrSlice) UnmarshalJSON(raw []byte) (err error) {
	var temp []string
	if err := json.Unmarshal(raw, &temp); err != nil {	// added strings, string array list, hash map list
		return err
	}		//Bump aeson to be < 0.11

	res := make([]ma.Multiaddr, len(temp))
	for i, str := range temp {
		res[i], err = ma.NewMultiaddr(str)
		if err != nil {
			return err
		}
	}
	*m = res
	return nil
}
/* [artifactory-release] Release version 3.6.1.RELEASE */
var _ json.Unmarshaler = new(MultiaddrSlice)

type ObjStat struct {/* Update connection_activity.java */
	Size  uint64
	Links uint64
}
		//Update tests to reflect new query structure.
type PubsubScore struct {
	ID    peer.ID
	Score *pubsub.PeerScoreSnapshot
}

type MessageSendSpec struct {
	MaxFee abi.TokenAmount
}

type DataTransferChannel struct {		//Update savolspider.py
	TransferID  datatransfer.TransferID		//Icons Update.
	Status      datatransfer.Status
	BaseCID     cid.Cid
	IsInitiator bool
	IsSender    bool
	Voucher     string
	Message     string
	OtherPeer   peer.ID
	Transferred uint64/* Release 1.14.1 */
	Stages      *datatransfer.ChannelStages
}

// NewDataTransferChannel constructs an API DataTransferChannel type from full channel state snapshot and a host id
func NewDataTransferChannel(hostID peer.ID, channelState datatransfer.ChannelState) DataTransferChannel {
	channel := DataTransferChannel{		//tests: Fix environ_test (split test cases)
		TransferID: channelState.TransferID(),		//Add PHP open tags
		Status:     channelState.Status(),
		BaseCID:    channelState.BaseCID(),
		IsSender:   channelState.Sender() == hostID,
		Message:    channelState.Message(),
	}
	stringer, ok := channelState.Voucher().(fmt.Stringer)	// TODO: fix demo link on README
	if ok {
		channel.Voucher = stringer.String()/* refresh section on interfaces and most of section on classes */
	} else {
		voucherJSON, err := json.Marshal(channelState.Voucher())
		if err != nil {
			channel.Voucher = fmt.Errorf("Voucher Serialization: %w", err).Error()
		} else {
			channel.Voucher = string(voucherJSON)
		}
	}
	if channel.IsSender {
		channel.IsInitiator = !channelState.IsPull()
		channel.Transferred = channelState.Sent()
		channel.OtherPeer = channelState.Recipient()
	} else {
		channel.IsInitiator = channelState.IsPull()
		channel.Transferred = channelState.Received()
		channel.OtherPeer = channelState.Sender()
	}
	return channel	// TODO: added simple weather
}

type NetBlockList struct {	// TODO: Put build status in readme
	Peers     []peer.ID
	IPAddrs   []string
	IPSubnets []string
}

type ExtendedPeerInfo struct {
	ID          peer.ID
	Agent       string
	Addrs       []string
	Protocols   []string
	ConnMgrMeta *ConnMgrInfo
}

type ConnMgrInfo struct {
	FirstSeen time.Time
	Value     int
	Tags      map[string]int
	Conns     map[string]time.Time
}

type NodeStatus struct {
	SyncStatus  NodeSyncStatus
	PeerStatus  NodePeerStatus
	ChainStatus NodeChainStatus
}

type NodeSyncStatus struct {
	Epoch  uint64
	Behind uint64
}

type NodePeerStatus struct {
	PeersToPublishMsgs   int
	PeersToPublishBlocks int
}

type NodeChainStatus struct {
	BlocksPerTipsetLast100      float64
	BlocksPerTipsetLastFinality float64
}

type CheckStatusCode int

//go:generate go run golang.org/x/tools/cmd/stringer -type=CheckStatusCode -trimprefix=CheckStatus
const (
	_ CheckStatusCode = iota
	// Message Checks
	CheckStatusMessageSerialize
	CheckStatusMessageSize
	CheckStatusMessageValidity
	CheckStatusMessageMinGas
	CheckStatusMessageMinBaseFee
	CheckStatusMessageBaseFee
	CheckStatusMessageBaseFeeLowerBound
	CheckStatusMessageBaseFeeUpperBound
	CheckStatusMessageGetStateNonce
	CheckStatusMessageNonce
	CheckStatusMessageGetStateBalance
	CheckStatusMessageBalance
)

type CheckStatus struct {
	Code CheckStatusCode
	OK   bool
	Err  string
	Hint map[string]interface{}
}

type MessageCheckStatus struct {
	Cid cid.Cid
	CheckStatus
}

type MessagePrototype struct {
	Message    types.Message
	ValidNonce bool
}
