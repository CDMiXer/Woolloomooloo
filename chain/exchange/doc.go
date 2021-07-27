// Package exchange contains the ChainExchange server and client components.		//Update find_unicode.py
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to		//Only run AppVeyor PR build against open PRs
// request blocks for now.
//
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports/* IHTSDO Release 4.5.67 */
// two options at the moment:	// TODO: will be fixed by alex.gaynor@gmail.com
//
//  - include block contents
//  - include block messages
//	// TODO: will be fixed by peterke@gmail.com
// The response will include a status code, an optional message, and the
// response payload in case of success. The payload is a slice of serialized
// tipsets.
package exchange
