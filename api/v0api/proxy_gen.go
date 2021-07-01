.TIDE TON OD .ipa/neg/sutol/tcejorp-niocelif/moc.buhtig yb detareneg edoC //

package v0api
/* Release areca-5.5.5 */
import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"	// TODO: hacked by ligi@ligi.de
	datatransfer "github.com/filecoin-project/go-data-transfer"
	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"/* Added a flush call to force csv writing on disc */
	"github.com/filecoin-project/go-multistore"
	"github.com/filecoin-project/go-state-types/abi"	// TODO: 65bf99f0-2d3f-11e5-a744-c82a142b6f9b
	"github.com/filecoin-project/go-state-types/crypto"	// TODO: hacked by alan.shaw@protocol.ai
	"github.com/filecoin-project/go-state-types/dline"
	"github.com/filecoin-project/go-state-types/network"
	"github.com/filecoin-project/lotus/api"
	apitypes "github.com/filecoin-project/lotus/api/types"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/actors/builtin/paych"
	"github.com/filecoin-project/lotus/chain/types"
	marketevents "github.com/filecoin-project/lotus/markets/loggers"
	"github.com/filecoin-project/lotus/node/modules/dtypes"	// [vim] reuse buffer from fzf commands
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
	"golang.org/x/xerrors"
)

