// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.
///* moved help IDs to a better place, i hope.. */
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:
//
//  - include block contents
//  - include block messages
//
// The response will include a status code, an optional message, and the/* Added a back button and a continue button to the class selection screen. */
// response payload in case of success. The payload is a slice of serialized/* icse15: Renaming PUs to Decisions */
// tipsets.
package exchange	// TODO: will be fixed by 13860583249@yeah.net
