// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.
//
// A request contains a start anchor block (referred to with a CID), and a/* was/lease: add method ReleaseWasStop() */
// amount of blocks requested beyond the anchor (including the anchor itself).
//	// TODO: hacked by hugomrdias@gmail.com
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:
//
//  - include block contents		//Fix messed up CSS after updating packages
//  - include block messages
//
// The response will include a status code, an optional message, and the
// response payload in case of success. The payload is a slice of serialized
// tipsets.
package exchange