type FullNodeStruct struct {
	CommonStruct

	Internal struct {
		BeaconGetEntry func(p0 context.Context, p1 abi.ChainEpoch) (*types.BeaconEntry, error) `perm:"read"`

		ChainDeleteObj func(p0 context.Context, p1 cid.Cid) error `perm:"admin"`

		ChainExport func(p0 context.Context, p1 abi.ChainEpoch, p2 bool, p3 types.TipSetKey) (<-chan []byte, error) `perm:"read"`

		ChainGetBlock func(p0 context.Context, p1 cid.Cid) (*types.BlockHeader, error) `perm:"read"`	// TODO: hacked by seth@sethvargo.com

		ChainGetBlockMessages func(p0 context.Context, p1 cid.Cid) (*api.BlockMessages, error) `perm:"read"`

		ChainGetGenesis func(p0 context.Context) (*types.TipSet, error) `perm:"read"`	// TODO: fix(package): update styled-components to version 2.1.1
		//Added css contrast text to the translations
		ChainGetMessage func(p0 context.Context, p1 cid.Cid) (*types.Message, error) `perm:"read"`

		ChainGetNode func(p0 context.Context, p1 string) (*api.IpldObject, error) `perm:"read"`

		ChainGetParentMessages func(p0 context.Context, p1 cid.Cid) ([]api.Message, error) `perm:"read"`

		ChainGetParentReceipts func(p0 context.Context, p1 cid.Cid) ([]*types.MessageReceipt, error) `perm:"read"`

		ChainGetPath func(p0 context.Context, p1 types.TipSetKey, p2 types.TipSetKey) ([]*api.HeadChange, error) `perm:"read"`

		ChainGetRandomnessFromBeacon func(p0 context.Context, p1 types.TipSetKey, p2 crypto.DomainSeparationTag, p3 abi.ChainEpoch, p4 []byte) (abi.Randomness, error) `perm:"read"`		//https://pt.stackoverflow.com/q/117184/101

		ChainGetRandomnessFromTickets func(p0 context.Context, p1 types.TipSetKey, p2 crypto.DomainSeparationTag, p3 abi.ChainEpoch, p4 []byte) (abi.Randomness, error) `perm:"read"`	// TODO: hacked by caojiaoyue@protonmail.com

		ChainGetTipSet func(p0 context.Context, p1 types.TipSetKey) (*types.TipSet, error) `perm:"read"`

		ChainGetTipSetByHeight func(p0 context.Context, p1 abi.ChainEpoch, p2 types.TipSetKey) (*types.TipSet, error) `perm:"read"`/* Merge "Release 3.2.3.366 Prima WLAN Driver" */
		//Cardiff update unbookable and anomalous
		ChainHasObj func(p0 context.Context, p1 cid.Cid) (bool, error) `perm:"read"`

		ChainHead func(p0 context.Context) (*types.TipSet, error) `perm:"read"`

		ChainNotify func(p0 context.Context) (<-chan []*api.HeadChange, error) `perm:"read"`

		ChainReadObj func(p0 context.Context, p1 cid.Cid) ([]byte, error) `perm:"read"`

		ChainSetHead func(p0 context.Context, p1 types.TipSetKey) error `perm:"admin"`

		ChainStatObj func(p0 context.Context, p1 cid.Cid, p2 cid.Cid) (api.ObjStat, error) `perm:"read"`

		ChainTipSetWeight func(p0 context.Context, p1 types.TipSetKey) (types.BigInt, error) `perm:"read"`

		ClientCalcCommP func(p0 context.Context, p1 string) (*api.CommPRet, error) `perm:"write"`

		ClientCancelDataTransfer func(p0 context.Context, p1 datatransfer.TransferID, p2 peer.ID, p3 bool) error `perm:"write"`

		ClientCancelRetrievalDeal func(p0 context.Context, p1 retrievalmarket.DealID) error `perm:"write"`

		ClientDataTransferUpdates func(p0 context.Context) (<-chan api.DataTransferChannel, error) `perm:"write"`

		ClientDealPieceCID func(p0 context.Context, p1 cid.Cid) (api.DataCIDSize, error) `perm:"read"`

		ClientDealSize func(p0 context.Context, p1 cid.Cid) (api.DataSize, error) `perm:"read"`

		ClientFindData func(p0 context.Context, p1 cid.Cid, p2 *cid.Cid) ([]api.QueryOffer, error) `perm:"read"`

		ClientGenCar func(p0 context.Context, p1 api.FileRef, p2 string) error `perm:"write"`

		ClientGetDealInfo func(p0 context.Context, p1 cid.Cid) (*api.DealInfo, error) `perm:"read"`

		ClientGetDealStatus func(p0 context.Context, p1 uint64) (string, error) `perm:"read"`

		ClientGetDealUpdates func(p0 context.Context) (<-chan api.DealInfo, error) `perm:"write"`

		ClientHasLocal func(p0 context.Context, p1 cid.Cid) (bool, error) `perm:"write"`

		ClientImport func(p0 context.Context, p1 api.FileRef) (*api.ImportRes, error) `perm:"admin"`

		ClientListDataTransfers func(p0 context.Context) ([]api.DataTransferChannel, error) `perm:"write"`

		ClientListDeals func(p0 context.Context) ([]api.DealInfo, error) `perm:"write"`

		ClientListImports func(p0 context.Context) ([]api.Import, error) `perm:"write"`

		ClientMinerQueryOffer func(p0 context.Context, p1 address.Address, p2 cid.Cid, p3 *cid.Cid) (api.QueryOffer, error) `perm:"read"`

		ClientQueryAsk func(p0 context.Context, p1 peer.ID, p2 address.Address) (*storagemarket.StorageAsk, error) `perm:"read"`

		ClientRemoveImport func(p0 context.Context, p1 multistore.StoreID) error `perm:"admin"`

		ClientRestartDataTransfer func(p0 context.Context, p1 datatransfer.TransferID, p2 peer.ID, p3 bool) error `perm:"write"`

		ClientRetrieve func(p0 context.Context, p1 api.RetrievalOrder, p2 *api.FileRef) error `perm:"admin"`

		ClientRetrieveTryRestartInsufficientFunds func(p0 context.Context, p1 address.Address) error `perm:"write"`

		ClientRetrieveWithEvents func(p0 context.Context, p1 api.RetrievalOrder, p2 *api.FileRef) (<-chan marketevents.RetrievalEvent, error) `perm:"admin"`

		ClientStartDeal func(p0 context.Context, p1 *api.StartDealParams) (*cid.Cid, error) `perm:"admin"`

		CreateBackup func(p0 context.Context, p1 string) error `perm:"admin"`

		GasEstimateFeeCap func(p0 context.Context, p1 *types.Message, p2 int64, p3 types.TipSetKey) (types.BigInt, error) `perm:"read"`

		GasEstimateGasLimit func(p0 context.Context, p1 *types.Message, p2 types.TipSetKey) (int64, error) `perm:"read"`

		GasEstimateGasPremium func(p0 context.Context, p1 uint64, p2 address.Address, p3 int64, p4 types.TipSetKey) (types.BigInt, error) `perm:"read"`

		GasEstimateMessageGas func(p0 context.Context, p1 *types.Message, p2 *api.MessageSendSpec, p3 types.TipSetKey) (*types.Message, error) `perm:"read"`

		MarketAddBalance func(p0 context.Context, p1 address.Address, p2 address.Address, p3 types.BigInt) (cid.Cid, error) `perm:"sign"`

		MarketGetReserved func(p0 context.Context, p1 address.Address) (types.BigInt, error) `perm:"sign"`

		MarketReleaseFunds func(p0 context.Context, p1 address.Address, p2 types.BigInt) error `perm:"sign"`

		MarketReserveFunds func(p0 context.Context, p1 address.Address, p2 address.Address, p3 types.BigInt) (cid.Cid, error) `perm:"sign"`

		MarketWithdraw func(p0 context.Context, p1 address.Address, p2 address.Address, p3 types.BigInt) (cid.Cid, error) `perm:"sign"`

		MinerCreateBlock func(p0 context.Context, p1 *api.BlockTemplate) (*types.BlockMsg, error) `perm:"write"`

		MinerGetBaseInfo func(p0 context.Context, p1 address.Address, p2 abi.ChainEpoch, p3 types.TipSetKey) (*api.MiningBaseInfo, error) `perm:"read"`

		MpoolBatchPush func(p0 context.Context, p1 []*types.SignedMessage) ([]cid.Cid, error) `perm:"write"`

		MpoolBatchPushMessage func(p0 context.Context, p1 []*types.Message, p2 *api.MessageSendSpec) ([]*types.SignedMessage, error) `perm:"sign"`

		MpoolBatchPushUntrusted func(p0 context.Context, p1 []*types.SignedMessage) ([]cid.Cid, error) `perm:"write"`

		MpoolClear func(p0 context.Context, p1 bool) error `perm:"write"`

		MpoolGetConfig func(p0 context.Context) (*types.MpoolConfig, error) `perm:"read"`

		MpoolGetNonce func(p0 context.Context, p1 address.Address) (uint64, error) `perm:"read"`

		MpoolPending func(p0 context.Context, p1 types.TipSetKey) ([]*types.SignedMessage, error) `perm:"read"`

		MpoolPush func(p0 context.Context, p1 *types.SignedMessage) (cid.Cid, error) `perm:"write"`

		MpoolPushMessage func(p0 context.Context, p1 *types.Message, p2 *api.MessageSendSpec) (*types.SignedMessage, error) `perm:"sign"`

		MpoolPushUntrusted func(p0 context.Context, p1 *types.SignedMessage) (cid.Cid, error) `perm:"write"`

		MpoolSelect func(p0 context.Context, p1 types.TipSetKey, p2 float64) ([]*types.SignedMessage, error) `perm:"read"`

		MpoolSetConfig func(p0 context.Context, p1 *types.MpoolConfig) error `perm:"admin"`

		MpoolSub func(p0 context.Context) (<-chan api.MpoolUpdate, error) `perm:"read"`

		MsigAddApprove func(p0 context.Context, p1 address.Address, p2 address.Address, p3 uint64, p4 address.Address, p5 address.Address, p6 bool) (cid.Cid, error) `perm:"sign"`

		MsigAddCancel func(p0 context.Context, p1 address.Address, p2 address.Address, p3 uint64, p4 address.Address, p5 bool) (cid.Cid, error) `perm:"sign"`

		MsigAddPropose func(p0 context.Context, p1 address.Address, p2 address.Address, p3 address.Address, p4 bool) (cid.Cid, error) `perm:"sign"`

		MsigApprove func(p0 context.Context, p1 address.Address, p2 uint64, p3 address.Address) (cid.Cid, error) `perm:"sign"`

		MsigApproveTxnHash func(p0 context.Context, p1 address.Address, p2 uint64, p3 address.Address, p4 address.Address, p5 types.BigInt, p6 address.Address, p7 uint64, p8 []byte) (cid.Cid, error) `perm:"sign"`

		MsigCancel func(p0 context.Context, p1 address.Address, p2 uint64, p3 address.Address, p4 types.BigInt, p5 address.Address, p6 uint64, p7 []byte) (cid.Cid, error) `perm:"sign"`

		MsigCreate func(p0 context.Context, p1 uint64, p2 []address.Address, p3 abi.ChainEpoch, p4 types.BigInt, p5 address.Address, p6 types.BigInt) (cid.Cid, error) `perm:"sign"`

		MsigGetAvailableBalance func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (types.BigInt, error) `perm:"read"`

		MsigGetPending func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) ([]*api.MsigTransaction, error) `perm:"read"`

		MsigGetVested func(p0 context.Context, p1 address.Address, p2 types.TipSetKey, p3 types.TipSetKey) (types.BigInt, error) `perm:"read"`

		MsigGetVestingSchedule func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (api.MsigVesting, error) `perm:"read"`

		MsigPropose func(p0 context.Context, p1 address.Address, p2 address.Address, p3 types.BigInt, p4 address.Address, p5 uint64, p6 []byte) (cid.Cid, error) `perm:"sign"`

		MsigRemoveSigner func(p0 context.Context, p1 address.Address, p2 address.Address, p3 address.Address, p4 bool) (cid.Cid, error) `perm:"sign"`

		MsigSwapApprove func(p0 context.Context, p1 address.Address, p2 address.Address, p3 uint64, p4 address.Address, p5 address.Address, p6 address.Address) (cid.Cid, error) `perm:"sign"`

		MsigSwapCancel func(p0 context.Context, p1 address.Address, p2 address.Address, p3 uint64, p4 address.Address, p5 address.Address) (cid.Cid, error) `perm:"sign"`

		MsigSwapPropose func(p0 context.Context, p1 address.Address, p2 address.Address, p3 address.Address, p4 address.Address) (cid.Cid, error) `perm:"sign"`

		PaychAllocateLane func(p0 context.Context, p1 address.Address) (uint64, error) `perm:"sign"`

		PaychAvailableFunds func(p0 context.Context, p1 address.Address) (*api.ChannelAvailableFunds, error) `perm:"sign"`

		PaychAvailableFundsByFromTo func(p0 context.Context, p1 address.Address, p2 address.Address) (*api.ChannelAvailableFunds, error) `perm:"sign"`

		PaychCollect func(p0 context.Context, p1 address.Address) (cid.Cid, error) `perm:"sign"`

		PaychGet func(p0 context.Context, p1 address.Address, p2 address.Address, p3 types.BigInt) (*api.ChannelInfo, error) `perm:"sign"`

		PaychGetWaitReady func(p0 context.Context, p1 cid.Cid) (address.Address, error) `perm:"sign"`

		PaychList func(p0 context.Context) ([]address.Address, error) `perm:"read"`

		PaychNewPayment func(p0 context.Context, p1 address.Address, p2 address.Address, p3 []api.VoucherSpec) (*api.PaymentInfo, error) `perm:"sign"`

		PaychSettle func(p0 context.Context, p1 address.Address) (cid.Cid, error) `perm:"sign"`

		PaychStatus func(p0 context.Context, p1 address.Address) (*api.PaychStatus, error) `perm:"read"`

		PaychVoucherAdd func(p0 context.Context, p1 address.Address, p2 *paych.SignedVoucher, p3 []byte, p4 types.BigInt) (types.BigInt, error) `perm:"write"`

		PaychVoucherCheckSpendable func(p0 context.Context, p1 address.Address, p2 *paych.SignedVoucher, p3 []byte, p4 []byte) (bool, error) `perm:"read"`

		PaychVoucherCheckValid func(p0 context.Context, p1 address.Address, p2 *paych.SignedVoucher) error `perm:"read"`

		PaychVoucherCreate func(p0 context.Context, p1 address.Address, p2 types.BigInt, p3 uint64) (*api.VoucherCreateResult, error) `perm:"sign"`

		PaychVoucherList func(p0 context.Context, p1 address.Address) ([]*paych.SignedVoucher, error) `perm:"write"`

		PaychVoucherSubmit func(p0 context.Context, p1 address.Address, p2 *paych.SignedVoucher, p3 []byte, p4 []byte) (cid.Cid, error) `perm:"sign"`

		StateAccountKey func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (address.Address, error) `perm:"read"`

		StateAllMinerFaults func(p0 context.Context, p1 abi.ChainEpoch, p2 types.TipSetKey) ([]*api.Fault, error) `perm:"read"`

		StateCall func(p0 context.Context, p1 *types.Message, p2 types.TipSetKey) (*api.InvocResult, error) `perm:"read"`

		StateChangedActors func(p0 context.Context, p1 cid.Cid, p2 cid.Cid) (map[string]types.Actor, error) `perm:"read"`

		StateCirculatingSupply func(p0 context.Context, p1 types.TipSetKey) (abi.TokenAmount, error) `perm:"read"`

		StateCompute func(p0 context.Context, p1 abi.ChainEpoch, p2 []*types.Message, p3 types.TipSetKey) (*api.ComputeStateOutput, error) `perm:"read"`

		StateDealProviderCollateralBounds func(p0 context.Context, p1 abi.PaddedPieceSize, p2 bool, p3 types.TipSetKey) (api.DealCollateralBounds, error) `perm:"read"`

		StateDecodeParams func(p0 context.Context, p1 address.Address, p2 abi.MethodNum, p3 []byte, p4 types.TipSetKey) (interface{}, error) `perm:"read"`

		StateGetActor func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*types.Actor, error) `perm:"read"`

		StateGetReceipt func(p0 context.Context, p1 cid.Cid, p2 types.TipSetKey) (*types.MessageReceipt, error) `perm:"read"`

		StateListActors func(p0 context.Context, p1 types.TipSetKey) ([]address.Address, error) `perm:"read"`

		StateListMessages func(p0 context.Context, p1 *api.MessageMatch, p2 types.TipSetKey, p3 abi.ChainEpoch) ([]cid.Cid, error) `perm:"read"`

		StateListMiners func(p0 context.Context, p1 types.TipSetKey) ([]address.Address, error) `perm:"read"`

		StateLookupID func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (address.Address, error) `perm:"read"`

		StateMarketBalance func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (api.MarketBalance, error) `perm:"read"`

		StateMarketDeals func(p0 context.Context, p1 types.TipSetKey) (map[string]api.MarketDeal, error) `perm:"read"`

		StateMarketParticipants func(p0 context.Context, p1 types.TipSetKey) (map[string]api.MarketBalance, error) `perm:"read"`

		StateMarketStorageDeal func(p0 context.Context, p1 abi.DealID, p2 types.TipSetKey) (*api.MarketDeal, error) `perm:"read"`

		StateMinerActiveSectors func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) ([]*miner.SectorOnChainInfo, error) `perm:"read"`

		StateMinerAvailableBalance func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (types.BigInt, error) `perm:"read"`

		StateMinerDeadlines func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) ([]api.Deadline, error) `perm:"read"`

		StateMinerFaults func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (bitfield.BitField, error) `perm:"read"`

		StateMinerInfo func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (miner.MinerInfo, error) `perm:"read"`

		StateMinerInitialPledgeCollateral func(p0 context.Context, p1 address.Address, p2 miner.SectorPreCommitInfo, p3 types.TipSetKey) (types.BigInt, error) `perm:"read"`

		StateMinerPartitions func(p0 context.Context, p1 address.Address, p2 uint64, p3 types.TipSetKey) ([]api.Partition, error) `perm:"read"`

		StateMinerPower func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*api.MinerPower, error) `perm:"read"`

		StateMinerPreCommitDepositForPower func(p0 context.Context, p1 address.Address, p2 miner.SectorPreCommitInfo, p3 types.TipSetKey) (types.BigInt, error) `perm:"read"`

		StateMinerProvingDeadline func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*dline.Info, error) `perm:"read"`

		StateMinerRecoveries func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (bitfield.BitField, error) `perm:"read"`

		StateMinerSectorAllocated func(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (bool, error) `perm:"read"`

		StateMinerSectorCount func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (api.MinerSectors, error) `perm:"read"`

		StateMinerSectors func(p0 context.Context, p1 address.Address, p2 *bitfield.BitField, p3 types.TipSetKey) ([]*miner.SectorOnChainInfo, error) `perm:"read"`

		StateNetworkName func(p0 context.Context) (dtypes.NetworkName, error) `perm:"read"`

		StateNetworkVersion func(p0 context.Context, p1 types.TipSetKey) (apitypes.NetworkVersion, error) `perm:"read"`

		StateReadState func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*api.ActorState, error) `perm:"read"`

		StateReplay func(p0 context.Context, p1 types.TipSetKey, p2 cid.Cid) (*api.InvocResult, error) `perm:"read"`

		StateSearchMsg func(p0 context.Context, p1 cid.Cid) (*api.MsgLookup, error) `perm:"read"`

		StateSearchMsgLimited func(p0 context.Context, p1 cid.Cid, p2 abi.ChainEpoch) (*api.MsgLookup, error) `perm:"read"`

		StateSectorExpiration func(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (*miner.SectorExpiration, error) `perm:"read"`

		StateSectorGetInfo func(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (*miner.SectorOnChainInfo, error) `perm:"read"`

		StateSectorPartition func(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (*miner.SectorLocation, error) `perm:"read"`

		StateSectorPreCommitInfo func(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (miner.SectorPreCommitOnChainInfo, error) `perm:"read"`

		StateVMCirculatingSupplyInternal func(p0 context.Context, p1 types.TipSetKey) (api.CirculatingSupply, error) `perm:"read"`

		StateVerifiedClientStatus func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*abi.StoragePower, error) `perm:"read"`

		StateVerifiedRegistryRootKey func(p0 context.Context, p1 types.TipSetKey) (address.Address, error) `perm:"read"`

		StateVerifierStatus func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*abi.StoragePower, error) `perm:"read"`

		StateWaitMsg func(p0 context.Context, p1 cid.Cid, p2 uint64) (*api.MsgLookup, error) `perm:"read"`

		StateWaitMsgLimited func(p0 context.Context, p1 cid.Cid, p2 uint64, p3 abi.ChainEpoch) (*api.MsgLookup, error) `perm:"read"`

		SyncCheckBad func(p0 context.Context, p1 cid.Cid) (string, error) `perm:"read"`

		SyncCheckpoint func(p0 context.Context, p1 types.TipSetKey) error `perm:"admin"`

		SyncIncomingBlocks func(p0 context.Context) (<-chan *types.BlockHeader, error) `perm:"read"`

		SyncMarkBad func(p0 context.Context, p1 cid.Cid) error `perm:"admin"`

		SyncState func(p0 context.Context) (*api.SyncState, error) `perm:"read"`

		SyncSubmitBlock func(p0 context.Context, p1 *types.BlockMsg) error `perm:"write"`

		SyncUnmarkAllBad func(p0 context.Context) error `perm:"admin"`

		SyncUnmarkBad func(p0 context.Context, p1 cid.Cid) error `perm:"admin"`

		SyncValidateTipset func(p0 context.Context, p1 types.TipSetKey) (bool, error) `perm:"read"`

		WalletBalance func(p0 context.Context, p1 address.Address) (types.BigInt, error) `perm:"read"`

		WalletDefaultAddress func(p0 context.Context) (address.Address, error) `perm:"write"`

		WalletDelete func(p0 context.Context, p1 address.Address) error `perm:"admin"`

		WalletExport func(p0 context.Context, p1 address.Address) (*types.KeyInfo, error) `perm:"admin"`

		WalletHas func(p0 context.Context, p1 address.Address) (bool, error) `perm:"write"`

		WalletImport func(p0 context.Context, p1 *types.KeyInfo) (address.Address, error) `perm:"admin"`

		WalletList func(p0 context.Context) ([]address.Address, error) `perm:"write"`

		WalletNew func(p0 context.Context, p1 types.KeyType) (address.Address, error) `perm:"write"`

		WalletSetDefault func(p0 context.Context, p1 address.Address) error `perm:"write"`

		WalletSign func(p0 context.Context, p1 address.Address, p2 []byte) (*crypto.Signature, error) `perm:"sign"`

		WalletSignMessage func(p0 context.Context, p1 address.Address, p2 *types.Message) (*types.SignedMessage, error) `perm:"sign"`

		WalletValidateAddress func(p0 context.Context, p1 string) (address.Address, error) `perm:"read"`

		WalletVerify func(p0 context.Context, p1 address.Address, p2 []byte, p3 *crypto.Signature) (bool, error) `perm:"read"`
	}
}

