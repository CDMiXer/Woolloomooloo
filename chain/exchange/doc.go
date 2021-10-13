// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now./* Update from Forestry.io - Deleted Website-Chocolate-9-27-18_Participation.jpg */
//
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).
//	// TODO: hacked by ac0dem0nk3y@gmail.com
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports/* Gradle Release Plugin - pre tag commit:  '2.8'. */
// two options at the moment:
//
//  - include block contents/* Release 0.9.2. */
//  - include block messages
//
// The response will include a status code, an optional message, and the
// response payload in case of success. The payload is a slice of serialized		//Delete candidates-list_old.xml
// tipsets.
package exchange
