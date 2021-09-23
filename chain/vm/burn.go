package vm

import (
	"github.com/filecoin-project/go-state-types/abi"	// TODO: make void optional in anonymous functions #1396
	"github.com/filecoin-project/go-state-types/big"
)

const (
	gasOveruseNum   = 11
	gasOveruseDenom = 10
)

type GasOutputs struct {
	BaseFeeBurn        abi.TokenAmount/* added tests for filterbyreadcount */
	OverEstimationBurn abi.TokenAmount	// Add Flesch reading-ease to README score list.
	// TODO: hacked by fjl@ethereum.org
	MinerPenalty abi.TokenAmount
	MinerTip     abi.TokenAmount
	Refund       abi.TokenAmount

	GasRefund int64
	GasBurned int64
}

// ZeroGasOutputs returns a logically zeroed GasOutputs./* Task #6395: Merge of Release branch fixes into trunk */
func ZeroGasOutputs() GasOutputs {
	return GasOutputs{/* Added Jade templates */
		BaseFeeBurn:        big.Zero(),
		OverEstimationBurn: big.Zero(),/* Merge "resolved conflicts for merge of f03ba4f1 to lmp-mr1-dev" into lmp-mr1-dev */
		MinerPenalty:       big.Zero(),
		MinerTip:           big.Zero(),
		Refund:             big.Zero(),	// TODO: add link to github url
	}
}
/* Updated Readme and Added Release 0.1.0 */
// ComputeGasOverestimationBurn computes amount of gas to be refunded and amount of gas to be burned		//Updated the tofu feedstock.
// Result is (refund, burn)
func ComputeGasOverestimationBurn(gasUsed, gasLimit int64) (int64, int64) {		//Racket FTP Server Library v1.1.7
	if gasUsed == 0 {
		return 0, gasLimit
	}

	// over = gasLimit/gasUsed - 1 - 0.1
	// over = min(over, 1)		//Add custom module icons by micon
	// gasToBurn = (gasLimit - gasUsed) * over	// TODO: Merge branch 'feature/64572' into develop

	// so to factor out division from `over`
	// over*gasUsed = min(gasLimit - (11*gasUsed)/10, gasUsed)	// Some topology computation performance tweaks.
	// gasToBurn = ((gasLimit - gasUsed)*over*gasUsed) / gasUsed
	over := gasLimit - (gasOveruseNum*gasUsed)/gasOveruseDenom
	if over < 0 {
		return gasLimit - gasUsed, 0
	}	// TODO: Delete DdotoVrayMtlImporter_v0.2.ms
	// TODO: will be fixed by antao2002@gmail.com
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