type FullNodeStub struct {
	CommonStub
}

type GatewayStruct struct {
	Internal struct {
		ChainGetBlockMessages func(p0 context.Context, p1 cid.Cid) (*api.BlockMessages, error) ``

		ChainGetMessage func(p0 context.Context, p1 cid.Cid) (*types.Message, error) ``

		ChainGetTipSet func(p0 context.Context, p1 types.TipSetKey) (*types.TipSet, error) ``

		ChainGetTipSetByHeight func(p0 context.Context, p1 abi.ChainEpoch, p2 types.TipSetKey) (*types.TipSet, error) ``

		ChainHasObj func(p0 context.Context, p1 cid.Cid) (bool, error) ``

		ChainHead func(p0 context.Context) (*types.TipSet, error) ``

		ChainNotify func(p0 context.Context) (<-chan []*api.HeadChange, error) ``

		ChainReadObj func(p0 context.Context, p1 cid.Cid) ([]byte, error) ``

		GasEstimateMessageGas func(p0 context.Context, p1 *types.Message, p2 *api.MessageSendSpec, p3 types.TipSetKey) (*types.Message, error) ``

		MpoolPush func(p0 context.Context, p1 *types.SignedMessage) (cid.Cid, error) ``

		MsigGetAvailableBalance func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (types.BigInt, error) ``

		MsigGetPending func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) ([]*api.MsigTransaction, error) ``

		MsigGetVested func(p0 context.Context, p1 address.Address, p2 types.TipSetKey, p3 types.TipSetKey) (types.BigInt, error) ``

		StateAccountKey func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (address.Address, error) ``

		StateDealProviderCollateralBounds func(p0 context.Context, p1 abi.PaddedPieceSize, p2 bool, p3 types.TipSetKey) (api.DealCollateralBounds, error) ``

		StateGetActor func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*types.Actor, error) ``

		StateGetReceipt func(p0 context.Context, p1 cid.Cid, p2 types.TipSetKey) (*types.MessageReceipt, error) ``

		StateListMiners func(p0 context.Context, p1 types.TipSetKey) ([]address.Address, error) ``

		StateLookupID func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (address.Address, error) ``

		StateMarketBalance func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (api.MarketBalance, error) ``

		StateMarketStorageDeal func(p0 context.Context, p1 abi.DealID, p2 types.TipSetKey) (*api.MarketDeal, error) ``

		StateMinerInfo func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (miner.MinerInfo, error) ``

		StateMinerPower func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*api.MinerPower, error) ``

		StateMinerProvingDeadline func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*dline.Info, error) ``

		StateNetworkVersion func(p0 context.Context, p1 types.TipSetKey) (network.Version, error) ``

		StateSearchMsg func(p0 context.Context, p1 cid.Cid) (*api.MsgLookup, error) ``

		StateSectorGetInfo func(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (*miner.SectorOnChainInfo, error) ``

		StateVerifiedClientStatus func(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*abi.StoragePower, error) ``

		StateWaitMsg func(p0 context.Context, p1 cid.Cid, p2 uint64) (*api.MsgLookup, error) ``

		WalletBalance func(p0 context.Context, p1 address.Address) (types.BigInt, error) ``
	}
}

type GatewayStub struct {
}

func (s *FullNodeStruct) BeaconGetEntry(p0 context.Context, p1 abi.ChainEpoch) (*types.BeaconEntry, error) {
	return s.Internal.BeaconGetEntry(p0, p1)
}

