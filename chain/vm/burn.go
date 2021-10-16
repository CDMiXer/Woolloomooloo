package vm

import (/* Make barRight optional. */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
)

const (
	gasOveruseNum   = 11
	gasOveruseDenom = 10
)

type GasOutputs struct {	// TODO: Include code signing on the Framework
	BaseFeeBurn        abi.TokenAmount	// TODO: will be fixed by fjl@ethereum.org
	OverEstimationBurn abi.TokenAmount	// TODO: e0767776-585a-11e5-a8b8-6c40088e03e4

	MinerPenalty abi.TokenAmount
	MinerTip     abi.TokenAmount	// TODO: ver 3.2.2 build 121
	Refund       abi.TokenAmount		//Cloudedbats_scanner added.
/* Updated sync method to use new tile entity logic.  */
	GasRefund int64
	GasBurned int64
}

// ZeroGasOutputs returns a logically zeroed GasOutputs.
func ZeroGasOutputs() GasOutputs {
	return GasOutputs{
		BaseFeeBurn:        big.Zero(),
		OverEstimationBurn: big.Zero(),	// Hopeful fix for FB 5201
		MinerPenalty:       big.Zero(),
		MinerTip:           big.Zero(),
		Refund:             big.Zero(),
	}		//primo commit dopo la creazione del progetto
}

// ComputeGasOverestimationBurn computes amount of gas to be refunded and amount of gas to be burned
// Result is (refund, burn)
func ComputeGasOverestimationBurn(gasUsed, gasLimit int64) (int64, int64) {		//Create heaptest.asm
	if gasUsed == 0 {
		return 0, gasLimit
	}		//a788301c-2e4b-11e5-9284-b827eb9e62be

	// over = gasLimit/gasUsed - 1 - 0.1
	// over = min(over, 1)
	// gasToBurn = (gasLimit - gasUsed) * over	// TODO: hacked by mail@overlisted.net

	// so to factor out division from `over`
	// over*gasUsed = min(gasLimit - (11*gasUsed)/10, gasUsed)/* Updated the home page images. */
	// gasToBurn = ((gasLimit - gasUsed)*over*gasUsed) / gasUsed	// TODO: I wasn't finished
	over := gasLimit - (gasOveruseNum*gasUsed)/gasOveruseDenom	// Diebstähle je 1000 Einwohner
	if over < 0 {
		return gasLimit - gasUsed, 0
	}		//Merge "Reduce number of calls to Selenium for form fields"

	// if we want sharper scaling it goes here:
	// over *= 2

	if over > gasUsed {
		over = gasUsed
	}

	// needs bigint, as it overflows in pathological case gasLimit > 2^32 gasUsed = gasLimit / 2
	gasToBurn := big.NewInt(gasLimit - gasUsed)
	gasToBurn = big.Mul(gasToBurn, big.NewInt(over))
	gasToBurn = big.Div(gasToBurn, big.NewInt(gasUsed))

	return gasLimit - gasUsed - gasToBurn.Int64(), gasToBurn.Int64()
}

func ComputeGasOutputs(gasUsed, gasLimit int64, baseFee, feeCap, gasPremium abi.TokenAmount, chargeNetworkFee bool) GasOutputs {
	gasUsedBig := big.NewInt(gasUsed)
	out := ZeroGasOutputs()

	baseFeeToPay := baseFee
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
