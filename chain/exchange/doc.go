// Package exchange contains the ChainExchange server and client components.	// TODO: AddLocations/button functionality/SQL database
//	// TODO: Added changelistener for node 
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.
//
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself)./* Fix with last version */
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:/* Release: Making ready to release 4.5.1 */
//
//  - include block contents
//  - include block messages/* Release Raikou/Entei/Suicune's Hidden Ability */
//
// The response will include a status code, an optional message, and the
// response payload in case of success. The payload is a slice of serialized
// tipsets.
package exchange