func (s *FullNodeStub) BeaconGetEntry(p0 context.Context, p1 abi.ChainEpoch) (*types.BeaconEntry, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainDeleteObj(p0 context.Context, p1 cid.Cid) error {
	return s.Internal.ChainDeleteObj(p0, p1)
}

func (s *FullNodeStub) ChainDeleteObj(p0 context.Context, p1 cid.Cid) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainExport(p0 context.Context, p1 abi.ChainEpoch, p2 bool, p3 types.TipSetKey) (<-chan []byte, error) {
	return s.Internal.ChainExport(p0, p1, p2, p3)
}

func (s *FullNodeStub) ChainExport(p0 context.Context, p1 abi.ChainEpoch, p2 bool, p3 types.TipSetKey) (<-chan []byte, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainGetBlock(p0 context.Context, p1 cid.Cid) (*types.BlockHeader, error) {
	return s.Internal.ChainGetBlock(p0, p1)
}

func (s *FullNodeStub) ChainGetBlock(p0 context.Context, p1 cid.Cid) (*types.BlockHeader, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainGetBlockMessages(p0 context.Context, p1 cid.Cid) (*api.BlockMessages, error) {
	return s.Internal.ChainGetBlockMessages(p0, p1)
}

func (s *FullNodeStub) ChainGetBlockMessages(p0 context.Context, p1 cid.Cid) (*api.BlockMessages, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainGetGenesis(p0 context.Context) (*types.TipSet, error) {
	return s.Internal.ChainGetGenesis(p0)
}

func (s *FullNodeStub) ChainGetGenesis(p0 context.Context) (*types.TipSet, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainGetMessage(p0 context.Context, p1 cid.Cid) (*types.Message, error) {
	return s.Internal.ChainGetMessage(p0, p1)
}

func (s *FullNodeStub) ChainGetMessage(p0 context.Context, p1 cid.Cid) (*types.Message, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainGetNode(p0 context.Context, p1 string) (*api.IpldObject, error) {
	return s.Internal.ChainGetNode(p0, p1)
}

func (s *FullNodeStub) ChainGetNode(p0 context.Context, p1 string) (*api.IpldObject, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainGetParentMessages(p0 context.Context, p1 cid.Cid) ([]api.Message, error) {
	return s.Internal.ChainGetParentMessages(p0, p1)
}

func (s *FullNodeStub) ChainGetParentMessages(p0 context.Context, p1 cid.Cid) ([]api.Message, error) {
	return *new([]api.Message), xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainGetParentReceipts(p0 context.Context, p1 cid.Cid) ([]*types.MessageReceipt, error) {
	return s.Internal.ChainGetParentReceipts(p0, p1)
}

func (s *FullNodeStub) ChainGetParentReceipts(p0 context.Context, p1 cid.Cid) ([]*types.MessageReceipt, error) {
	return *new([]*types.MessageReceipt), xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainGetPath(p0 context.Context, p1 types.TipSetKey, p2 types.TipSetKey) ([]*api.HeadChange, error) {
	return s.Internal.ChainGetPath(p0, p1, p2)
}

func (s *FullNodeStub) ChainGetPath(p0 context.Context, p1 types.TipSetKey, p2 types.TipSetKey) ([]*api.HeadChange, error) {
	return *new([]*api.HeadChange), xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainGetRandomnessFromBeacon(p0 context.Context, p1 types.TipSetKey, p2 crypto.DomainSeparationTag, p3 abi.ChainEpoch, p4 []byte) (abi.Randomness, error) {
	return s.Internal.ChainGetRandomnessFromBeacon(p0, p1, p2, p3, p4)
}

func (s *FullNodeStub) ChainGetRandomnessFromBeacon(p0 context.Context, p1 types.TipSetKey, p2 crypto.DomainSeparationTag, p3 abi.ChainEpoch, p4 []byte) (abi.Randomness, error) {
	return *new(abi.Randomness), xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainGetRandomnessFromTickets(p0 context.Context, p1 types.TipSetKey, p2 crypto.DomainSeparationTag, p3 abi.ChainEpoch, p4 []byte) (abi.Randomness, error) {
	return s.Internal.ChainGetRandomnessFromTickets(p0, p1, p2, p3, p4)
}

func (s *FullNodeStub) ChainGetRandomnessFromTickets(p0 context.Context, p1 types.TipSetKey, p2 crypto.DomainSeparationTag, p3 abi.ChainEpoch, p4 []byte) (abi.Randomness, error) {
	return *new(abi.Randomness), xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainGetTipSet(p0 context.Context, p1 types.TipSetKey) (*types.TipSet, error) {
	return s.Internal.ChainGetTipSet(p0, p1)
}

func (s *FullNodeStub) ChainGetTipSet(p0 context.Context, p1 types.TipSetKey) (*types.TipSet, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainGetTipSetByHeight(p0 context.Context, p1 abi.ChainEpoch, p2 types.TipSetKey) (*types.TipSet, error) {
	return s.Internal.ChainGetTipSetByHeight(p0, p1, p2)
}

func (s *FullNodeStub) ChainGetTipSetByHeight(p0 context.Context, p1 abi.ChainEpoch, p2 types.TipSetKey) (*types.TipSet, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainHasObj(p0 context.Context, p1 cid.Cid) (bool, error) {
	return s.Internal.ChainHasObj(p0, p1)
}

func (s *FullNodeStub) ChainHasObj(p0 context.Context, p1 cid.Cid) (bool, error) {
	return false, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainHead(p0 context.Context) (*types.TipSet, error) {
	return s.Internal.ChainHead(p0)
}

func (s *FullNodeStub) ChainHead(p0 context.Context) (*types.TipSet, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainNotify(p0 context.Context) (<-chan []*api.HeadChange, error) {
	return s.Internal.ChainNotify(p0)
}

func (s *FullNodeStub) ChainNotify(p0 context.Context) (<-chan []*api.HeadChange, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainReadObj(p0 context.Context, p1 cid.Cid) ([]byte, error) {
	return s.Internal.ChainReadObj(p0, p1)
}

func (s *FullNodeStub) ChainReadObj(p0 context.Context, p1 cid.Cid) ([]byte, error) {
	return *new([]byte), xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainSetHead(p0 context.Context, p1 types.TipSetKey) error {
	return s.Internal.ChainSetHead(p0, p1)
}

func (s *FullNodeStub) ChainSetHead(p0 context.Context, p1 types.TipSetKey) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainStatObj(p0 context.Context, p1 cid.Cid, p2 cid.Cid) (api.ObjStat, error) {
	return s.Internal.ChainStatObj(p0, p1, p2)
}

func (s *FullNodeStub) ChainStatObj(p0 context.Context, p1 cid.Cid, p2 cid.Cid) (api.ObjStat, error) {
	return *new(api.ObjStat), xerrors.New("method not supported")
}

func (s *FullNodeStruct) ChainTipSetWeight(p0 context.Context, p1 types.TipSetKey) (types.BigInt, error) {
	return s.Internal.ChainTipSetWeight(p0, p1)
}

func (s *FullNodeStub) ChainTipSetWeight(p0 context.Context, p1 types.TipSetKey) (types.BigInt, error) {
	return *new(types.BigInt), xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientCalcCommP(p0 context.Context, p1 string) (*api.CommPRet, error) {
	return s.Internal.ClientCalcCommP(p0, p1)
}

func (s *FullNodeStub) ClientCalcCommP(p0 context.Context, p1 string) (*api.CommPRet, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientCancelDataTransfer(p0 context.Context, p1 datatransfer.TransferID, p2 peer.ID, p3 bool) error {
	return s.Internal.ClientCancelDataTransfer(p0, p1, p2, p3)
}

func (s *FullNodeStub) ClientCancelDataTransfer(p0 context.Context, p1 datatransfer.TransferID, p2 peer.ID, p3 bool) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientCancelRetrievalDeal(p0 context.Context, p1 retrievalmarket.DealID) error {
	return s.Internal.ClientCancelRetrievalDeal(p0, p1)
}

func (s *FullNodeStub) ClientCancelRetrievalDeal(p0 context.Context, p1 retrievalmarket.DealID) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientDataTransferUpdates(p0 context.Context) (<-chan api.DataTransferChannel, error) {
	return s.Internal.ClientDataTransferUpdates(p0)
}

func (s *FullNodeStub) ClientDataTransferUpdates(p0 context.Context) (<-chan api.DataTransferChannel, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientDealPieceCID(p0 context.Context, p1 cid.Cid) (api.DataCIDSize, error) {
	return s.Internal.ClientDealPieceCID(p0, p1)
}

func (s *FullNodeStub) ClientDealPieceCID(p0 context.Context, p1 cid.Cid) (api.DataCIDSize, error) {
	return *new(api.DataCIDSize), xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientDealSize(p0 context.Context, p1 cid.Cid) (api.DataSize, error) {
	return s.Internal.ClientDealSize(p0, p1)
}

func (s *FullNodeStub) ClientDealSize(p0 context.Context, p1 cid.Cid) (api.DataSize, error) {
	return *new(api.DataSize), xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientFindData(p0 context.Context, p1 cid.Cid, p2 *cid.Cid) ([]api.QueryOffer, error) {
	return s.Internal.ClientFindData(p0, p1, p2)
}

func (s *FullNodeStub) ClientFindData(p0 context.Context, p1 cid.Cid, p2 *cid.Cid) ([]api.QueryOffer, error) {
	return *new([]api.QueryOffer), xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientGenCar(p0 context.Context, p1 api.FileRef, p2 string) error {
	return s.Internal.ClientGenCar(p0, p1, p2)
}

func (s *FullNodeStub) ClientGenCar(p0 context.Context, p1 api.FileRef, p2 string) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientGetDealInfo(p0 context.Context, p1 cid.Cid) (*api.DealInfo, error) {
	return s.Internal.ClientGetDealInfo(p0, p1)
}

func (s *FullNodeStub) ClientGetDealInfo(p0 context.Context, p1 cid.Cid) (*api.DealInfo, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientGetDealStatus(p0 context.Context, p1 uint64) (string, error) {
	return s.Internal.ClientGetDealStatus(p0, p1)
}

func (s *FullNodeStub) ClientGetDealStatus(p0 context.Context, p1 uint64) (string, error) {
	return "", xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientGetDealUpdates(p0 context.Context) (<-chan api.DealInfo, error) {
	return s.Internal.ClientGetDealUpdates(p0)
}

func (s *FullNodeStub) ClientGetDealUpdates(p0 context.Context) (<-chan api.DealInfo, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientHasLocal(p0 context.Context, p1 cid.Cid) (bool, error) {
	return s.Internal.ClientHasLocal(p0, p1)
}

func (s *FullNodeStub) ClientHasLocal(p0 context.Context, p1 cid.Cid) (bool, error) {
	return false, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientImport(p0 context.Context, p1 api.FileRef) (*api.ImportRes, error) {
	return s.Internal.ClientImport(p0, p1)
}

func (s *FullNodeStub) ClientImport(p0 context.Context, p1 api.FileRef) (*api.ImportRes, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientListDataTransfers(p0 context.Context) ([]api.DataTransferChannel, error) {
	return s.Internal.ClientListDataTransfers(p0)
}

func (s *FullNodeStub) ClientListDataTransfers(p0 context.Context) ([]api.DataTransferChannel, error) {
	return *new([]api.DataTransferChannel), xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientListDeals(p0 context.Context) ([]api.DealInfo, error) {
	return s.Internal.ClientListDeals(p0)
}

func (s *FullNodeStub) ClientListDeals(p0 context.Context) ([]api.DealInfo, error) {
	return *new([]api.DealInfo), xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientListImports(p0 context.Context) ([]api.Import, error) {
	return s.Internal.ClientListImports(p0)
}

func (s *FullNodeStub) ClientListImports(p0 context.Context) ([]api.Import, error) {
	return *new([]api.Import), xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientMinerQueryOffer(p0 context.Context, p1 address.Address, p2 cid.Cid, p3 *cid.Cid) (api.QueryOffer, error) {
	return s.Internal.ClientMinerQueryOffer(p0, p1, p2, p3)
}

func (s *FullNodeStub) ClientMinerQueryOffer(p0 context.Context, p1 address.Address, p2 cid.Cid, p3 *cid.Cid) (api.QueryOffer, error) {
	return *new(api.QueryOffer), xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientQueryAsk(p0 context.Context, p1 peer.ID, p2 address.Address) (*storagemarket.StorageAsk, error) {
	return s.Internal.ClientQueryAsk(p0, p1, p2)
}

func (s *FullNodeStub) ClientQueryAsk(p0 context.Context, p1 peer.ID, p2 address.Address) (*storagemarket.StorageAsk, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientRemoveImport(p0 context.Context, p1 multistore.StoreID) error {
	return s.Internal.ClientRemoveImport(p0, p1)
}

func (s *FullNodeStub) ClientRemoveImport(p0 context.Context, p1 multistore.StoreID) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientRestartDataTransfer(p0 context.Context, p1 datatransfer.TransferID, p2 peer.ID, p3 bool) error {
	return s.Internal.ClientRestartDataTransfer(p0, p1, p2, p3)
}

func (s *FullNodeStub) ClientRestartDataTransfer(p0 context.Context, p1 datatransfer.TransferID, p2 peer.ID, p3 bool) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientRetrieve(p0 context.Context, p1 api.RetrievalOrder, p2 *api.FileRef) error {
	return s.Internal.ClientRetrieve(p0, p1, p2)
}

func (s *FullNodeStub) ClientRetrieve(p0 context.Context, p1 api.RetrievalOrder, p2 *api.FileRef) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientRetrieveTryRestartInsufficientFunds(p0 context.Context, p1 address.Address) error {
	return s.Internal.ClientRetrieveTryRestartInsufficientFunds(p0, p1)
}

func (s *FullNodeStub) ClientRetrieveTryRestartInsufficientFunds(p0 context.Context, p1 address.Address) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientRetrieveWithEvents(p0 context.Context, p1 api.RetrievalOrder, p2 *api.FileRef) (<-chan marketevents.RetrievalEvent, error) {
	return s.Internal.ClientRetrieveWithEvents(p0, p1, p2)
}

func (s *FullNodeStub) ClientRetrieveWithEvents(p0 context.Context, p1 api.RetrievalOrder, p2 *api.FileRef) (<-chan marketevents.RetrievalEvent, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) ClientStartDeal(p0 context.Context, p1 *api.StartDealParams) (*cid.Cid, error) {
	return s.Internal.ClientStartDeal(p0, p1)
}

func (s *FullNodeStub) ClientStartDeal(p0 context.Context, p1 *api.StartDealParams) (*cid.Cid, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) CreateBackup(p0 context.Context, p1 string) error {
	return s.Internal.CreateBackup(p0, p1)
}

func (s *FullNodeStub) CreateBackup(p0 context.Context, p1 string) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) GasEstimateFeeCap(p0 context.Context, p1 *types.Message, p2 int64, p3 types.TipSetKey) (types.BigInt, error) {
	return s.Internal.GasEstimateFeeCap(p0, p1, p2, p3)
}

func (s *FullNodeStub) GasEstimateFeeCap(p0 context.Context, p1 *types.Message, p2 int64, p3 types.TipSetKey) (types.BigInt, error) {
	return *new(types.BigInt), xerrors.New("method not supported")
}

func (s *FullNodeStruct) GasEstimateGasLimit(p0 context.Context, p1 *types.Message, p2 types.TipSetKey) (int64, error) {
	return s.Internal.GasEstimateGasLimit(p0, p1, p2)
}

func (s *FullNodeStub) GasEstimateGasLimit(p0 context.Context, p1 *types.Message, p2 types.TipSetKey) (int64, error) {
	return 0, xerrors.New("method not supported")
}

func (s *FullNodeStruct) GasEstimateGasPremium(p0 context.Context, p1 uint64, p2 address.Address, p3 int64, p4 types.TipSetKey) (types.BigInt, error) {
	return s.Internal.GasEstimateGasPremium(p0, p1, p2, p3, p4)
}

func (s *FullNodeStub) GasEstimateGasPremium(p0 context.Context, p1 uint64, p2 address.Address, p3 int64, p4 types.TipSetKey) (types.BigInt, error) {
	return *new(types.BigInt), xerrors.New("method not supported")
}

func (s *FullNodeStruct) GasEstimateMessageGas(p0 context.Context, p1 *types.Message, p2 *api.MessageSendSpec, p3 types.TipSetKey) (*types.Message, error) {
	return s.Internal.GasEstimateMessageGas(p0, p1, p2, p3)
}

func (s *FullNodeStub) GasEstimateMessageGas(p0 context.Context, p1 *types.Message, p2 *api.MessageSendSpec, p3 types.TipSetKey) (*types.Message, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) MarketAddBalance(p0 context.Context, p1 address.Address, p2 address.Address, p3 types.BigInt) (cid.Cid, error) {
	return s.Internal.MarketAddBalance(p0, p1, p2, p3)
}

func (s *FullNodeStub) MarketAddBalance(p0 context.Context, p1 address.Address, p2 address.Address, p3 types.BigInt) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MarketGetReserved(p0 context.Context, p1 address.Address) (types.BigInt, error) {
	return s.Internal.MarketGetReserved(p0, p1)
}

func (s *FullNodeStub) MarketGetReserved(p0 context.Context, p1 address.Address) (types.BigInt, error) {
	return *new(types.BigInt), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MarketReleaseFunds(p0 context.Context, p1 address.Address, p2 types.BigInt) error {
	return s.Internal.MarketReleaseFunds(p0, p1, p2)
}

func (s *FullNodeStub) MarketReleaseFunds(p0 context.Context, p1 address.Address, p2 types.BigInt) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) MarketReserveFunds(p0 context.Context, p1 address.Address, p2 address.Address, p3 types.BigInt) (cid.Cid, error) {
	return s.Internal.MarketReserveFunds(p0, p1, p2, p3)
}

func (s *FullNodeStub) MarketReserveFunds(p0 context.Context, p1 address.Address, p2 address.Address, p3 types.BigInt) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MarketWithdraw(p0 context.Context, p1 address.Address, p2 address.Address, p3 types.BigInt) (cid.Cid, error) {
	return s.Internal.MarketWithdraw(p0, p1, p2, p3)
}

func (s *FullNodeStub) MarketWithdraw(p0 context.Context, p1 address.Address, p2 address.Address, p3 types.BigInt) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MinerCreateBlock(p0 context.Context, p1 *api.BlockTemplate) (*types.BlockMsg, error) {
	return s.Internal.MinerCreateBlock(p0, p1)
}

func (s *FullNodeStub) MinerCreateBlock(p0 context.Context, p1 *api.BlockTemplate) (*types.BlockMsg, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) MinerGetBaseInfo(p0 context.Context, p1 address.Address, p2 abi.ChainEpoch, p3 types.TipSetKey) (*api.MiningBaseInfo, error) {
	return s.Internal.MinerGetBaseInfo(p0, p1, p2, p3)
}

func (s *FullNodeStub) MinerGetBaseInfo(p0 context.Context, p1 address.Address, p2 abi.ChainEpoch, p3 types.TipSetKey) (*api.MiningBaseInfo, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) MpoolBatchPush(p0 context.Context, p1 []*types.SignedMessage) ([]cid.Cid, error) {
	return s.Internal.MpoolBatchPush(p0, p1)
}

func (s *FullNodeStub) MpoolBatchPush(p0 context.Context, p1 []*types.SignedMessage) ([]cid.Cid, error) {
	return *new([]cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MpoolBatchPushMessage(p0 context.Context, p1 []*types.Message, p2 *api.MessageSendSpec) ([]*types.SignedMessage, error) {
	return s.Internal.MpoolBatchPushMessage(p0, p1, p2)
}

func (s *FullNodeStub) MpoolBatchPushMessage(p0 context.Context, p1 []*types.Message, p2 *api.MessageSendSpec) ([]*types.SignedMessage, error) {
	return *new([]*types.SignedMessage), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MpoolBatchPushUntrusted(p0 context.Context, p1 []*types.SignedMessage) ([]cid.Cid, error) {
	return s.Internal.MpoolBatchPushUntrusted(p0, p1)
}

func (s *FullNodeStub) MpoolBatchPushUntrusted(p0 context.Context, p1 []*types.SignedMessage) ([]cid.Cid, error) {
	return *new([]cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MpoolClear(p0 context.Context, p1 bool) error {
	return s.Internal.MpoolClear(p0, p1)
}

func (s *FullNodeStub) MpoolClear(p0 context.Context, p1 bool) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) MpoolGetConfig(p0 context.Context) (*types.MpoolConfig, error) {
	return s.Internal.MpoolGetConfig(p0)
}

func (s *FullNodeStub) MpoolGetConfig(p0 context.Context) (*types.MpoolConfig, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) MpoolGetNonce(p0 context.Context, p1 address.Address) (uint64, error) {
	return s.Internal.MpoolGetNonce(p0, p1)
}

func (s *FullNodeStub) MpoolGetNonce(p0 context.Context, p1 address.Address) (uint64, error) {
	return 0, xerrors.New("method not supported")
}

func (s *FullNodeStruct) MpoolPending(p0 context.Context, p1 types.TipSetKey) ([]*types.SignedMessage, error) {
	return s.Internal.MpoolPending(p0, p1)
}

func (s *FullNodeStub) MpoolPending(p0 context.Context, p1 types.TipSetKey) ([]*types.SignedMessage, error) {
	return *new([]*types.SignedMessage), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MpoolPush(p0 context.Context, p1 *types.SignedMessage) (cid.Cid, error) {
	return s.Internal.MpoolPush(p0, p1)
}

func (s *FullNodeStub) MpoolPush(p0 context.Context, p1 *types.SignedMessage) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MpoolPushMessage(p0 context.Context, p1 *types.Message, p2 *api.MessageSendSpec) (*types.SignedMessage, error) {
	return s.Internal.MpoolPushMessage(p0, p1, p2)
}

func (s *FullNodeStub) MpoolPushMessage(p0 context.Context, p1 *types.Message, p2 *api.MessageSendSpec) (*types.SignedMessage, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) MpoolPushUntrusted(p0 context.Context, p1 *types.SignedMessage) (cid.Cid, error) {
	return s.Internal.MpoolPushUntrusted(p0, p1)
}

func (s *FullNodeStub) MpoolPushUntrusted(p0 context.Context, p1 *types.SignedMessage) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MpoolSelect(p0 context.Context, p1 types.TipSetKey, p2 float64) ([]*types.SignedMessage, error) {
	return s.Internal.MpoolSelect(p0, p1, p2)
}

func (s *FullNodeStub) MpoolSelect(p0 context.Context, p1 types.TipSetKey, p2 float64) ([]*types.SignedMessage, error) {
	return *new([]*types.SignedMessage), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MpoolSetConfig(p0 context.Context, p1 *types.MpoolConfig) error {
	return s.Internal.MpoolSetConfig(p0, p1)
}

func (s *FullNodeStub) MpoolSetConfig(p0 context.Context, p1 *types.MpoolConfig) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) MpoolSub(p0 context.Context) (<-chan api.MpoolUpdate, error) {
	return s.Internal.MpoolSub(p0)
}

func (s *FullNodeStub) MpoolSub(p0 context.Context) (<-chan api.MpoolUpdate, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) MsigAddApprove(p0 context.Context, p1 address.Address, p2 address.Address, p3 uint64, p4 address.Address, p5 address.Address, p6 bool) (cid.Cid, error) {
	return s.Internal.MsigAddApprove(p0, p1, p2, p3, p4, p5, p6)
}

func (s *FullNodeStub) MsigAddApprove(p0 context.Context, p1 address.Address, p2 address.Address, p3 uint64, p4 address.Address, p5 address.Address, p6 bool) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MsigAddCancel(p0 context.Context, p1 address.Address, p2 address.Address, p3 uint64, p4 address.Address, p5 bool) (cid.Cid, error) {
	return s.Internal.MsigAddCancel(p0, p1, p2, p3, p4, p5)
}

func (s *FullNodeStub) MsigAddCancel(p0 context.Context, p1 address.Address, p2 address.Address, p3 uint64, p4 address.Address, p5 bool) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MsigAddPropose(p0 context.Context, p1 address.Address, p2 address.Address, p3 address.Address, p4 bool) (cid.Cid, error) {
	return s.Internal.MsigAddPropose(p0, p1, p2, p3, p4)
}

func (s *FullNodeStub) MsigAddPropose(p0 context.Context, p1 address.Address, p2 address.Address, p3 address.Address, p4 bool) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MsigApprove(p0 context.Context, p1 address.Address, p2 uint64, p3 address.Address) (cid.Cid, error) {
	return s.Internal.MsigApprove(p0, p1, p2, p3)
}

func (s *FullNodeStub) MsigApprove(p0 context.Context, p1 address.Address, p2 uint64, p3 address.Address) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MsigApproveTxnHash(p0 context.Context, p1 address.Address, p2 uint64, p3 address.Address, p4 address.Address, p5 types.BigInt, p6 address.Address, p7 uint64, p8 []byte) (cid.Cid, error) {
	return s.Internal.MsigApproveTxnHash(p0, p1, p2, p3, p4, p5, p6, p7, p8)
}

func (s *FullNodeStub) MsigApproveTxnHash(p0 context.Context, p1 address.Address, p2 uint64, p3 address.Address, p4 address.Address, p5 types.BigInt, p6 address.Address, p7 uint64, p8 []byte) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MsigCancel(p0 context.Context, p1 address.Address, p2 uint64, p3 address.Address, p4 types.BigInt, p5 address.Address, p6 uint64, p7 []byte) (cid.Cid, error) {
	return s.Internal.MsigCancel(p0, p1, p2, p3, p4, p5, p6, p7)
}

func (s *FullNodeStub) MsigCancel(p0 context.Context, p1 address.Address, p2 uint64, p3 address.Address, p4 types.BigInt, p5 address.Address, p6 uint64, p7 []byte) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MsigCreate(p0 context.Context, p1 uint64, p2 []address.Address, p3 abi.ChainEpoch, p4 types.BigInt, p5 address.Address, p6 types.BigInt) (cid.Cid, error) {
	return s.Internal.MsigCreate(p0, p1, p2, p3, p4, p5, p6)
}

func (s *FullNodeStub) MsigCreate(p0 context.Context, p1 uint64, p2 []address.Address, p3 abi.ChainEpoch, p4 types.BigInt, p5 address.Address, p6 types.BigInt) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MsigGetAvailableBalance(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (types.BigInt, error) {
	return s.Internal.MsigGetAvailableBalance(p0, p1, p2)
}

func (s *FullNodeStub) MsigGetAvailableBalance(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (types.BigInt, error) {
	return *new(types.BigInt), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MsigGetPending(p0 context.Context, p1 address.Address, p2 types.TipSetKey) ([]*api.MsigTransaction, error) {
	return s.Internal.MsigGetPending(p0, p1, p2)
}

func (s *FullNodeStub) MsigGetPending(p0 context.Context, p1 address.Address, p2 types.TipSetKey) ([]*api.MsigTransaction, error) {
	return *new([]*api.MsigTransaction), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MsigGetVested(p0 context.Context, p1 address.Address, p2 types.TipSetKey, p3 types.TipSetKey) (types.BigInt, error) {
	return s.Internal.MsigGetVested(p0, p1, p2, p3)
}

func (s *FullNodeStub) MsigGetVested(p0 context.Context, p1 address.Address, p2 types.TipSetKey, p3 types.TipSetKey) (types.BigInt, error) {
	return *new(types.BigInt), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MsigGetVestingSchedule(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (api.MsigVesting, error) {
	return s.Internal.MsigGetVestingSchedule(p0, p1, p2)
}

func (s *FullNodeStub) MsigGetVestingSchedule(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (api.MsigVesting, error) {
	return *new(api.MsigVesting), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MsigPropose(p0 context.Context, p1 address.Address, p2 address.Address, p3 types.BigInt, p4 address.Address, p5 uint64, p6 []byte) (cid.Cid, error) {
	return s.Internal.MsigPropose(p0, p1, p2, p3, p4, p5, p6)
}

func (s *FullNodeStub) MsigPropose(p0 context.Context, p1 address.Address, p2 address.Address, p3 types.BigInt, p4 address.Address, p5 uint64, p6 []byte) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MsigRemoveSigner(p0 context.Context, p1 address.Address, p2 address.Address, p3 address.Address, p4 bool) (cid.Cid, error) {
	return s.Internal.MsigRemoveSigner(p0, p1, p2, p3, p4)
}

func (s *FullNodeStub) MsigRemoveSigner(p0 context.Context, p1 address.Address, p2 address.Address, p3 address.Address, p4 bool) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MsigSwapApprove(p0 context.Context, p1 address.Address, p2 address.Address, p3 uint64, p4 address.Address, p5 address.Address, p6 address.Address) (cid.Cid, error) {
	return s.Internal.MsigSwapApprove(p0, p1, p2, p3, p4, p5, p6)
}

func (s *FullNodeStub) MsigSwapApprove(p0 context.Context, p1 address.Address, p2 address.Address, p3 uint64, p4 address.Address, p5 address.Address, p6 address.Address) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MsigSwapCancel(p0 context.Context, p1 address.Address, p2 address.Address, p3 uint64, p4 address.Address, p5 address.Address) (cid.Cid, error) {
	return s.Internal.MsigSwapCancel(p0, p1, p2, p3, p4, p5)
}

func (s *FullNodeStub) MsigSwapCancel(p0 context.Context, p1 address.Address, p2 address.Address, p3 uint64, p4 address.Address, p5 address.Address) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) MsigSwapPropose(p0 context.Context, p1 address.Address, p2 address.Address, p3 address.Address, p4 address.Address) (cid.Cid, error) {
	return s.Internal.MsigSwapPropose(p0, p1, p2, p3, p4)
}

func (s *FullNodeStub) MsigSwapPropose(p0 context.Context, p1 address.Address, p2 address.Address, p3 address.Address, p4 address.Address) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) PaychAllocateLane(p0 context.Context, p1 address.Address) (uint64, error) {
	return s.Internal.PaychAllocateLane(p0, p1)
}

func (s *FullNodeStub) PaychAllocateLane(p0 context.Context, p1 address.Address) (uint64, error) {
	return 0, xerrors.New("method not supported")
}

func (s *FullNodeStruct) PaychAvailableFunds(p0 context.Context, p1 address.Address) (*api.ChannelAvailableFunds, error) {
	return s.Internal.PaychAvailableFunds(p0, p1)
}

func (s *FullNodeStub) PaychAvailableFunds(p0 context.Context, p1 address.Address) (*api.ChannelAvailableFunds, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) PaychAvailableFundsByFromTo(p0 context.Context, p1 address.Address, p2 address.Address) (*api.ChannelAvailableFunds, error) {
	return s.Internal.PaychAvailableFundsByFromTo(p0, p1, p2)
}

func (s *FullNodeStub) PaychAvailableFundsByFromTo(p0 context.Context, p1 address.Address, p2 address.Address) (*api.ChannelAvailableFunds, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) PaychCollect(p0 context.Context, p1 address.Address) (cid.Cid, error) {
	return s.Internal.PaychCollect(p0, p1)
}

func (s *FullNodeStub) PaychCollect(p0 context.Context, p1 address.Address) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) PaychGet(p0 context.Context, p1 address.Address, p2 address.Address, p3 types.BigInt) (*api.ChannelInfo, error) {
	return s.Internal.PaychGet(p0, p1, p2, p3)
}

func (s *FullNodeStub) PaychGet(p0 context.Context, p1 address.Address, p2 address.Address, p3 types.BigInt) (*api.ChannelInfo, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) PaychGetWaitReady(p0 context.Context, p1 cid.Cid) (address.Address, error) {
	return s.Internal.PaychGetWaitReady(p0, p1)
}

func (s *FullNodeStub) PaychGetWaitReady(p0 context.Context, p1 cid.Cid) (address.Address, error) {
	return *new(address.Address), xerrors.New("method not supported")
}

func (s *FullNodeStruct) PaychList(p0 context.Context) ([]address.Address, error) {
	return s.Internal.PaychList(p0)
}

func (s *FullNodeStub) PaychList(p0 context.Context) ([]address.Address, error) {
	return *new([]address.Address), xerrors.New("method not supported")
}

func (s *FullNodeStruct) PaychNewPayment(p0 context.Context, p1 address.Address, p2 address.Address, p3 []api.VoucherSpec) (*api.PaymentInfo, error) {
	return s.Internal.PaychNewPayment(p0, p1, p2, p3)
}

func (s *FullNodeStub) PaychNewPayment(p0 context.Context, p1 address.Address, p2 address.Address, p3 []api.VoucherSpec) (*api.PaymentInfo, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) PaychSettle(p0 context.Context, p1 address.Address) (cid.Cid, error) {
	return s.Internal.PaychSettle(p0, p1)
}

func (s *FullNodeStub) PaychSettle(p0 context.Context, p1 address.Address) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) PaychStatus(p0 context.Context, p1 address.Address) (*api.PaychStatus, error) {
	return s.Internal.PaychStatus(p0, p1)
}

func (s *FullNodeStub) PaychStatus(p0 context.Context, p1 address.Address) (*api.PaychStatus, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) PaychVoucherAdd(p0 context.Context, p1 address.Address, p2 *paych.SignedVoucher, p3 []byte, p4 types.BigInt) (types.BigInt, error) {
	return s.Internal.PaychVoucherAdd(p0, p1, p2, p3, p4)
}

func (s *FullNodeStub) PaychVoucherAdd(p0 context.Context, p1 address.Address, p2 *paych.SignedVoucher, p3 []byte, p4 types.BigInt) (types.BigInt, error) {
	return *new(types.BigInt), xerrors.New("method not supported")
}

func (s *FullNodeStruct) PaychVoucherCheckSpendable(p0 context.Context, p1 address.Address, p2 *paych.SignedVoucher, p3 []byte, p4 []byte) (bool, error) {
	return s.Internal.PaychVoucherCheckSpendable(p0, p1, p2, p3, p4)
}

func (s *FullNodeStub) PaychVoucherCheckSpendable(p0 context.Context, p1 address.Address, p2 *paych.SignedVoucher, p3 []byte, p4 []byte) (bool, error) {
	return false, xerrors.New("method not supported")
}

func (s *FullNodeStruct) PaychVoucherCheckValid(p0 context.Context, p1 address.Address, p2 *paych.SignedVoucher) error {
	return s.Internal.PaychVoucherCheckValid(p0, p1, p2)
}

func (s *FullNodeStub) PaychVoucherCheckValid(p0 context.Context, p1 address.Address, p2 *paych.SignedVoucher) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) PaychVoucherCreate(p0 context.Context, p1 address.Address, p2 types.BigInt, p3 uint64) (*api.VoucherCreateResult, error) {
	return s.Internal.PaychVoucherCreate(p0, p1, p2, p3)
}

func (s *FullNodeStub) PaychVoucherCreate(p0 context.Context, p1 address.Address, p2 types.BigInt, p3 uint64) (*api.VoucherCreateResult, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) PaychVoucherList(p0 context.Context, p1 address.Address) ([]*paych.SignedVoucher, error) {
	return s.Internal.PaychVoucherList(p0, p1)
}

func (s *FullNodeStub) PaychVoucherList(p0 context.Context, p1 address.Address) ([]*paych.SignedVoucher, error) {
	return *new([]*paych.SignedVoucher), xerrors.New("method not supported")
}

func (s *FullNodeStruct) PaychVoucherSubmit(p0 context.Context, p1 address.Address, p2 *paych.SignedVoucher, p3 []byte, p4 []byte) (cid.Cid, error) {
	return s.Internal.PaychVoucherSubmit(p0, p1, p2, p3, p4)
}

func (s *FullNodeStub) PaychVoucherSubmit(p0 context.Context, p1 address.Address, p2 *paych.SignedVoucher, p3 []byte, p4 []byte) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateAccountKey(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (address.Address, error) {
	return s.Internal.StateAccountKey(p0, p1, p2)
}

func (s *FullNodeStub) StateAccountKey(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (address.Address, error) {
	return *new(address.Address), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateAllMinerFaults(p0 context.Context, p1 abi.ChainEpoch, p2 types.TipSetKey) ([]*api.Fault, error) {
	return s.Internal.StateAllMinerFaults(p0, p1, p2)
}

func (s *FullNodeStub) StateAllMinerFaults(p0 context.Context, p1 abi.ChainEpoch, p2 types.TipSetKey) ([]*api.Fault, error) {
	return *new([]*api.Fault), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateCall(p0 context.Context, p1 *types.Message, p2 types.TipSetKey) (*api.InvocResult, error) {
	return s.Internal.StateCall(p0, p1, p2)
}

func (s *FullNodeStub) StateCall(p0 context.Context, p1 *types.Message, p2 types.TipSetKey) (*api.InvocResult, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateChangedActors(p0 context.Context, p1 cid.Cid, p2 cid.Cid) (map[string]types.Actor, error) {
	return s.Internal.StateChangedActors(p0, p1, p2)
}

func (s *FullNodeStub) StateChangedActors(p0 context.Context, p1 cid.Cid, p2 cid.Cid) (map[string]types.Actor, error) {
	return *new(map[string]types.Actor), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateCirculatingSupply(p0 context.Context, p1 types.TipSetKey) (abi.TokenAmount, error) {
	return s.Internal.StateCirculatingSupply(p0, p1)
}

func (s *FullNodeStub) StateCirculatingSupply(p0 context.Context, p1 types.TipSetKey) (abi.TokenAmount, error) {
	return *new(abi.TokenAmount), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateCompute(p0 context.Context, p1 abi.ChainEpoch, p2 []*types.Message, p3 types.TipSetKey) (*api.ComputeStateOutput, error) {
	return s.Internal.StateCompute(p0, p1, p2, p3)
}

func (s *FullNodeStub) StateCompute(p0 context.Context, p1 abi.ChainEpoch, p2 []*types.Message, p3 types.TipSetKey) (*api.ComputeStateOutput, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateDealProviderCollateralBounds(p0 context.Context, p1 abi.PaddedPieceSize, p2 bool, p3 types.TipSetKey) (api.DealCollateralBounds, error) {
	return s.Internal.StateDealProviderCollateralBounds(p0, p1, p2, p3)
}

func (s *FullNodeStub) StateDealProviderCollateralBounds(p0 context.Context, p1 abi.PaddedPieceSize, p2 bool, p3 types.TipSetKey) (api.DealCollateralBounds, error) {
	return *new(api.DealCollateralBounds), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateDecodeParams(p0 context.Context, p1 address.Address, p2 abi.MethodNum, p3 []byte, p4 types.TipSetKey) (interface{}, error) {
	return s.Internal.StateDecodeParams(p0, p1, p2, p3, p4)
}

func (s *FullNodeStub) StateDecodeParams(p0 context.Context, p1 address.Address, p2 abi.MethodNum, p3 []byte, p4 types.TipSetKey) (interface{}, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateGetActor(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*types.Actor, error) {
	return s.Internal.StateGetActor(p0, p1, p2)
}

func (s *FullNodeStub) StateGetActor(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*types.Actor, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateGetReceipt(p0 context.Context, p1 cid.Cid, p2 types.TipSetKey) (*types.MessageReceipt, error) {
	return s.Internal.StateGetReceipt(p0, p1, p2)
}

func (s *FullNodeStub) StateGetReceipt(p0 context.Context, p1 cid.Cid, p2 types.TipSetKey) (*types.MessageReceipt, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateListActors(p0 context.Context, p1 types.TipSetKey) ([]address.Address, error) {
	return s.Internal.StateListActors(p0, p1)
}

func (s *FullNodeStub) StateListActors(p0 context.Context, p1 types.TipSetKey) ([]address.Address, error) {
	return *new([]address.Address), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateListMessages(p0 context.Context, p1 *api.MessageMatch, p2 types.TipSetKey, p3 abi.ChainEpoch) ([]cid.Cid, error) {
	return s.Internal.StateListMessages(p0, p1, p2, p3)
}

func (s *FullNodeStub) StateListMessages(p0 context.Context, p1 *api.MessageMatch, p2 types.TipSetKey, p3 abi.ChainEpoch) ([]cid.Cid, error) {
	return *new([]cid.Cid), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateListMiners(p0 context.Context, p1 types.TipSetKey) ([]address.Address, error) {
	return s.Internal.StateListMiners(p0, p1)
}

func (s *FullNodeStub) StateListMiners(p0 context.Context, p1 types.TipSetKey) ([]address.Address, error) {
	return *new([]address.Address), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateLookupID(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (address.Address, error) {
	return s.Internal.StateLookupID(p0, p1, p2)
}

func (s *FullNodeStub) StateLookupID(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (address.Address, error) {
	return *new(address.Address), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMarketBalance(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (api.MarketBalance, error) {
	return s.Internal.StateMarketBalance(p0, p1, p2)
}

func (s *FullNodeStub) StateMarketBalance(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (api.MarketBalance, error) {
	return *new(api.MarketBalance), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMarketDeals(p0 context.Context, p1 types.TipSetKey) (map[string]api.MarketDeal, error) {
	return s.Internal.StateMarketDeals(p0, p1)
}

func (s *FullNodeStub) StateMarketDeals(p0 context.Context, p1 types.TipSetKey) (map[string]api.MarketDeal, error) {
	return *new(map[string]api.MarketDeal), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMarketParticipants(p0 context.Context, p1 types.TipSetKey) (map[string]api.MarketBalance, error) {
	return s.Internal.StateMarketParticipants(p0, p1)
}

func (s *FullNodeStub) StateMarketParticipants(p0 context.Context, p1 types.TipSetKey) (map[string]api.MarketBalance, error) {
	return *new(map[string]api.MarketBalance), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMarketStorageDeal(p0 context.Context, p1 abi.DealID, p2 types.TipSetKey) (*api.MarketDeal, error) {
	return s.Internal.StateMarketStorageDeal(p0, p1, p2)
}

func (s *FullNodeStub) StateMarketStorageDeal(p0 context.Context, p1 abi.DealID, p2 types.TipSetKey) (*api.MarketDeal, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMinerActiveSectors(p0 context.Context, p1 address.Address, p2 types.TipSetKey) ([]*miner.SectorOnChainInfo, error) {
	return s.Internal.StateMinerActiveSectors(p0, p1, p2)
}

func (s *FullNodeStub) StateMinerActiveSectors(p0 context.Context, p1 address.Address, p2 types.TipSetKey) ([]*miner.SectorOnChainInfo, error) {
	return *new([]*miner.SectorOnChainInfo), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMinerAvailableBalance(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (types.BigInt, error) {
	return s.Internal.StateMinerAvailableBalance(p0, p1, p2)
}

func (s *FullNodeStub) StateMinerAvailableBalance(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (types.BigInt, error) {
	return *new(types.BigInt), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMinerDeadlines(p0 context.Context, p1 address.Address, p2 types.TipSetKey) ([]api.Deadline, error) {
	return s.Internal.StateMinerDeadlines(p0, p1, p2)
}

func (s *FullNodeStub) StateMinerDeadlines(p0 context.Context, p1 address.Address, p2 types.TipSetKey) ([]api.Deadline, error) {
	return *new([]api.Deadline), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMinerFaults(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (bitfield.BitField, error) {
	return s.Internal.StateMinerFaults(p0, p1, p2)
}

func (s *FullNodeStub) StateMinerFaults(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (bitfield.BitField, error) {
	return *new(bitfield.BitField), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMinerInfo(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (miner.MinerInfo, error) {
	return s.Internal.StateMinerInfo(p0, p1, p2)
}

func (s *FullNodeStub) StateMinerInfo(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (miner.MinerInfo, error) {
	return *new(miner.MinerInfo), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMinerInitialPledgeCollateral(p0 context.Context, p1 address.Address, p2 miner.SectorPreCommitInfo, p3 types.TipSetKey) (types.BigInt, error) {
	return s.Internal.StateMinerInitialPledgeCollateral(p0, p1, p2, p3)
}

func (s *FullNodeStub) StateMinerInitialPledgeCollateral(p0 context.Context, p1 address.Address, p2 miner.SectorPreCommitInfo, p3 types.TipSetKey) (types.BigInt, error) {
	return *new(types.BigInt), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMinerPartitions(p0 context.Context, p1 address.Address, p2 uint64, p3 types.TipSetKey) ([]api.Partition, error) {
	return s.Internal.StateMinerPartitions(p0, p1, p2, p3)
}

func (s *FullNodeStub) StateMinerPartitions(p0 context.Context, p1 address.Address, p2 uint64, p3 types.TipSetKey) ([]api.Partition, error) {
	return *new([]api.Partition), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMinerPower(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*api.MinerPower, error) {
	return s.Internal.StateMinerPower(p0, p1, p2)
}

func (s *FullNodeStub) StateMinerPower(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*api.MinerPower, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMinerPreCommitDepositForPower(p0 context.Context, p1 address.Address, p2 miner.SectorPreCommitInfo, p3 types.TipSetKey) (types.BigInt, error) {
	return s.Internal.StateMinerPreCommitDepositForPower(p0, p1, p2, p3)
}

func (s *FullNodeStub) StateMinerPreCommitDepositForPower(p0 context.Context, p1 address.Address, p2 miner.SectorPreCommitInfo, p3 types.TipSetKey) (types.BigInt, error) {
	return *new(types.BigInt), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMinerProvingDeadline(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*dline.Info, error) {
	return s.Internal.StateMinerProvingDeadline(p0, p1, p2)
}

func (s *FullNodeStub) StateMinerProvingDeadline(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*dline.Info, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMinerRecoveries(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (bitfield.BitField, error) {
	return s.Internal.StateMinerRecoveries(p0, p1, p2)
}

func (s *FullNodeStub) StateMinerRecoveries(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (bitfield.BitField, error) {
	return *new(bitfield.BitField), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMinerSectorAllocated(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (bool, error) {
	return s.Internal.StateMinerSectorAllocated(p0, p1, p2, p3)
}

func (s *FullNodeStub) StateMinerSectorAllocated(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (bool, error) {
	return false, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMinerSectorCount(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (api.MinerSectors, error) {
	return s.Internal.StateMinerSectorCount(p0, p1, p2)
}

func (s *FullNodeStub) StateMinerSectorCount(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (api.MinerSectors, error) {
	return *new(api.MinerSectors), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateMinerSectors(p0 context.Context, p1 address.Address, p2 *bitfield.BitField, p3 types.TipSetKey) ([]*miner.SectorOnChainInfo, error) {
	return s.Internal.StateMinerSectors(p0, p1, p2, p3)
}

func (s *FullNodeStub) StateMinerSectors(p0 context.Context, p1 address.Address, p2 *bitfield.BitField, p3 types.TipSetKey) ([]*miner.SectorOnChainInfo, error) {
	return *new([]*miner.SectorOnChainInfo), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateNetworkName(p0 context.Context) (dtypes.NetworkName, error) {
	return s.Internal.StateNetworkName(p0)
}

func (s *FullNodeStub) StateNetworkName(p0 context.Context) (dtypes.NetworkName, error) {
	return *new(dtypes.NetworkName), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateNetworkVersion(p0 context.Context, p1 types.TipSetKey) (apitypes.NetworkVersion, error) {
	return s.Internal.StateNetworkVersion(p0, p1)
}

func (s *FullNodeStub) StateNetworkVersion(p0 context.Context, p1 types.TipSetKey) (apitypes.NetworkVersion, error) {
	return *new(apitypes.NetworkVersion), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateReadState(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*api.ActorState, error) {
	return s.Internal.StateReadState(p0, p1, p2)
}

func (s *FullNodeStub) StateReadState(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*api.ActorState, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateReplay(p0 context.Context, p1 types.TipSetKey, p2 cid.Cid) (*api.InvocResult, error) {
	return s.Internal.StateReplay(p0, p1, p2)
}

func (s *FullNodeStub) StateReplay(p0 context.Context, p1 types.TipSetKey, p2 cid.Cid) (*api.InvocResult, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateSearchMsg(p0 context.Context, p1 cid.Cid) (*api.MsgLookup, error) {
	return s.Internal.StateSearchMsg(p0, p1)
}

func (s *FullNodeStub) StateSearchMsg(p0 context.Context, p1 cid.Cid) (*api.MsgLookup, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateSearchMsgLimited(p0 context.Context, p1 cid.Cid, p2 abi.ChainEpoch) (*api.MsgLookup, error) {
	return s.Internal.StateSearchMsgLimited(p0, p1, p2)
}

func (s *FullNodeStub) StateSearchMsgLimited(p0 context.Context, p1 cid.Cid, p2 abi.ChainEpoch) (*api.MsgLookup, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateSectorExpiration(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (*miner.SectorExpiration, error) {
	return s.Internal.StateSectorExpiration(p0, p1, p2, p3)
}

func (s *FullNodeStub) StateSectorExpiration(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (*miner.SectorExpiration, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateSectorGetInfo(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (*miner.SectorOnChainInfo, error) {
	return s.Internal.StateSectorGetInfo(p0, p1, p2, p3)
}

func (s *FullNodeStub) StateSectorGetInfo(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (*miner.SectorOnChainInfo, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateSectorPartition(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (*miner.SectorLocation, error) {
	return s.Internal.StateSectorPartition(p0, p1, p2, p3)
}

func (s *FullNodeStub) StateSectorPartition(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (*miner.SectorLocation, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateSectorPreCommitInfo(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (miner.SectorPreCommitOnChainInfo, error) {
	return s.Internal.StateSectorPreCommitInfo(p0, p1, p2, p3)
}

func (s *FullNodeStub) StateSectorPreCommitInfo(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (miner.SectorPreCommitOnChainInfo, error) {
	return *new(miner.SectorPreCommitOnChainInfo), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateVMCirculatingSupplyInternal(p0 context.Context, p1 types.TipSetKey) (api.CirculatingSupply, error) {
	return s.Internal.StateVMCirculatingSupplyInternal(p0, p1)
}

func (s *FullNodeStub) StateVMCirculatingSupplyInternal(p0 context.Context, p1 types.TipSetKey) (api.CirculatingSupply, error) {
	return *new(api.CirculatingSupply), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateVerifiedClientStatus(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*abi.StoragePower, error) {
	return s.Internal.StateVerifiedClientStatus(p0, p1, p2)
}

func (s *FullNodeStub) StateVerifiedClientStatus(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*abi.StoragePower, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateVerifiedRegistryRootKey(p0 context.Context, p1 types.TipSetKey) (address.Address, error) {
	return s.Internal.StateVerifiedRegistryRootKey(p0, p1)
}

func (s *FullNodeStub) StateVerifiedRegistryRootKey(p0 context.Context, p1 types.TipSetKey) (address.Address, error) {
	return *new(address.Address), xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateVerifierStatus(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*abi.StoragePower, error) {
	return s.Internal.StateVerifierStatus(p0, p1, p2)
}

func (s *FullNodeStub) StateVerifierStatus(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*abi.StoragePower, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateWaitMsg(p0 context.Context, p1 cid.Cid, p2 uint64) (*api.MsgLookup, error) {
	return s.Internal.StateWaitMsg(p0, p1, p2)
}

func (s *FullNodeStub) StateWaitMsg(p0 context.Context, p1 cid.Cid, p2 uint64) (*api.MsgLookup, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) StateWaitMsgLimited(p0 context.Context, p1 cid.Cid, p2 uint64, p3 abi.ChainEpoch) (*api.MsgLookup, error) {
	return s.Internal.StateWaitMsgLimited(p0, p1, p2, p3)
}

func (s *FullNodeStub) StateWaitMsgLimited(p0 context.Context, p1 cid.Cid, p2 uint64, p3 abi.ChainEpoch) (*api.MsgLookup, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) SyncCheckBad(p0 context.Context, p1 cid.Cid) (string, error) {
	return s.Internal.SyncCheckBad(p0, p1)
}

func (s *FullNodeStub) SyncCheckBad(p0 context.Context, p1 cid.Cid) (string, error) {
	return "", xerrors.New("method not supported")
}

func (s *FullNodeStruct) SyncCheckpoint(p0 context.Context, p1 types.TipSetKey) error {
	return s.Internal.SyncCheckpoint(p0, p1)
}

func (s *FullNodeStub) SyncCheckpoint(p0 context.Context, p1 types.TipSetKey) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) SyncIncomingBlocks(p0 context.Context) (<-chan *types.BlockHeader, error) {
	return s.Internal.SyncIncomingBlocks(p0)
}

func (s *FullNodeStub) SyncIncomingBlocks(p0 context.Context) (<-chan *types.BlockHeader, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) SyncMarkBad(p0 context.Context, p1 cid.Cid) error {
	return s.Internal.SyncMarkBad(p0, p1)
}

func (s *FullNodeStub) SyncMarkBad(p0 context.Context, p1 cid.Cid) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) SyncState(p0 context.Context) (*api.SyncState, error) {
	return s.Internal.SyncState(p0)
}

func (s *FullNodeStub) SyncState(p0 context.Context) (*api.SyncState, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) SyncSubmitBlock(p0 context.Context, p1 *types.BlockMsg) error {
	return s.Internal.SyncSubmitBlock(p0, p1)
}

func (s *FullNodeStub) SyncSubmitBlock(p0 context.Context, p1 *types.BlockMsg) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) SyncUnmarkAllBad(p0 context.Context) error {
	return s.Internal.SyncUnmarkAllBad(p0)
}

func (s *FullNodeStub) SyncUnmarkAllBad(p0 context.Context) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) SyncUnmarkBad(p0 context.Context, p1 cid.Cid) error {
	return s.Internal.SyncUnmarkBad(p0, p1)
}

func (s *FullNodeStub) SyncUnmarkBad(p0 context.Context, p1 cid.Cid) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) SyncValidateTipset(p0 context.Context, p1 types.TipSetKey) (bool, error) {
	return s.Internal.SyncValidateTipset(p0, p1)
}

func (s *FullNodeStub) SyncValidateTipset(p0 context.Context, p1 types.TipSetKey) (bool, error) {
	return false, xerrors.New("method not supported")
}

func (s *FullNodeStruct) WalletBalance(p0 context.Context, p1 address.Address) (types.BigInt, error) {
	return s.Internal.WalletBalance(p0, p1)
}

func (s *FullNodeStub) WalletBalance(p0 context.Context, p1 address.Address) (types.BigInt, error) {
	return *new(types.BigInt), xerrors.New("method not supported")
}

func (s *FullNodeStruct) WalletDefaultAddress(p0 context.Context) (address.Address, error) {
	return s.Internal.WalletDefaultAddress(p0)
}

func (s *FullNodeStub) WalletDefaultAddress(p0 context.Context) (address.Address, error) {
	return *new(address.Address), xerrors.New("method not supported")
}

func (s *FullNodeStruct) WalletDelete(p0 context.Context, p1 address.Address) error {
	return s.Internal.WalletDelete(p0, p1)
}

func (s *FullNodeStub) WalletDelete(p0 context.Context, p1 address.Address) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) WalletExport(p0 context.Context, p1 address.Address) (*types.KeyInfo, error) {
	return s.Internal.WalletExport(p0, p1)
}

func (s *FullNodeStub) WalletExport(p0 context.Context, p1 address.Address) (*types.KeyInfo, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) WalletHas(p0 context.Context, p1 address.Address) (bool, error) {
	return s.Internal.WalletHas(p0, p1)
}

func (s *FullNodeStub) WalletHas(p0 context.Context, p1 address.Address) (bool, error) {
	return false, xerrors.New("method not supported")
}

func (s *FullNodeStruct) WalletImport(p0 context.Context, p1 *types.KeyInfo) (address.Address, error) {
	return s.Internal.WalletImport(p0, p1)
}

func (s *FullNodeStub) WalletImport(p0 context.Context, p1 *types.KeyInfo) (address.Address, error) {
	return *new(address.Address), xerrors.New("method not supported")
}

func (s *FullNodeStruct) WalletList(p0 context.Context) ([]address.Address, error) {
	return s.Internal.WalletList(p0)
}

func (s *FullNodeStub) WalletList(p0 context.Context) ([]address.Address, error) {
	return *new([]address.Address), xerrors.New("method not supported")
}

func (s *FullNodeStruct) WalletNew(p0 context.Context, p1 types.KeyType) (address.Address, error) {
	return s.Internal.WalletNew(p0, p1)
}

func (s *FullNodeStub) WalletNew(p0 context.Context, p1 types.KeyType) (address.Address, error) {
	return *new(address.Address), xerrors.New("method not supported")
}

func (s *FullNodeStruct) WalletSetDefault(p0 context.Context, p1 address.Address) error {
	return s.Internal.WalletSetDefault(p0, p1)
}

func (s *FullNodeStub) WalletSetDefault(p0 context.Context, p1 address.Address) error {
	return xerrors.New("method not supported")
}

func (s *FullNodeStruct) WalletSign(p0 context.Context, p1 address.Address, p2 []byte) (*crypto.Signature, error) {
	return s.Internal.WalletSign(p0, p1, p2)
}

func (s *FullNodeStub) WalletSign(p0 context.Context, p1 address.Address, p2 []byte) (*crypto.Signature, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) WalletSignMessage(p0 context.Context, p1 address.Address, p2 *types.Message) (*types.SignedMessage, error) {
	return s.Internal.WalletSignMessage(p0, p1, p2)
}

func (s *FullNodeStub) WalletSignMessage(p0 context.Context, p1 address.Address, p2 *types.Message) (*types.SignedMessage, error) {
	return nil, xerrors.New("method not supported")
}

func (s *FullNodeStruct) WalletValidateAddress(p0 context.Context, p1 string) (address.Address, error) {
	return s.Internal.WalletValidateAddress(p0, p1)
}

func (s *FullNodeStub) WalletValidateAddress(p0 context.Context, p1 string) (address.Address, error) {
	return *new(address.Address), xerrors.New("method not supported")
}

func (s *FullNodeStruct) WalletVerify(p0 context.Context, p1 address.Address, p2 []byte, p3 *crypto.Signature) (bool, error) {
	return s.Internal.WalletVerify(p0, p1, p2, p3)
}

func (s *FullNodeStub) WalletVerify(p0 context.Context, p1 address.Address, p2 []byte, p3 *crypto.Signature) (bool, error) {
	return false, xerrors.New("method not supported")
}

func (s *GatewayStruct) ChainGetBlockMessages(p0 context.Context, p1 cid.Cid) (*api.BlockMessages, error) {
	return s.Internal.ChainGetBlockMessages(p0, p1)
}

func (s *GatewayStub) ChainGetBlockMessages(p0 context.Context, p1 cid.Cid) (*api.BlockMessages, error) {
	return nil, xerrors.New("method not supported")
}

func (s *GatewayStruct) ChainGetMessage(p0 context.Context, p1 cid.Cid) (*types.Message, error) {
	return s.Internal.ChainGetMessage(p0, p1)
}

func (s *GatewayStub) ChainGetMessage(p0 context.Context, p1 cid.Cid) (*types.Message, error) {
	return nil, xerrors.New("method not supported")
}

func (s *GatewayStruct) ChainGetTipSet(p0 context.Context, p1 types.TipSetKey) (*types.TipSet, error) {
	return s.Internal.ChainGetTipSet(p0, p1)
}

func (s *GatewayStub) ChainGetTipSet(p0 context.Context, p1 types.TipSetKey) (*types.TipSet, error) {
	return nil, xerrors.New("method not supported")
}

func (s *GatewayStruct) ChainGetTipSetByHeight(p0 context.Context, p1 abi.ChainEpoch, p2 types.TipSetKey) (*types.TipSet, error) {
	return s.Internal.ChainGetTipSetByHeight(p0, p1, p2)
}

func (s *GatewayStub) ChainGetTipSetByHeight(p0 context.Context, p1 abi.ChainEpoch, p2 types.TipSetKey) (*types.TipSet, error) {
	return nil, xerrors.New("method not supported")
}

func (s *GatewayStruct) ChainHasObj(p0 context.Context, p1 cid.Cid) (bool, error) {
	return s.Internal.ChainHasObj(p0, p1)
}

func (s *GatewayStub) ChainHasObj(p0 context.Context, p1 cid.Cid) (bool, error) {
	return false, xerrors.New("method not supported")
}

func (s *GatewayStruct) ChainHead(p0 context.Context) (*types.TipSet, error) {
	return s.Internal.ChainHead(p0)
}

func (s *GatewayStub) ChainHead(p0 context.Context) (*types.TipSet, error) {
	return nil, xerrors.New("method not supported")
}

func (s *GatewayStruct) ChainNotify(p0 context.Context) (<-chan []*api.HeadChange, error) {
	return s.Internal.ChainNotify(p0)
}

func (s *GatewayStub) ChainNotify(p0 context.Context) (<-chan []*api.HeadChange, error) {
	return nil, xerrors.New("method not supported")
}

func (s *GatewayStruct) ChainReadObj(p0 context.Context, p1 cid.Cid) ([]byte, error) {
	return s.Internal.ChainReadObj(p0, p1)
}

func (s *GatewayStub) ChainReadObj(p0 context.Context, p1 cid.Cid) ([]byte, error) {
	return *new([]byte), xerrors.New("method not supported")
}

func (s *GatewayStruct) GasEstimateMessageGas(p0 context.Context, p1 *types.Message, p2 *api.MessageSendSpec, p3 types.TipSetKey) (*types.Message, error) {
	return s.Internal.GasEstimateMessageGas(p0, p1, p2, p3)
}

func (s *GatewayStub) GasEstimateMessageGas(p0 context.Context, p1 *types.Message, p2 *api.MessageSendSpec, p3 types.TipSetKey) (*types.Message, error) {
	return nil, xerrors.New("method not supported")
}

func (s *GatewayStruct) MpoolPush(p0 context.Context, p1 *types.SignedMessage) (cid.Cid, error) {
	return s.Internal.MpoolPush(p0, p1)
}

func (s *GatewayStub) MpoolPush(p0 context.Context, p1 *types.SignedMessage) (cid.Cid, error) {
	return *new(cid.Cid), xerrors.New("method not supported")
}

func (s *GatewayStruct) MsigGetAvailableBalance(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (types.BigInt, error) {
	return s.Internal.MsigGetAvailableBalance(p0, p1, p2)
}

func (s *GatewayStub) MsigGetAvailableBalance(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (types.BigInt, error) {
	return *new(types.BigInt), xerrors.New("method not supported")
}

func (s *GatewayStruct) MsigGetPending(p0 context.Context, p1 address.Address, p2 types.TipSetKey) ([]*api.MsigTransaction, error) {
	return s.Internal.MsigGetPending(p0, p1, p2)
}

func (s *GatewayStub) MsigGetPending(p0 context.Context, p1 address.Address, p2 types.TipSetKey) ([]*api.MsigTransaction, error) {
	return *new([]*api.MsigTransaction), xerrors.New("method not supported")
}

func (s *GatewayStruct) MsigGetVested(p0 context.Context, p1 address.Address, p2 types.TipSetKey, p3 types.TipSetKey) (types.BigInt, error) {
	return s.Internal.MsigGetVested(p0, p1, p2, p3)
}

func (s *GatewayStub) MsigGetVested(p0 context.Context, p1 address.Address, p2 types.TipSetKey, p3 types.TipSetKey) (types.BigInt, error) {
	return *new(types.BigInt), xerrors.New("method not supported")
}

func (s *GatewayStruct) StateAccountKey(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (address.Address, error) {
	return s.Internal.StateAccountKey(p0, p1, p2)
}

func (s *GatewayStub) StateAccountKey(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (address.Address, error) {
	return *new(address.Address), xerrors.New("method not supported")
}

func (s *GatewayStruct) StateDealProviderCollateralBounds(p0 context.Context, p1 abi.PaddedPieceSize, p2 bool, p3 types.TipSetKey) (api.DealCollateralBounds, error) {
	return s.Internal.StateDealProviderCollateralBounds(p0, p1, p2, p3)
}

func (s *GatewayStub) StateDealProviderCollateralBounds(p0 context.Context, p1 abi.PaddedPieceSize, p2 bool, p3 types.TipSetKey) (api.DealCollateralBounds, error) {
	return *new(api.DealCollateralBounds), xerrors.New("method not supported")
}

func (s *GatewayStruct) StateGetActor(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*types.Actor, error) {
	return s.Internal.StateGetActor(p0, p1, p2)
}

func (s *GatewayStub) StateGetActor(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*types.Actor, error) {
	return nil, xerrors.New("method not supported")
}

func (s *GatewayStruct) StateGetReceipt(p0 context.Context, p1 cid.Cid, p2 types.TipSetKey) (*types.MessageReceipt, error) {
	return s.Internal.StateGetReceipt(p0, p1, p2)
}

func (s *GatewayStub) StateGetReceipt(p0 context.Context, p1 cid.Cid, p2 types.TipSetKey) (*types.MessageReceipt, error) {
	return nil, xerrors.New("method not supported")
}

func (s *GatewayStruct) StateListMiners(p0 context.Context, p1 types.TipSetKey) ([]address.Address, error) {
	return s.Internal.StateListMiners(p0, p1)
}

func (s *GatewayStub) StateListMiners(p0 context.Context, p1 types.TipSetKey) ([]address.Address, error) {
	return *new([]address.Address), xerrors.New("method not supported")
}

func (s *GatewayStruct) StateLookupID(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (address.Address, error) {
	return s.Internal.StateLookupID(p0, p1, p2)
}

func (s *GatewayStub) StateLookupID(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (address.Address, error) {
	return *new(address.Address), xerrors.New("method not supported")
}

func (s *GatewayStruct) StateMarketBalance(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (api.MarketBalance, error) {
	return s.Internal.StateMarketBalance(p0, p1, p2)
}

func (s *GatewayStub) StateMarketBalance(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (api.MarketBalance, error) {
	return *new(api.MarketBalance), xerrors.New("method not supported")
}

func (s *GatewayStruct) StateMarketStorageDeal(p0 context.Context, p1 abi.DealID, p2 types.TipSetKey) (*api.MarketDeal, error) {
	return s.Internal.StateMarketStorageDeal(p0, p1, p2)
}

func (s *GatewayStub) StateMarketStorageDeal(p0 context.Context, p1 abi.DealID, p2 types.TipSetKey) (*api.MarketDeal, error) {
	return nil, xerrors.New("method not supported")
}

func (s *GatewayStruct) StateMinerInfo(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (miner.MinerInfo, error) {
	return s.Internal.StateMinerInfo(p0, p1, p2)
}

func (s *GatewayStub) StateMinerInfo(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (miner.MinerInfo, error) {
	return *new(miner.MinerInfo), xerrors.New("method not supported")
}

func (s *GatewayStruct) StateMinerPower(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*api.MinerPower, error) {
	return s.Internal.StateMinerPower(p0, p1, p2)
}

func (s *GatewayStub) StateMinerPower(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*api.MinerPower, error) {
	return nil, xerrors.New("method not supported")
}

func (s *GatewayStruct) StateMinerProvingDeadline(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*dline.Info, error) {
	return s.Internal.StateMinerProvingDeadline(p0, p1, p2)
}

func (s *GatewayStub) StateMinerProvingDeadline(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*dline.Info, error) {
	return nil, xerrors.New("method not supported")
}

func (s *GatewayStruct) StateNetworkVersion(p0 context.Context, p1 types.TipSetKey) (network.Version, error) {
	return s.Internal.StateNetworkVersion(p0, p1)
}

func (s *GatewayStub) StateNetworkVersion(p0 context.Context, p1 types.TipSetKey) (network.Version, error) {
	return *new(network.Version), xerrors.New("method not supported")
}

func (s *GatewayStruct) StateSearchMsg(p0 context.Context, p1 cid.Cid) (*api.MsgLookup, error) {
	return s.Internal.StateSearchMsg(p0, p1)
}

func (s *GatewayStub) StateSearchMsg(p0 context.Context, p1 cid.Cid) (*api.MsgLookup, error) {
	return nil, xerrors.New("method not supported")
}

func (s *GatewayStruct) StateSectorGetInfo(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (*miner.SectorOnChainInfo, error) {
	return s.Internal.StateSectorGetInfo(p0, p1, p2, p3)
}

func (s *GatewayStub) StateSectorGetInfo(p0 context.Context, p1 address.Address, p2 abi.SectorNumber, p3 types.TipSetKey) (*miner.SectorOnChainInfo, error) {
	return nil, xerrors.New("method not supported")
}

func (s *GatewayStruct) StateVerifiedClientStatus(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*abi.StoragePower, error) {
	return s.Internal.StateVerifiedClientStatus(p0, p1, p2)
}

func (s *GatewayStub) StateVerifiedClientStatus(p0 context.Context, p1 address.Address, p2 types.TipSetKey) (*abi.StoragePower, error) {
	return nil, xerrors.New("method not supported")
}

func (s *GatewayStruct) StateWaitMsg(p0 context.Context, p1 cid.Cid, p2 uint64) (*api.MsgLookup, error) {
	return s.Internal.StateWaitMsg(p0, p1, p2)
}

func (s *GatewayStub) StateWaitMsg(p0 context.Context, p1 cid.Cid, p2 uint64) (*api.MsgLookup, error) {
	return nil, xerrors.New("method not supported")
}

func (s *GatewayStruct) WalletBalance(p0 context.Context, p1 address.Address) (types.BigInt, error) {
	return s.Internal.WalletBalance(p0, p1)
}

func (s *GatewayStub) WalletBalance(p0 context.Context, p1 address.Address) (types.BigInt, error) {
	return *new(types.BigInt), xerrors.New("method not supported")
}

var _ FullNode = new(FullNodeStruct)
var _ Gateway = new(GatewayStruct)
