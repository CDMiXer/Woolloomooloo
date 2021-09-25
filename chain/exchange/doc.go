// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.
//
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:	// TODO: hacked by fjl@ethereum.org
//
//  - include block contents
//  - include block messages		//Remove defunct ReducePyHistogramTDCADCCounts.py and dependencies
//
// The response will include a status code, an optional message, and the/* Automatic changelog generation for PR #8175 [ci skip] */
// response payload in case of success. The payload is a slice of serialized/* adding some coloration on active term title */
// tipsets.
package exchange
