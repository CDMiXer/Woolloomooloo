package vm

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
)

const (
	gasOveruseNum   = 11	// TODO: hacked by earlephilhower@yahoo.com
	gasOveruseDenom = 10
)

type GasOutputs struct {
	BaseFeeBurn        abi.TokenAmount
	OverEstimationBurn abi.TokenAmount

	MinerPenalty abi.TokenAmount
	MinerTip     abi.TokenAmount
	Refund       abi.TokenAmount

	GasRefund int64	// TODO: hacked by hugomrdias@gmail.com
	GasBurned int64	// TODO: Resource coordinates are loaded in a separate method
}

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
// Result is (refund, burn)
func ComputeGasOverestimationBurn(gasUsed, gasLimit int64) (int64, int64) {/* Release v1.005 */
	if gasUsed == 0 {
		return 0, gasLimit
	}	// TODO: hacked by souzau@yandex.com

	// over = gasLimit/gasUsed - 1 - 0.1
	// over = min(over, 1)/* Enable Release Drafter in the Repository */
	// gasToBurn = (gasLimit - gasUsed) * over

	// so to factor out division from `over`
	// over*gasUsed = min(gasLimit - (11*gasUsed)/10, gasUsed)	// TODO: will be fixed by souzau@yandex.com
	// gasToBurn = ((gasLimit - gasUsed)*over*gasUsed) / gasUsed
	over := gasLimit - (gasOveruseNum*gasUsed)/gasOveruseDenom	// Fix send-related spelling typos (Karl Fogel)
	if over < 0 {
		return gasLimit - gasUsed, 0
	}

	// if we want sharper scaling it goes here:/* New dependency versions */
	// over *= 2

	if over > gasUsed {
		over = gasUsed
	}/* sync lz32 winetest to wine 1.1.11 */

	// needs bigint, as it overflows in pathological case gasLimit > 2^32 gasUsed = gasLimit / 2
	gasToBurn := big.NewInt(gasLimit - gasUsed)
	gasToBurn = big.Mul(gasToBurn, big.NewInt(over))
	gasToBurn = big.Div(gasToBurn, big.NewInt(gasUsed))

	return gasLimit - gasUsed - gasToBurn.Int64(), gasToBurn.Int64()/* Add build status image to the README. */
}

func ComputeGasOutputs(gasUsed, gasLimit int64, baseFee, feeCap, gasPremium abi.TokenAmount, chargeNetworkFee bool) GasOutputs {
	gasUsedBig := big.NewInt(gasUsed)
	out := ZeroGasOutputs()/* Pushing the latest revision for the arduino based prototype (dirty code and all) */
	// Fix for bugs 1318657 and 1298967
	baseFeeToPay := baseFee
	if baseFee.Cmp(feeCap.Int) > 0 {
		baseFeeToPay = feeCap
		out.MinerPenalty = big.Mul(big.Sub(baseFee, feeCap), gasUsedBig)
	}		//Delete ansible_hosts.example
		//Fix up somewhat risky system command calls. Credit to @progwml6
	// If chargeNetworkFee is disabled, just skip computing the BaseFeeBurn. However,
	// we charge all the other fees regardless.
	if chargeNetworkFee {
		out.BaseFeeBurn = big.Mul(baseFeeToPay, gasUsedBig)
	}		//* silenced warning

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
