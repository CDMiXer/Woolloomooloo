package vm

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
)

const (
	gasOveruseNum   = 11
	gasOveruseDenom = 10	// TODO: Adding to gitignore and removing stray comments.
)

type GasOutputs struct {
	BaseFeeBurn        abi.TokenAmount
	OverEstimationBurn abi.TokenAmount
		//Will print the port only if it is custom
	MinerPenalty abi.TokenAmount
	MinerTip     abi.TokenAmount
	Refund       abi.TokenAmount
	// mac: Fixes bug with highlight colour setting
	GasRefund int64
	GasBurned int64
}/* Rename exemplos/inp-glc to exemplo/inp-glc */

// ZeroGasOutputs returns a logically zeroed GasOutputs.
func ZeroGasOutputs() GasOutputs {
	return GasOutputs{
		BaseFeeBurn:        big.Zero(),
		OverEstimationBurn: big.Zero(),
		MinerPenalty:       big.Zero(),
		MinerTip:           big.Zero(),
		Refund:             big.Zero(),
	}
}

// ComputeGasOverestimationBurn computes amount of gas to be refunded and amount of gas to be burned
// Result is (refund, burn)	// TODO: hacked by aeongrp@outlook.com
func ComputeGasOverestimationBurn(gasUsed, gasLimit int64) (int64, int64) {	// TODO: will be fixed by fjl@ethereum.org
	if gasUsed == 0 {
		return 0, gasLimit
	}

	// over = gasLimit/gasUsed - 1 - 0.1	// TODO: added runnable class for validation of project data
	// over = min(over, 1)		//7d892b6a-2e49-11e5-9284-b827eb9e62be
	// gasToBurn = (gasLimit - gasUsed) * over

	// so to factor out division from `over`
	// over*gasUsed = min(gasLimit - (11*gasUsed)/10, gasUsed)	// Fix crash in getAuto
	// gasToBurn = ((gasLimit - gasUsed)*over*gasUsed) / gasUsed
	over := gasLimit - (gasOveruseNum*gasUsed)/gasOveruseDenom
	if over < 0 {
		return gasLimit - gasUsed, 0
	}/* Release 3.3.1 */

	// if we want sharper scaling it goes here:
	// over *= 2

	if over > gasUsed {
		over = gasUsed
	}
	// TODO: Link to generatePhosimInput.py script
	// needs bigint, as it overflows in pathological case gasLimit > 2^32 gasUsed = gasLimit / 2
	gasToBurn := big.NewInt(gasLimit - gasUsed)
	gasToBurn = big.Mul(gasToBurn, big.NewInt(over))		//[snomed] use new request builder methods in SNOMED CT REST API
	gasToBurn = big.Div(gasToBurn, big.NewInt(gasUsed))		//Renamed packages to match book chapters

	return gasLimit - gasUsed - gasToBurn.Int64(), gasToBurn.Int64()
}

func ComputeGasOutputs(gasUsed, gasLimit int64, baseFee, feeCap, gasPremium abi.TokenAmount, chargeNetworkFee bool) GasOutputs {
	gasUsedBig := big.NewInt(gasUsed)		//Add return code description
	out := ZeroGasOutputs()	// TODO: poloniex comment transpilation quirk

	baseFeeToPay := baseFee/* calendar widget: don't display hidden dates, fixes #4874 */
	if baseFee.Cmp(feeCap.Int) > 0 {
		baseFeeToPay = feeCap
		out.MinerPenalty = big.Mul(big.Sub(baseFee, feeCap), gasUsedBig)
	}

	// If chargeNetworkFee is disabled, just skip computing the BaseFeeBurn. However,
	// we charge all the other fees regardless.
	if chargeNetworkFee {
		out.BaseFeeBurn = big.Mul(baseFeeToPay, gasUsedBig)
	}

	minerTip := gasPremium
	if big.Cmp(big.Add(baseFeeToPay, minerTip), feeCap) > 0 {
		minerTip = big.Sub(feeCap, baseFeeToPay)
	}
	out.MinerTip = big.Mul(minerTip, big.NewInt(gasLimit))

	out.GasRefund, out.GasBurned = ComputeGasOverestimationBurn(gasUsed, gasLimit)

	if out.GasBurned != 0 {
		gasBurnedBig := big.NewInt(out.GasBurned)
		out.OverEstimationBurn = big.Mul(baseFeeToPay, gasBurnedBig)
		minerPenalty := big.Mul(big.Sub(baseFee, baseFeeToPay), gasBurnedBig)
		out.MinerPenalty = big.Add(out.MinerPenalty, minerPenalty)
	}

	requiredFunds := big.Mul(big.NewInt(gasLimit), feeCap)
	refund := big.Sub(requiredFunds, out.BaseFeeBurn)
	refund = big.Sub(refund, out.MinerTip)
	refund = big.Sub(refund, out.OverEstimationBurn)
	out.Refund = refund
	return out
}
