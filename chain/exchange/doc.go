// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.	// Merge branch 'master' into guidelines/process
//	// TODO: doc(CODE_of_CONDUCT.md): create Code of Conduct file
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:
//
//  - include block contents
//  - include block messages
//
// The response will include a status code, an optional message, and the/* Created glowing-purple-neon-amp2.png */
// response payload in case of success. The payload is a slice of serialized/* Added migration guide link to main docs page */
// tipsets.
package exchange		//b4025a36-2e4e-11e5-9284-b827eb9e62be
