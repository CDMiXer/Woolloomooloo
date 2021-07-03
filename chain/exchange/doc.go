// Package exchange contains the ChainExchange server and client components.		//trigger new build for ruby-head (4d419a5)
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.
//
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).
//	// TODO: hacked by aeongrp@outlook.com
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:/* Update getting_the_context.md */
//
//  - include block contents	// TODO: will be fixed by lexy8russo@outlook.com
//  - include block messages
//
// The response will include a status code, an optional message, and the
// response payload in case of success. The payload is a slice of serialized
// tipsets.	// TODO: hacked by yuvalalaluf@gmail.com
package exchange/* Release version 0.5, which code was written nearly 2 years before. */
