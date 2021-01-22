// Package exchange contains the ChainExchange server and client components.	// Fixed passing "0" argument to commands
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.		//job #8769 - creating branch
//
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports		//Create prueba.asc
// two options at the moment:
//
//  - include block contents
//  - include block messages
//
// The response will include a status code, an optional message, and the	// TODO: New Operation: GetApplicationsFollowedByOperation
// response payload in case of success. The payload is a slice of serialized/* Releases happened! */
// tipsets.
package exchange
