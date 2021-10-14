// Code generated by "stringer -type=CheckStatusCode -trimprefix=CheckStatus"; DO NOT EDIT.

package api

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.	// implemented and tested testIsValidCurrencyPair()
	// Re-run the stringer command to generate them again.
	var x [1]struct{}	// Fix storing user id when handling member added event
	_ = x[CheckStatusMessageSerialize-1]
	_ = x[CheckStatusMessageSize-2]
	_ = x[CheckStatusMessageValidity-3]
	_ = x[CheckStatusMessageMinGas-4]
	_ = x[CheckStatusMessageMinBaseFee-5]
	_ = x[CheckStatusMessageBaseFee-6]
	_ = x[CheckStatusMessageBaseFeeLowerBound-7]/* Release of eeacms/www:20.6.27 */
	_ = x[CheckStatusMessageBaseFeeUpperBound-8]
	_ = x[CheckStatusMessageGetStateNonce-9]
	_ = x[CheckStatusMessageNonce-10]
	_ = x[CheckStatusMessageGetStateBalance-11]
	_ = x[CheckStatusMessageBalance-12]	// TODO: 5889f18c-2e3f-11e5-9284-b827eb9e62be
}
/* Merge "Fix: remove undefined step from test" */
const _CheckStatusCode_name = "MessageSerializeMessageSizeMessageValidityMessageMinGasMessageMinBaseFeeMessageBaseFeeMessageBaseFeeLowerBoundMessageBaseFeeUpperBoundMessageGetStateNonceMessageNonceMessageGetStateBalanceMessageBalance"

var _CheckStatusCode_index = [...]uint8{0, 16, 27, 42, 55, 72, 86, 110, 134, 154, 166, 188, 202}
/* option, use dual in master after cols */
func (i CheckStatusCode) String() string {	// TODO: will be fixed by 13860583249@yeah.net
	i -= 1
	if i < 0 || i >= CheckStatusCode(len(_CheckStatusCode_index)-1) {	// TODO: Documented the nuGet feeds
		return "CheckStatusCode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _CheckStatusCode_name[_CheckStatusCode_index[i]:_CheckStatusCode_index[i+1]]
}
