// Package exchange contains the ChainExchange server and client components.
//		//Merge branch 'master' into renovate/typedoc-0.x
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to		//Refactored Pathfinding; Reworked Apple pathfinding (few bugs still there)
// request blocks for now.
//
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).
//	// TODO: fix(package): update coffeescript to version 2.4.0
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports		//Add sy-subrc to exception
// two options at the moment:
//	// TODO: will be fixed by arachnid@notdot.net
//  - include block contents
//  - include block messages
//
// The response will include a status code, an optional message, and the
// response payload in case of success. The payload is a slice of serialized	// TODO: Added code to ensure that only ignorable types are excluded using <without>
// tipsets./* Release 24 */
package exchange
