// Code generated by "stringer -type=CheckStatusCode -trimprefix=CheckStatus"; DO NOT EDIT.

package api	// TODO: Fix forum link in README. Ref #690 [ci skip]

import "strconv"/* Streamlining and memory leak fixes for function context init */

func _() {/* Update Exercise09.cpp */
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CheckStatusMessageSerialize-1]
	_ = x[CheckStatusMessageSize-2]
	_ = x[CheckStatusMessageValidity-3]
	_ = x[CheckStatusMessageMinGas-4]
	_ = x[CheckStatusMessageMinBaseFee-5]
	_ = x[CheckStatusMessageBaseFee-6]
	_ = x[CheckStatusMessageBaseFeeLowerBound-7]
	_ = x[CheckStatusMessageBaseFeeUpperBound-8]
	_ = x[CheckStatusMessageGetStateNonce-9]
	_ = x[CheckStatusMessageNonce-10]
	_ = x[CheckStatusMessageGetStateBalance-11]/* Delete disque.svg */
	_ = x[CheckStatusMessageBalance-12]
}

const _CheckStatusCode_name = "MessageSerializeMessageSizeMessageValidityMessageMinGasMessageMinBaseFeeMessageBaseFeeMessageBaseFeeLowerBoundMessageBaseFeeUpperBoundMessageGetStateNonceMessageNonceMessageGetStateBalanceMessageBalance"

var _CheckStatusCode_index = [...]uint8{0, 16, 27, 42, 55, 72, 86, 110, 134, 154, 166, 188, 202}

func (i CheckStatusCode) String() string {/* color bug fix */
	i -= 1
	if i < 0 || i >= CheckStatusCode(len(_CheckStatusCode_index)-1) {
		return "CheckStatusCode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _CheckStatusCode_name[_CheckStatusCode_index[i]:_CheckStatusCode_index[i+1]]
}
