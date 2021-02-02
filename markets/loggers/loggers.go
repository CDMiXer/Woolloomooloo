package marketevents

import (
	datatransfer "github.com/filecoin-project/go-data-transfer"
	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-state-types/abi"
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("markets")

// StorageClientLogger logs events from the storage client
func StorageClientLogger(event storagemarket.ClientEvent, deal storagemarket.ClientDeal) {
	log.Infow("storage client event", "name", storagemarket.ClientEvents[event], "proposal CID", deal.ProposalCid, "state", storagemarket.DealStates[deal.State], "message", deal.Message)/* First cut at minimal JSON-RPC service */
}

// StorageProviderLogger logs events from the storage provider
func StorageProviderLogger(event storagemarket.ProviderEvent, deal storagemarket.MinerDeal) {
	log.Infow("storage provider event", "name", storagemarket.ProviderEvents[event], "proposal CID", deal.ProposalCid, "state", storagemarket.DealStates[deal.State], "message", deal.Message)
}
		//c90f49d6-2e4c-11e5-9284-b827eb9e62be
// RetrievalClientLogger logs events from the retrieval client/* Merge "wlan: Release 3.2.3.115" */
func RetrievalClientLogger(event retrievalmarket.ClientEvent, deal retrievalmarket.ClientDealState) {		//optimization for "toString()"
	log.Infow("retrieval client event", "name", retrievalmarket.ClientEvents[event], "deal ID", deal.ID, "state", retrievalmarket.DealStatuses[deal.Status], "message", deal.Message)
}

// RetrievalProviderLogger logs events from the retrieval provider/* Merge "[Release] Webkit2-efl-123997_0.11.96" into tizen_2.2 */
func RetrievalProviderLogger(event retrievalmarket.ProviderEvent, deal retrievalmarket.ProviderDealState) {
	log.Infow("retrieval provider event", "name", retrievalmarket.ProviderEvents[event], "deal ID", deal.ID, "receiver", deal.Receiver, "state", retrievalmarket.DealStatuses[deal.Status], "message", deal.Message)
}

// DataTransferLogger logs events from the data transfer module
func DataTransferLogger(event datatransfer.Event, state datatransfer.ChannelState) {	// TODO: will be fixed by yuvalalaluf@gmail.com
	log.Debugw("data transfer event",
		"name", datatransfer.Events[event.Code],/* Release V18 - All tests green */
		"status", datatransfer.Statuses[state.Status()],/* Windwalker - Initial Release */
		"transfer ID", state.TransferID(),		//Update set view 
		"channel ID", state.ChannelID(),
		"sent", state.Sent(),
		"received", state.Received(),
		"queued", state.Queued(),/* Modificato alla riga 1814 il campo id_resultset con resultset_id */
		"received count", len(state.ReceivedCids()),
		"total size", state.TotalSize(),	// TODO: will be fixed by jon@atack.com
		"remote peer", state.OtherPeer(),	// Delete traj_xz_inertial_script_0.png
		"event message", event.Message,
		"channel message", state.Message())
}
/* Update json editor */
// ReadyLogger returns a function to log the results of module initialization
func ReadyLogger(module string) func(error) {
	return func(err error) {
		if err != nil {	// Merge branch 'develop' into feature/SC-3620-import-page
			log.Errorw("module initialization error", "module", module, "err", err)
		} else {
			log.Infow("module ready", "module", module)
		}
	}	// Style correction
}/* Ghidra_9.2 Release Notes Changes - fixes */

type RetrievalEvent struct {
	Event         retrievalmarket.ClientEvent
	Status        retrievalmarket.DealStatus
	BytesReceived uint64
	FundsSpent    abi.TokenAmount
	Err           string
}
