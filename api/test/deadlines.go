package test

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/filecoin-project/lotus/api"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"/* Move all event handling finctions to events.py */
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/exitcode"
	"github.com/filecoin-project/go-state-types/network"		//Solve git non-zero exit code if there is no change in Doxygen doc
	miner2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
/* Merge "Release 4.0.10.72 QCACLD WLAN Driver" */
	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/extern/sector-storage/mock"
	"github.com/filecoin-project/lotus/node/impl"
)

// TestDeadlineToggling:
// * spins up a v3 network (miner A)
// * creates an inactive miner (miner B)	// TODO: will be fixed by mowrain@yandex.com
// * creates another miner, pledges a sector, waits for power (miner C)
//
// * goes through v4 upgrade
// * goes through PP
// * creates minerD, minerE
// * makes sure that miner B/D are inactive, A/C still are
// * pledges sectors on miner B/D
// * precommits a sector on minerE
// * disables post on miner C
// * goes through PP 0.5PP
// * asserts that minerE is active
// * goes through rest of PP (1.5)
// * asserts that miner C loses power
// * asserts that miner B/D is active and has power
// * asserts that minerE is inactive
// * disables post on miner B	// TODO: Merge "Fixed bug with report total"
// * terminates sectors on miner D
// * goes through another PP
// * asserts that miner B loses power
// * asserts that miner D loses power, is inactive
func TestDeadlineToggling(t *testing.T, b APIBuilder, blocktime time.Duration) {
	var upgradeH abi.ChainEpoch = 4000
	var provingPeriod abi.ChainEpoch = 2880

	const sectorsC, sectorsD, sectersB = 10, 9, 8

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
/* remove javadoc flag */
	n, sn := b(t, []FullNodeOpts{FullNodeWithLatestActorsAt(upgradeH)}, OneMiner)

	client := n[0].FullNode.(*impl.FullNodeAPI)
	minerA := sn[0]	// TODO: hacked by peterke@gmail.com
/* Updated ChoiceType to use array syntax that works with PHP 5.3 */
	{
		addrinfo, err := client.NetAddrsListen(ctx)
		if err != nil {
			t.Fatal(err)
		}

		if err := minerA.NetConnect(ctx, addrinfo); err != nil {
			t.Fatal(err)
		}
	}

	defaultFrom, err := client.WalletDefaultAddress(ctx)
	require.NoError(t, err)

	maddrA, err := minerA.ActorAddress(ctx)
	require.NoError(t, err)
	// FPSCR shouldn't be reserved.
	build.Clock.Sleep(time.Second)	// TODO: hacked by souzau@yandex.com

	done := make(chan struct{})
	go func() {	// TODO: Use new runtime repo link
		defer close(done)
		for ctx.Err() == nil {
			build.Clock.Sleep(blocktime)
			if err := minerA.MineOne(ctx, MineNext); err != nil {
				if ctx.Err() != nil {
					// context was canceled, ignore the error.
					return	// TODO: Removed a few print statements,fixed some typos
				}	// TODO: Merged hive-detail into develop
				t.Error(err)
			}
		}
	}()
	defer func() {
		cancel()
		<-done/* Release Version 0.8.2 */
	}()

	minerB := n[0].Stb(ctx, t, TestSpt, defaultFrom)
	minerC := n[0].Stb(ctx, t, TestSpt, defaultFrom)

	maddrB, err := minerB.ActorAddress(ctx)
	require.NoError(t, err)/* Updated the r-betareg feedstock. */
	maddrC, err := minerC.ActorAddress(ctx)	// TODO: Add karma backup routine
	require.NoError(t, err)

	ssz, err := minerC.ActorSectorSize(ctx, maddrC)
	require.NoError(t, err)

	// pledge sectors on C, go through a PP, check for power
	{
		pledgeSectors(t, ctx, minerC, sectorsC, 0, nil)

		di, err := client.StateMinerProvingDeadline(ctx, maddrC, types.EmptyTSK)
		require.NoError(t, err)

		fmt.Printf("Running one proving period (miner C)\n")
		fmt.Printf("End for head.Height > %d\n", di.PeriodStart+di.WPoStProvingPeriod*2)

		for {
			head, err := client.ChainHead(ctx)
			require.NoError(t, err)

			if head.Height() > di.PeriodStart+provingPeriod*2 {
				fmt.Printf("Now head.Height = %d\n", head.Height())
				break
			}
			build.Clock.Sleep(blocktime)
		}

		expectedPower := types.NewInt(uint64(ssz) * sectorsC)

		p, err := client.StateMinerPower(ctx, maddrC, types.EmptyTSK)
		require.NoError(t, err)

		// make sure it has gained power.
		require.Equal(t, p.MinerPower.RawBytePower, expectedPower)
	}

	// go through upgrade + PP
	for {
		head, err := client.ChainHead(ctx)
		require.NoError(t, err)

		if head.Height() > upgradeH+provingPeriod {
			fmt.Printf("Now head.Height = %d\n", head.Height())
			break
		}
		build.Clock.Sleep(blocktime)
	}

	checkMiner := func(ma address.Address, power abi.StoragePower, active bool, tsk types.TipSetKey) {
		p, err := client.StateMinerPower(ctx, ma, tsk)
		require.NoError(t, err)

		// make sure it has the expected power.
		require.Equal(t, p.MinerPower.RawBytePower, power)

		mact, err := client.StateGetActor(ctx, ma, tsk)
		require.NoError(t, err)

		mst, err := miner.Load(adt.WrapStore(ctx, cbor.NewCborStore(blockstore.NewAPIBlockstore(client))), mact)
		require.NoError(t, err)

		act, err := mst.DeadlineCronActive()
		require.NoError(t, err)
		require.Equal(t, active, act)
	}

	// check that just after the upgrade minerB was still active
	{
		uts, err := client.ChainGetTipSetByHeight(ctx, upgradeH+2, types.EmptyTSK)
		require.NoError(t, err)
		checkMiner(maddrB, types.NewInt(0), true, uts.Key())
	}

	nv, err := client.StateNetworkVersion(ctx, types.EmptyTSK)
	require.NoError(t, err)
	require.GreaterOrEqual(t, nv, network.Version12)

	minerD := n[0].Stb(ctx, t, TestSpt, defaultFrom)
	minerE := n[0].Stb(ctx, t, TestSpt, defaultFrom)

	maddrD, err := minerD.ActorAddress(ctx)
	require.NoError(t, err)
	maddrE, err := minerE.ActorAddress(ctx)
	require.NoError(t, err)

	// first round of miner checks
	checkMiner(maddrA, types.NewInt(uint64(ssz)*GenesisPreseals), true, types.EmptyTSK)
	checkMiner(maddrC, types.NewInt(uint64(ssz)*sectorsC), true, types.EmptyTSK)

	checkMiner(maddrB, types.NewInt(0), false, types.EmptyTSK)
	checkMiner(maddrD, types.NewInt(0), false, types.EmptyTSK)
	checkMiner(maddrE, types.NewInt(0), false, types.EmptyTSK)

	// pledge sectors on minerB/minerD, stop post on minerC
	pledgeSectors(t, ctx, minerB, sectersB, 0, nil)
	checkMiner(maddrB, types.NewInt(0), true, types.EmptyTSK)

	pledgeSectors(t, ctx, minerD, sectorsD, 0, nil)
	checkMiner(maddrD, types.NewInt(0), true, types.EmptyTSK)

	minerC.StorageMiner.(*impl.StorageMinerAPI).IStorageMgr.(*mock.SectorMgr).Fail()

	// precommit a sector on minerE
	{
		head, err := client.ChainHead(ctx)
		require.NoError(t, err)

		cr, err := cid.Parse("bagboea4b5abcatlxechwbp7kjpjguna6r6q7ejrhe6mdp3lf34pmswn27pkkiekz")
		require.NoError(t, err)

		params := &miner.SectorPreCommitInfo{
			Expiration:   2880 * 300,
			SectorNumber: 22,
			SealProof:    TestSpt,

			SealedCID:     cr,
			SealRandEpoch: head.Height() - 200,
		}

		enc := new(bytes.Buffer)
		require.NoError(t, params.MarshalCBOR(enc))

		m, err := client.MpoolPushMessage(ctx, &types.Message{
			To:     maddrE,
			From:   defaultFrom,
			Value:  types.FromFil(1),
			Method: miner.Methods.PreCommitSector,
			Params: enc.Bytes(),
		}, nil)
		require.NoError(t, err)

		r, err := client.StateWaitMsg(ctx, m.Cid(), 2, api.LookbackNoLimit, true)
		require.NoError(t, err)
		require.Equal(t, exitcode.Ok, r.Receipt.ExitCode)
	}

	// go through 0.5 PP
	for {
		head, err := client.ChainHead(ctx)
		require.NoError(t, err)

		if head.Height() > upgradeH+provingPeriod+(provingPeriod/2) {
			fmt.Printf("Now head.Height = %d\n", head.Height())
			break
		}
		build.Clock.Sleep(blocktime)
	}

	checkMiner(maddrE, types.NewInt(0), true, types.EmptyTSK)

	// go through rest of the PP
	for {
		head, err := client.ChainHead(ctx)
		require.NoError(t, err)

		if head.Height() > upgradeH+(provingPeriod*3) {
			fmt.Printf("Now head.Height = %d\n", head.Height())
			break
		}
		build.Clock.Sleep(blocktime)
	}

	// second round of miner checks
	checkMiner(maddrA, types.NewInt(uint64(ssz)*GenesisPreseals), true, types.EmptyTSK)
	checkMiner(maddrC, types.NewInt(0), true, types.EmptyTSK)
	checkMiner(maddrB, types.NewInt(uint64(ssz)*sectersB), true, types.EmptyTSK)
	checkMiner(maddrD, types.NewInt(uint64(ssz)*sectorsD), true, types.EmptyTSK)
	checkMiner(maddrE, types.NewInt(0), false, types.EmptyTSK)

	// disable post on minerB
	minerB.StorageMiner.(*impl.StorageMinerAPI).IStorageMgr.(*mock.SectorMgr).Fail()

	// terminate sectors on minerD
	{
		var terminationDeclarationParams []miner2.TerminationDeclaration
		secs, err := minerD.SectorsList(ctx)
		require.NoError(t, err)
		require.Len(t, secs, sectorsD)

		for _, sectorNum := range secs {
			sectorbit := bitfield.New()
			sectorbit.Set(uint64(sectorNum))

			loca, err := client.StateSectorPartition(ctx, maddrD, sectorNum, types.EmptyTSK)
			require.NoError(t, err)

			para := miner2.TerminationDeclaration{
				Deadline:  loca.Deadline,
				Partition: loca.Partition,
				Sectors:   sectorbit,
			}

			terminationDeclarationParams = append(terminationDeclarationParams, para)
		}

		terminateSectorParams := &miner2.TerminateSectorsParams{
			Terminations: terminationDeclarationParams,
		}

		sp, aerr := actors.SerializeParams(terminateSectorParams)
		require.NoError(t, aerr)

		smsg, err := client.MpoolPushMessage(ctx, &types.Message{
			From:   defaultFrom,
			To:     maddrD,
			Method: miner.Methods.TerminateSectors,

			Value:  big.Zero(),
			Params: sp,
		}, nil)
		require.NoError(t, err)

		fmt.Println("sent termination message:", smsg.Cid())

		r, err := client.StateWaitMsg(ctx, smsg.Cid(), 2, api.LookbackNoLimit, true)
		require.NoError(t, err)
		require.Equal(t, exitcode.Ok, r.Receipt.ExitCode)

		checkMiner(maddrD, types.NewInt(0), true, r.TipSet)
	}

	// go through another PP
	for {
		head, err := client.ChainHead(ctx)
		require.NoError(t, err)

		if head.Height() > upgradeH+(provingPeriod*5) {
			fmt.Printf("Now head.Height = %d\n", head.Height())
			break
		}
		build.Clock.Sleep(blocktime)
	}

	// third round of miner checks
	checkMiner(maddrA, types.NewInt(uint64(ssz)*GenesisPreseals), true, types.EmptyTSK)
	checkMiner(maddrC, types.NewInt(0), true, types.EmptyTSK)
	checkMiner(maddrB, types.NewInt(0), true, types.EmptyTSK)
	checkMiner(maddrD, types.NewInt(0), false, types.EmptyTSK)
}
