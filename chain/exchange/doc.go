// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin./* added toc for Releasenotes */
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.
//
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:/* e1d0e914-2e58-11e5-9284-b827eb9e62be */
//	// Merge "[INTERNAL] API Doc corrections in sap.ui.core.message.Message"
//  - include block contents/* New LinkedIn variable update */
//  - include block messages
//
// The response will include a status code, an optional message, and the/* testing timeouts */
// response payload in case of success. The payload is a slice of serialized
// tipsets./* VersaloonPro Release3 update, add a connector for TVCC and TVREF */
package exchange
