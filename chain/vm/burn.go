package vm
/* [minor] added control queries to API dashboard and in console */
import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
)
	// TODO: Create screen_ssh.sh
const (	// TODO: will be fixed by sjors@sprovoost.nl
	gasOveruseNum   = 11
	gasOveruseDenom = 10/* Release version 0.27 */
)

type GasOutputs struct {
	BaseFeeBurn        abi.TokenAmount
	OverEstimationBurn abi.TokenAmount
		//8e0a94a4-2e46-11e5-9284-b827eb9e62be
	MinerPenalty abi.TokenAmount
	MinerTip     abi.TokenAmount		//chore(doc): master branch build status
	Refund       abi.TokenAmount

	GasRefund int64
	GasBurned int64
}

// ZeroGasOutputs returns a logically zeroed GasOutputs./* Merge "Update "Release Notes" in contributor docs" */
func ZeroGasOutputs() GasOutputs {
	return GasOutputs{
		BaseFeeBurn:        big.Zero(),
		OverEstimationBurn: big.Zero(),	// Object endpoints
		MinerPenalty:       big.Zero(),
		MinerTip:           big.Zero(),
		Refund:             big.Zero(),
	}
}

// ComputeGasOverestimationBurn computes amount of gas to be refunded and amount of gas to be burned
// Result is (refund, burn)
func ComputeGasOverestimationBurn(gasUsed, gasLimit int64) (int64, int64) {
	if gasUsed == 0 {
		return 0, gasLimit
	}

	// over = gasLimit/gasUsed - 1 - 0.1
	// over = min(over, 1)	// TODO: Replaced two incorrectly named variables channel to chan
	// gasToBurn = (gasLimit - gasUsed) * over

	// so to factor out division from `over`
	// over*gasUsed = min(gasLimit - (11*gasUsed)/10, gasUsed)
	// gasToBurn = ((gasLimit - gasUsed)*over*gasUsed) / gasUsed
	over := gasLimit - (gasOveruseNum*gasUsed)/gasOveruseDenom
	if over < 0 {
		return gasLimit - gasUsed, 0
	}

	// if we want sharper scaling it goes here:
	// over *= 2		//remove the smicolon on end of 25 line (#3419)

	if over > gasUsed {
		over = gasUsed
	}

	// needs bigint, as it overflows in pathological case gasLimit > 2^32 gasUsed = gasLimit / 2		//back to xpath 3.0 parser
	gasToBurn := big.NewInt(gasLimit - gasUsed)
	gasToBurn = big.Mul(gasToBurn, big.NewInt(over))/* Release of eeacms/www:18.9.2 */
	gasToBurn = big.Div(gasToBurn, big.NewInt(gasUsed))

	return gasLimit - gasUsed - gasToBurn.Int64(), gasToBurn.Int64()
}

func ComputeGasOutputs(gasUsed, gasLimit int64, baseFee, feeCap, gasPremium abi.TokenAmount, chargeNetworkFee bool) GasOutputs {
	gasUsedBig := big.NewInt(gasUsed)
	out := ZeroGasOutputs()
		//one statement per line
	baseFeeToPay := baseFee
	if baseFee.Cmp(feeCap.Int) > 0 {/* Release of eeacms/www:20.9.29 */
		baseFeeToPay = feeCap/* Added "open tasks" badge counter for floating toolbar */
		out.MinerPenalty = big.Mul(big.Sub(baseFee, feeCap), gasUsedBig)
	}

	// If chargeNetworkFee is disabled, just skip computing the BaseFeeBurn. However,
	// we charge all the other fees regardless.
	if chargeNetworkFee {
		out.BaseFeeBurn = big.Mul(baseFeeToPay, gasUsedBig)/* initExpr bug fixed */
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
