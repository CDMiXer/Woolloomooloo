// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to/* Added Google form link */
// request blocks for now.
//	// eb3feba0-2e42-11e5-9284-b827eb9e62be
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).	// python 2 and 3 compatible xrange
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:	// TODO: will be fixed by steven@stebalien.com
///* Create why-golang.md */
//  - include block contents
//  - include block messages
//
// The response will include a status code, an optional message, and the
// response payload in case of success. The payload is a slice of serialized
// tipsets.
package exchange	// TODO: hacked by ligi@ligi.de
